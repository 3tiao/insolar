//
// Copyright 2019 Insolar Technologies GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package genesis

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	"github.com/insolar/insolar/application/contract/noderecord"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/record"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/internal/ledger/artifact"
	"github.com/insolar/insolar/ledger/storage"
	"github.com/insolar/insolar/logicrunner/artifacts"
	"github.com/insolar/insolar/platformpolicy"
	"github.com/insolar/insolar/testutils"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const testDataPath = "gentestdata"

func mockArtifactClient(t *testing.T) *artifacts.ClientMock {
	amMock := artifacts.NewClientMock(t)
	amMock.RegisterRequestFunc = func(p context.Context, p1 insolar.Reference, p2 insolar.Parcel) (r *insolar.ID, r1 error) {
		id := testutils.RandomID()
		return &id, nil
	}
	amMock.ActivateObjectFunc = func(p context.Context, p1 insolar.Reference, p2 insolar.Reference, p3 insolar.Reference, p4 insolar.Reference, p5 bool, p6 []byte) (r artifacts.ObjectDescriptor, r1 error) {
		return artifacts.NewObjectDescriptorMock(t), nil
	}
	amMock.RegisterResultFunc = func(p context.Context, p1 insolar.Reference, p2 insolar.Reference, p3 []byte) (r *insolar.ID, r1 error) {
		id := testutils.RandomID()
		return &id, nil
	}
	return amMock
}

func mockArtifactManager(t *testing.T) artifact.Manager {
	osMock := storage.NewObjectStorageMock(t)

	osMock.SetRecordFunc = func(p context.Context, p1 insolar.ID, p2 insolar.PulseNumber, p3 record.VirtualRecord) (r *insolar.ID, r1 error) {
		id := testutils.RandomID()
		return &id, nil
	}

	return &artifact.Scope{
		PulseNumber:                insolar.FirstPulseNumber,
		ObjectStorage:              osMock,
		PlatformCryptographyScheme: platformpolicy.NewPlatformCryptographyScheme(),
	}
}

func mockArtifactClientWithRegisterRequestError(t *testing.T) *artifacts.ClientMock {
	amMock := artifacts.NewClientMock(t)
	amMock.RegisterRequestFunc = func(p context.Context, p1 insolar.Reference, p2 insolar.Parcel) (r *insolar.ID, r1 error) {
		return nil, errors.New("test reasons")
	}
	return amMock
}

func mockGenesis(t *testing.T, ac artifacts.Client, am artifact.Manager) *Genesis {
	ref := testutils.RandomRef()
	var discoveryNodes []Node
	discoveryNodes = append(discoveryNodes,
		Node{
			Role: "virtual",
		},
		Node{
			Role: "virtual",
		},
	)
	g := &Genesis{
		config: &Config{
			ReuseKeys:        true,
			DiscoveryNodes:   discoveryNodes,
			DiscoveryKeysDir: testDataPath,
			NodeKeysDir:      testDataPath,
		},
		rootDomainRef: &ref,
		nodeDomainRef: &ref,

		ArtifactsClient: ac,
		ArtifactManager: am,
	}
	return g
}

func mockContractBuilder(t *testing.T, g *Genesis) *ContractsBuilder {
	ref := testutils.RandomRef()
	cb := NewContractBuilder(g.ArtifactsClient, g.ArtifactManager)
	cb.Prototypes[nodeRecord] = &ref
	return cb
}

type genesisWithDataSuite struct {
	suite.Suite
}

func NewGenesisWithDataSuite() *genesisWithDataSuite {
	return &genesisWithDataSuite{
		Suite: suite.Suite{},
	}
}

// Init and run suite
func TestGenesisWithData(t *testing.T) {
	suite.Run(t, NewGenesisWithDataSuite())
}

func (s *genesisWithDataSuite) AfterTest(suiteName, testName string) {
	err := os.RemoveAll(testDataPath)
	if err != nil {
		s.T().Error("Can't remove testing data after test done", err)
	}
}

func (s *genesisWithDataSuite) TestCreateKeys() {
	g := &Genesis{
		config: &Config{
			KeysNameFormat: "/node_%d.json",
		},
	}
	ctx := inslogger.TestContext(s.T())
	amount := 5

	err := g.createKeys(ctx, testDataPath, amount)
	require.Nil(s.T(), err)

	files, _ := ioutil.ReadDir(testDataPath)
	require.Equal(s.T(), amount, len(files))
}

func (s *genesisWithDataSuite) TestUploadKeys_DontReuse() {
	g := Genesis{
		config: &Config{
			ReuseKeys: false,
		},
	}
	ctx := inslogger.TestContext(s.T())
	amount := 5

	info, err := g.uploadKeys(ctx, testDataPath, amount)
	require.Nil(s.T(), err)

	require.Equal(s.T(), amount, len(info))
}

func (s *genesisWithDataSuite) TestUploadKeys_Reuse() {
	g := Genesis{
		config: &Config{
			ReuseKeys: true,
		},
	}
	ctx := inslogger.TestContext(s.T())
	amount := 5
	err := g.createKeys(ctx, testDataPath, amount)
	require.Nil(s.T(), err)

	info, err := g.uploadKeys(ctx, testDataPath, amount)
	require.Nil(s.T(), err)

	require.Equal(s.T(), amount, len(info))
}

func (s *genesisWithDataSuite) TestUploadKeys_Reuse_WrongAmount() {
	g := Genesis{
		config: &Config{
			ReuseKeys: true,
		},
	}
	ctx := inslogger.TestContext(s.T())
	amount := 5
	err := g.createKeys(ctx, testDataPath, amount+5)
	require.Nil(s.T(), err)

	_, err = g.uploadKeys(ctx, testDataPath, amount)
	require.NotNil(s.T(), err)
	require.Contains(s.T(), err.Error(), "[ uploadKeys ] amount of nodes != amount of files in directory")
}

func TestUploadKeys_Reuse_DirNotExist(t *testing.T) {
	g := Genesis{
		config: &Config{
			ReuseKeys: true,
		},
	}
	ctx := inslogger.TestContext(t)
	amount := 5

	_, err := g.uploadKeys(ctx, testDataPath, amount)
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "[ uploadKeys ] can't read dir")
}

func TestActivateNodeRecord_RegisterRequest_Err(t *testing.T) {
	am := mockArtifactManager(t)
	ac := mockArtifactClientWithRegisterRequestError(t)
	g := mockGenesis(t, ac, am)
	cb := mockContractBuilder(t, g)
	ctx := inslogger.TestContext(t)
	publicKey := "fancy_public_key"
	name := "fancy_name"
	record := &noderecord.NodeRecord{
		Record: noderecord.RecordInfo{
			PublicKey: publicKey,
			Role:      insolar.StaticRoleVirtual,
		},
	}

	_, err := g.activateNodeRecord(ctx, cb, record, name)
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "[ activateNodeRecord ] Couldn't register request: test reasons")
}

func TestActivateNodeRecord_Activate_Err(t *testing.T) {
	am := mockArtifactManager(t)

	ac := artifacts.NewClientMock(t)
	ac.RegisterRequestFunc = func(p context.Context, p1 insolar.Reference, p2 insolar.Parcel) (r *insolar.ID, r1 error) {
		id := testutils.RandomID()
		return &id, nil
	}
	ac.ActivateObjectFunc = func(p context.Context, p1 insolar.Reference, p2 insolar.Reference, p3 insolar.Reference, p4 insolar.Reference, p5 bool, p6 []byte) (r artifacts.ObjectDescriptor, r1 error) {
		return nil, errors.New("test reasons")
	}

	g := mockGenesis(t, ac, am)
	cb := mockContractBuilder(t, g)
	ctx := inslogger.TestContext(t)
	publicKey := "fancy_public_key"
	name := "fancy_name"
	record := &noderecord.NodeRecord{
		Record: noderecord.RecordInfo{
			PublicKey: publicKey,
			Role:      insolar.StaticRoleVirtual,
		},
	}

	_, err := g.activateNodeRecord(ctx, cb, record, name)
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "[ activateNodeRecord ] Could'n activateNodeRecord node object: test reasons")
}

func TestActivateNodeRecord_RegisterResult_Err(t *testing.T) {
	am := mockArtifactManager(t)

	ac := artifacts.NewClientMock(t)
	ac.RegisterRequestFunc = func(p context.Context, p1 insolar.Reference, p2 insolar.Parcel) (r *insolar.ID, r1 error) {
		id := testutils.RandomID()
		return &id, nil
	}
	ac.ActivateObjectFunc = func(p context.Context, p1 insolar.Reference, p2 insolar.Reference, p3 insolar.Reference, p4 insolar.Reference, p5 bool, p6 []byte) (r artifacts.ObjectDescriptor, r1 error) {
		return artifacts.NewObjectDescriptorMock(t), nil
	}
	ac.RegisterResultFunc = func(p context.Context, p1 insolar.Reference, p2 insolar.Reference, p3 []byte) (r *insolar.ID, r1 error) {
		return nil, errors.New("test reasons")
	}

	g := mockGenesis(t, ac, am)
	cb := mockContractBuilder(t, g)
	ctx := inslogger.TestContext(t)
	publicKey := "fancy_public_key"
	name := "fancy_name"
	record := &noderecord.NodeRecord{
		Record: noderecord.RecordInfo{
			PublicKey: publicKey,
			Role:      insolar.StaticRoleVirtual,
		},
	}

	_, err := g.activateNodeRecord(ctx, cb, record, name)
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "[ activateNodeRecord ] Couldn't register result: test reasons")
}

func TestActivateNodeRecord(t *testing.T) {
	ac := mockArtifactClient(t)
	am := mockArtifactManager(t)
	g := mockGenesis(t, ac, am)

	cb := mockContractBuilder(t, g)
	ctx := inslogger.TestContext(t)
	publicKey := "fancy_public_key"
	name := "fancy_name"
	record := &noderecord.NodeRecord{
		Record: noderecord.RecordInfo{
			PublicKey: publicKey,
			Role:      insolar.StaticRoleVirtual,
		},
	}

	contract, err := g.activateNodeRecord(ctx, cb, record, name)
	require.Nil(t, err)
	require.NotNil(t, contract)
}

func TestActivateNodes_Err(t *testing.T) {
	ac := mockArtifactClientWithRegisterRequestError(t)
	am := mockArtifactManager(t)
	g := mockGenesis(t, ac, am)
	cb := mockContractBuilder(t, g)
	ctx := inslogger.TestContext(t)

	var nodes []nodeInfo
	nodes = append(nodes,
		nodeInfo{
			publicKey: "test_pk_1",
		},
		nodeInfo{
			publicKey: "test_pk_2",
		},
	)

	_, err := g.activateNodes(ctx, cb, nodes)
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "[ activateNodes ] Couldn't activateNodeRecord node instance:")
}

func TestActivateNodes(t *testing.T) {
	ac := mockArtifactClient(t)
	am := mockArtifactManager(t)
	g := mockGenesis(t, ac, am)
	cb := mockContractBuilder(t, g)
	ctx := inslogger.TestContext(t)

	var nodes []nodeInfo
	nodes = append(nodes,
		nodeInfo{
			publicKey: "test_pk_1",
		},
		nodeInfo{
			publicKey: "test_pk_2",
		},
	)

	updatedNodes, err := g.activateNodes(ctx, cb, nodes)
	require.Nil(t, err)
	require.Len(t, updatedNodes, len(nodes))
	for i := 0; i < len(nodes); i++ {
		require.Equal(t, nodes[i].publicKey, updatedNodes[i].publicKey)
		require.NotNil(t, updatedNodes[i].ref)
	}
}

func TestActivateDiscoveryNodes_DiffLen(t *testing.T) {
	ac := mockArtifactClient(t)
	am := mockArtifactManager(t)
	g := mockGenesis(t, ac, am)

	cb := mockContractBuilder(t, g)
	ctx := inslogger.TestContext(t)

	var nodes []nodeInfo
	nodes = append(nodes,
		nodeInfo{
			publicKey: "test_pk_1",
		},
	)

	_, err := g.activateDiscoveryNodes(ctx, cb, nodes)
	require.EqualError(t, err, "[ activateDiscoveryNodes ] len of nodesInfo param must be equal to len of DiscoveryNodes in genesis config")
}

func TestActivateDiscoveryNodes_Err(t *testing.T) {
	ac := mockArtifactClientWithRegisterRequestError(t)
	am := mockArtifactManager(t)
	g := mockGenesis(t, ac, am)

	cb := mockContractBuilder(t, g)
	ctx := inslogger.TestContext(t)

	var nodes []nodeInfo
	nodes = append(nodes,
		nodeInfo{
			publicKey: "test_pk_1",
		},
		nodeInfo{
			publicKey: "test_pk_2",
		},
	)
	require.Len(t, nodes, len(g.config.DiscoveryNodes))

	_, err := g.activateDiscoveryNodes(ctx, cb, nodes)
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "[ activateDiscoveryNodes ] Couldn't activateNodeRecord node instance:")
}

func TestActivateDiscoveryNodes(t *testing.T) {
	ac := mockArtifactClient(t)
	am := mockArtifactManager(t)
	g := mockGenesis(t, ac, am)

	cb := mockContractBuilder(t, g)
	ctx := inslogger.TestContext(t)

	var nodes []nodeInfo
	nodes = append(nodes,
		nodeInfo{
			publicKey: "test_pk_1",
		},
		nodeInfo{
			publicKey: "test_pk_2",
		},
	)
	require.Len(t, nodes, len(g.config.DiscoveryNodes))

	genesisNodes, err := g.activateDiscoveryNodes(ctx, cb, nodes)
	require.Nil(t, err)
	require.Len(t, genesisNodes, len(g.config.DiscoveryNodes))
	for i := 0; i < len(g.config.DiscoveryNodes); i++ {
		require.Equal(t, g.config.DiscoveryNodes[i].Role, genesisNodes[i].role)
		require.Equal(t, nodes[i].publicKey, genesisNodes[i].node.PublicKey)
		require.NotNil(t, genesisNodes[i].ref)
	}
}

func (s *genesisWithDataSuite) TestAddDiscoveryIndex_ActivateErr() {
	ac := mockArtifactClientWithRegisterRequestError(s.T())
	am := mockArtifactManager(s.T())
	g := mockGenesis(s.T(), ac, am)
	cb := mockContractBuilder(s.T(), g)
	ctx := inslogger.TestContext(s.T())
	err := g.createKeys(ctx, testDataPath, len(g.config.DiscoveryNodes))
	require.Nil(s.T(), err)

	indexMap := make(map[string]string)

	genesisNodes, resIndexMap, err := g.addDiscoveryIndex(ctx, cb, indexMap)
	require.NotNil(s.T(), err)
	require.Contains(s.T(), err.Error(), "[ addDiscoveryIndex ]: [ activateDiscoveryNodes ] Couldn't activateNodeRecord node instance")
	require.Empty(s.T(), genesisNodes)
	require.Empty(s.T(), resIndexMap)
}

func TestAddDiscoveryIndex_UploadErr(t *testing.T) {
	ac := mockArtifactClientWithRegisterRequestError(t)
	am := mockArtifactManager(t)
	g := mockGenesis(t, ac, am)
	g.config.DiscoveryKeysDir = "not_existed_testDataPath"
	cb := mockContractBuilder(t, g)
	ctx := inslogger.TestContext(t)

	indexMap := make(map[string]string)

	genesisNodes, resIndexMap, err := g.addDiscoveryIndex(ctx, cb, indexMap)
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "[ addDiscoveryIndex ]: [ uploadKeys ] can't read dir")
	require.Empty(t, genesisNodes)
	require.Empty(t, resIndexMap)
}

func (s *genesisWithDataSuite) TestAddDiscoveryIndex() {
	ac := mockArtifactClient(s.T())
	am := mockArtifactManager(s.T())
	g := mockGenesis(s.T(), ac, am)
	cb := mockContractBuilder(s.T(), g)
	ctx := inslogger.TestContext(s.T())
	err := g.createKeys(ctx, testDataPath, len(g.config.DiscoveryNodes))
	require.Nil(s.T(), err)

	indexMap := make(map[string]string)

	genesisNodes, resIndexMap, err := g.addDiscoveryIndex(ctx, cb, indexMap)
	require.Nil(s.T(), err)
	require.Len(s.T(), genesisNodes, len(g.config.DiscoveryNodes))
	require.Len(s.T(), resIndexMap, len(g.config.DiscoveryNodes))
}

func (s *genesisWithDataSuite) TestAddIndex_ActivateErr() {
	ac := mockArtifactClientWithRegisterRequestError(s.T())
	am := mockArtifactManager(s.T())
	g := mockGenesis(s.T(), ac, am)
	cb := mockContractBuilder(s.T(), g)
	ctx := inslogger.TestContext(s.T())
	err := g.createKeys(ctx, testDataPath, nodeAmount)
	require.Nil(s.T(), err)

	indexMap := make(map[string]string)

	resIndexMap, err := g.addIndex(ctx, cb, indexMap)
	require.NotNil(s.T(), err)
	require.Contains(s.T(), err.Error(), "[ addIndex ]: [ activateNodes ] Couldn't activateNodeRecord node instance")
	require.Empty(s.T(), resIndexMap)
}

func TestAddIndex_UploadErr(t *testing.T) {
	ac := mockArtifactClientWithRegisterRequestError(t)
	am := mockArtifactManager(t)
	g := mockGenesis(t, ac, am)
	g.config.NodeKeysDir = "not_existed_testDataPath"
	cb := mockContractBuilder(t, g)
	ctx := inslogger.TestContext(t)

	indexMap := make(map[string]string)

	resIndexMap, err := g.addIndex(ctx, cb, indexMap)
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "[ addIndex ]: [ uploadKeys ] can't read dir")
	require.Empty(t, resIndexMap)
}

func (s *genesisWithDataSuite) TestAddIndex() {
	ac := mockArtifactClient(s.T())
	am := mockArtifactManager(s.T())
	g := mockGenesis(s.T(), ac, am)
	cb := mockContractBuilder(s.T(), g)
	ctx := inslogger.TestContext(s.T())
	err := g.createKeys(ctx, testDataPath, nodeAmount)
	require.Nil(s.T(), err)

	indexMap := make(map[string]string)

	resIndexMap, err := g.addIndex(ctx, cb, indexMap)
	require.Nil(s.T(), err)
	require.Len(s.T(), resIndexMap, nodeAmount)
}
