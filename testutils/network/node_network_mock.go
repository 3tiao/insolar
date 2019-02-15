package network

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.9
The original interface "NodeNetwork" can be found in github.com/insolar/insolar/core
*/
import (
	"sync/atomic"
	"time"

	"github.com/gojuno/minimock"
	core "github.com/insolar/insolar/core"

	testify_assert "github.com/stretchr/testify/assert"
)

//NodeNetworkMock implements github.com/insolar/insolar/core.NodeNetwork
type NodeNetworkMock struct {
	t minimock.Tester

	GetActiveNodeFunc       func(p core.RecordRef) (r core.Node)
	GetActiveNodeCounter    uint64
	GetActiveNodePreCounter uint64
	GetActiveNodeMock       mNodeNetworkMockGetActiveNode

	GetActiveNodesFunc       func() (r []core.Node)
	GetActiveNodesCounter    uint64
	GetActiveNodesPreCounter uint64
	GetActiveNodesMock       mNodeNetworkMockGetActiveNodes

	GetActiveNodesByRoleFunc       func(p core.DynamicRole) (r []core.RecordRef)
	GetActiveNodesByRoleCounter    uint64
	GetActiveNodesByRolePreCounter uint64
	GetActiveNodesByRoleMock       mNodeNetworkMockGetActiveNodesByRole

	GetOriginFunc       func() (r core.Node)
	GetOriginCounter    uint64
	GetOriginPreCounter uint64
	GetOriginMock       mNodeNetworkMockGetOrigin

	GetStateFunc       func() (r core.NodeNetworkState)
	GetStateCounter    uint64
	GetStatePreCounter uint64
	GetStateMock       mNodeNetworkMockGetState
}

//NewNodeNetworkMock returns a mock for github.com/insolar/insolar/core.NodeNetwork
func NewNodeNetworkMock(t minimock.Tester) *NodeNetworkMock {
	m := &NodeNetworkMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetActiveNodeMock = mNodeNetworkMockGetActiveNode{mock: m}
	m.GetActiveNodesMock = mNodeNetworkMockGetActiveNodes{mock: m}
	m.GetActiveNodesByRoleMock = mNodeNetworkMockGetActiveNodesByRole{mock: m}
	m.GetOriginMock = mNodeNetworkMockGetOrigin{mock: m}
	m.GetStateMock = mNodeNetworkMockGetState{mock: m}

	return m
}

type mNodeNetworkMockGetActiveNode struct {
	mock              *NodeNetworkMock
	mainExpectation   *NodeNetworkMockGetActiveNodeExpectation
	expectationSeries []*NodeNetworkMockGetActiveNodeExpectation
}

type NodeNetworkMockGetActiveNodeExpectation struct {
	input  *NodeNetworkMockGetActiveNodeInput
	result *NodeNetworkMockGetActiveNodeResult
}

type NodeNetworkMockGetActiveNodeInput struct {
	p core.RecordRef
}

type NodeNetworkMockGetActiveNodeResult struct {
	r core.Node
}

//Expect specifies that invocation of NodeNetwork.GetActiveNode is expected from 1 to Infinity times
func (m *mNodeNetworkMockGetActiveNode) Expect(p core.RecordRef) *mNodeNetworkMockGetActiveNode {
	m.mock.GetActiveNodeFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &NodeNetworkMockGetActiveNodeExpectation{}
	}
	m.mainExpectation.input = &NodeNetworkMockGetActiveNodeInput{p}
	return m
}

//Return specifies results of invocation of NodeNetwork.GetActiveNode
func (m *mNodeNetworkMockGetActiveNode) Return(r core.Node) *NodeNetworkMock {
	m.mock.GetActiveNodeFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &NodeNetworkMockGetActiveNodeExpectation{}
	}
	m.mainExpectation.result = &NodeNetworkMockGetActiveNodeResult{r}
	return m.mock
}

//ExpectOnce specifies that invocation of NodeNetwork.GetActiveNode is expected once
func (m *mNodeNetworkMockGetActiveNode) ExpectOnce(p core.RecordRef) *NodeNetworkMockGetActiveNodeExpectation {
	m.mock.GetActiveNodeFunc = nil
	m.mainExpectation = nil

	expectation := &NodeNetworkMockGetActiveNodeExpectation{}
	expectation.input = &NodeNetworkMockGetActiveNodeInput{p}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *NodeNetworkMockGetActiveNodeExpectation) Return(r core.Node) {
	e.result = &NodeNetworkMockGetActiveNodeResult{r}
}

//Set uses given function f as a mock of NodeNetwork.GetActiveNode method
func (m *mNodeNetworkMockGetActiveNode) Set(f func(p core.RecordRef) (r core.Node)) *NodeNetworkMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.GetActiveNodeFunc = f
	return m.mock
}

//GetActiveNode implements github.com/insolar/insolar/core.NodeNetwork interface
func (m *NodeNetworkMock) GetActiveNode(p core.RecordRef) (r core.Node) {
	counter := atomic.AddUint64(&m.GetActiveNodePreCounter, 1)
	defer atomic.AddUint64(&m.GetActiveNodeCounter, 1)

	if len(m.GetActiveNodeMock.expectationSeries) > 0 {
		if counter > uint64(len(m.GetActiveNodeMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to NodeNetworkMock.GetActiveNode. %v", p)
			return
		}

		input := m.GetActiveNodeMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, NodeNetworkMockGetActiveNodeInput{p}, "NodeNetwork.GetActiveNode got unexpected parameters")

		result := m.GetActiveNodeMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the NodeNetworkMock.GetActiveNode")
			return
		}

		r = result.r

		return
	}

	if m.GetActiveNodeMock.mainExpectation != nil {

		input := m.GetActiveNodeMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, NodeNetworkMockGetActiveNodeInput{p}, "NodeNetwork.GetActiveNode got unexpected parameters")
		}

		result := m.GetActiveNodeMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the NodeNetworkMock.GetActiveNode")
		}

		r = result.r

		return
	}

	if m.GetActiveNodeFunc == nil {
		m.t.Fatalf("Unexpected call to NodeNetworkMock.GetActiveNode. %v", p)
		return
	}

	return m.GetActiveNodeFunc(p)
}

//GetActiveNodeMinimockCounter returns a count of NodeNetworkMock.GetActiveNodeFunc invocations
func (m *NodeNetworkMock) GetActiveNodeMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.GetActiveNodeCounter)
}

//GetActiveNodeMinimockPreCounter returns the value of NodeNetworkMock.GetActiveNode invocations
func (m *NodeNetworkMock) GetActiveNodeMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.GetActiveNodePreCounter)
}

//GetActiveNodeFinished returns true if mock invocations count is ok
func (m *NodeNetworkMock) GetActiveNodeFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.GetActiveNodeMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.GetActiveNodeCounter) == uint64(len(m.GetActiveNodeMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.GetActiveNodeMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.GetActiveNodeCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.GetActiveNodeFunc != nil {
		return atomic.LoadUint64(&m.GetActiveNodeCounter) > 0
	}

	return true
}

type mNodeNetworkMockGetActiveNodes struct {
	mock              *NodeNetworkMock
	mainExpectation   *NodeNetworkMockGetActiveNodesExpectation
	expectationSeries []*NodeNetworkMockGetActiveNodesExpectation
}

type NodeNetworkMockGetActiveNodesExpectation struct {
	result *NodeNetworkMockGetActiveNodesResult
}

type NodeNetworkMockGetActiveNodesResult struct {
	r []core.Node
}

//Expect specifies that invocation of NodeNetwork.GetActiveNodes is expected from 1 to Infinity times
func (m *mNodeNetworkMockGetActiveNodes) Expect() *mNodeNetworkMockGetActiveNodes {
	m.mock.GetActiveNodesFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &NodeNetworkMockGetActiveNodesExpectation{}
	}

	return m
}

//Return specifies results of invocation of NodeNetwork.GetActiveNodes
func (m *mNodeNetworkMockGetActiveNodes) Return(r []core.Node) *NodeNetworkMock {
	m.mock.GetActiveNodesFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &NodeNetworkMockGetActiveNodesExpectation{}
	}
	m.mainExpectation.result = &NodeNetworkMockGetActiveNodesResult{r}
	return m.mock
}

//ExpectOnce specifies that invocation of NodeNetwork.GetActiveNodes is expected once
func (m *mNodeNetworkMockGetActiveNodes) ExpectOnce() *NodeNetworkMockGetActiveNodesExpectation {
	m.mock.GetActiveNodesFunc = nil
	m.mainExpectation = nil

	expectation := &NodeNetworkMockGetActiveNodesExpectation{}

	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *NodeNetworkMockGetActiveNodesExpectation) Return(r []core.Node) {
	e.result = &NodeNetworkMockGetActiveNodesResult{r}
}

//Set uses given function f as a mock of NodeNetwork.GetActiveNodes method
func (m *mNodeNetworkMockGetActiveNodes) Set(f func() (r []core.Node)) *NodeNetworkMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.GetActiveNodesFunc = f
	return m.mock
}

//GetActiveNodes implements github.com/insolar/insolar/core.NodeNetwork interface
func (m *NodeNetworkMock) GetActiveNodes() (r []core.Node) {
	counter := atomic.AddUint64(&m.GetActiveNodesPreCounter, 1)
	defer atomic.AddUint64(&m.GetActiveNodesCounter, 1)

	if len(m.GetActiveNodesMock.expectationSeries) > 0 {
		if counter > uint64(len(m.GetActiveNodesMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to NodeNetworkMock.GetActiveNodes.")
			return
		}

		result := m.GetActiveNodesMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the NodeNetworkMock.GetActiveNodes")
			return
		}

		r = result.r

		return
	}

	if m.GetActiveNodesMock.mainExpectation != nil {

		result := m.GetActiveNodesMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the NodeNetworkMock.GetActiveNodes")
		}

		r = result.r

		return
	}

	if m.GetActiveNodesFunc == nil {
		m.t.Fatalf("Unexpected call to NodeNetworkMock.GetActiveNodes.")
		return
	}

	return m.GetActiveNodesFunc()
}

//GetActiveNodesMinimockCounter returns a count of NodeNetworkMock.GetActiveNodesFunc invocations
func (m *NodeNetworkMock) GetActiveNodesMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.GetActiveNodesCounter)
}

//GetActiveNodesMinimockPreCounter returns the value of NodeNetworkMock.GetActiveNodes invocations
func (m *NodeNetworkMock) GetActiveNodesMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.GetActiveNodesPreCounter)
}

//GetActiveNodesFinished returns true if mock invocations count is ok
func (m *NodeNetworkMock) GetActiveNodesFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.GetActiveNodesMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.GetActiveNodesCounter) == uint64(len(m.GetActiveNodesMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.GetActiveNodesMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.GetActiveNodesCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.GetActiveNodesFunc != nil {
		return atomic.LoadUint64(&m.GetActiveNodesCounter) > 0
	}

	return true
}

type mNodeNetworkMockGetActiveNodesByRole struct {
	mock              *NodeNetworkMock
	mainExpectation   *NodeNetworkMockGetActiveNodesByRoleExpectation
	expectationSeries []*NodeNetworkMockGetActiveNodesByRoleExpectation
}

type NodeNetworkMockGetActiveNodesByRoleExpectation struct {
	input  *NodeNetworkMockGetActiveNodesByRoleInput
	result *NodeNetworkMockGetActiveNodesByRoleResult
}

type NodeNetworkMockGetActiveNodesByRoleInput struct {
	p core.DynamicRole
}

type NodeNetworkMockGetActiveNodesByRoleResult struct {
	r []core.RecordRef
}

//Expect specifies that invocation of NodeNetwork.GetActiveNodesByRole is expected from 1 to Infinity times
func (m *mNodeNetworkMockGetActiveNodesByRole) Expect(p core.DynamicRole) *mNodeNetworkMockGetActiveNodesByRole {
	m.mock.GetActiveNodesByRoleFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &NodeNetworkMockGetActiveNodesByRoleExpectation{}
	}
	m.mainExpectation.input = &NodeNetworkMockGetActiveNodesByRoleInput{p}
	return m
}

//Return specifies results of invocation of NodeNetwork.GetActiveNodesByRole
func (m *mNodeNetworkMockGetActiveNodesByRole) Return(r []core.RecordRef) *NodeNetworkMock {
	m.mock.GetActiveNodesByRoleFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &NodeNetworkMockGetActiveNodesByRoleExpectation{}
	}
	m.mainExpectation.result = &NodeNetworkMockGetActiveNodesByRoleResult{r}
	return m.mock
}

//ExpectOnce specifies that invocation of NodeNetwork.GetActiveNodesByRole is expected once
func (m *mNodeNetworkMockGetActiveNodesByRole) ExpectOnce(p core.DynamicRole) *NodeNetworkMockGetActiveNodesByRoleExpectation {
	m.mock.GetActiveNodesByRoleFunc = nil
	m.mainExpectation = nil

	expectation := &NodeNetworkMockGetActiveNodesByRoleExpectation{}
	expectation.input = &NodeNetworkMockGetActiveNodesByRoleInput{p}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *NodeNetworkMockGetActiveNodesByRoleExpectation) Return(r []core.RecordRef) {
	e.result = &NodeNetworkMockGetActiveNodesByRoleResult{r}
}

//Set uses given function f as a mock of NodeNetwork.GetActiveNodesByRole method
func (m *mNodeNetworkMockGetActiveNodesByRole) Set(f func(p core.DynamicRole) (r []core.RecordRef)) *NodeNetworkMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.GetActiveNodesByRoleFunc = f
	return m.mock
}

//GetActiveNodesByRole implements github.com/insolar/insolar/core.NodeNetwork interface
func (m *NodeNetworkMock) GetActiveNodesByRole(p core.DynamicRole) (r []core.RecordRef) {
	counter := atomic.AddUint64(&m.GetActiveNodesByRolePreCounter, 1)
	defer atomic.AddUint64(&m.GetActiveNodesByRoleCounter, 1)

	if len(m.GetActiveNodesByRoleMock.expectationSeries) > 0 {
		if counter > uint64(len(m.GetActiveNodesByRoleMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to NodeNetworkMock.GetActiveNodesByRole. %v", p)
			return
		}

		input := m.GetActiveNodesByRoleMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, NodeNetworkMockGetActiveNodesByRoleInput{p}, "NodeNetwork.GetActiveNodesByRole got unexpected parameters")

		result := m.GetActiveNodesByRoleMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the NodeNetworkMock.GetActiveNodesByRole")
			return
		}

		r = result.r

		return
	}

	if m.GetActiveNodesByRoleMock.mainExpectation != nil {

		input := m.GetActiveNodesByRoleMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, NodeNetworkMockGetActiveNodesByRoleInput{p}, "NodeNetwork.GetActiveNodesByRole got unexpected parameters")
		}

		result := m.GetActiveNodesByRoleMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the NodeNetworkMock.GetActiveNodesByRole")
		}

		r = result.r

		return
	}

	if m.GetActiveNodesByRoleFunc == nil {
		m.t.Fatalf("Unexpected call to NodeNetworkMock.GetActiveNodesByRole. %v", p)
		return
	}

	return m.GetActiveNodesByRoleFunc(p)
}

//GetActiveNodesByRoleMinimockCounter returns a count of NodeNetworkMock.GetActiveNodesByRoleFunc invocations
func (m *NodeNetworkMock) GetActiveNodesByRoleMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.GetActiveNodesByRoleCounter)
}

//GetActiveNodesByRoleMinimockPreCounter returns the value of NodeNetworkMock.GetActiveNodesByRole invocations
func (m *NodeNetworkMock) GetActiveNodesByRoleMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.GetActiveNodesByRolePreCounter)
}

//GetActiveNodesByRoleFinished returns true if mock invocations count is ok
func (m *NodeNetworkMock) GetActiveNodesByRoleFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.GetActiveNodesByRoleMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.GetActiveNodesByRoleCounter) == uint64(len(m.GetActiveNodesByRoleMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.GetActiveNodesByRoleMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.GetActiveNodesByRoleCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.GetActiveNodesByRoleFunc != nil {
		return atomic.LoadUint64(&m.GetActiveNodesByRoleCounter) > 0
	}

	return true
}

type mNodeNetworkMockGetOrigin struct {
	mock              *NodeNetworkMock
	mainExpectation   *NodeNetworkMockGetOriginExpectation
	expectationSeries []*NodeNetworkMockGetOriginExpectation
}

type NodeNetworkMockGetOriginExpectation struct {
	result *NodeNetworkMockGetOriginResult
}

type NodeNetworkMockGetOriginResult struct {
	r core.Node
}

//Expect specifies that invocation of NodeNetwork.GetOrigin is expected from 1 to Infinity times
func (m *mNodeNetworkMockGetOrigin) Expect() *mNodeNetworkMockGetOrigin {
	m.mock.GetOriginFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &NodeNetworkMockGetOriginExpectation{}
	}

	return m
}

//Return specifies results of invocation of NodeNetwork.GetOrigin
func (m *mNodeNetworkMockGetOrigin) Return(r core.Node) *NodeNetworkMock {
	m.mock.GetOriginFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &NodeNetworkMockGetOriginExpectation{}
	}
	m.mainExpectation.result = &NodeNetworkMockGetOriginResult{r}
	return m.mock
}

//ExpectOnce specifies that invocation of NodeNetwork.GetOrigin is expected once
func (m *mNodeNetworkMockGetOrigin) ExpectOnce() *NodeNetworkMockGetOriginExpectation {
	m.mock.GetOriginFunc = nil
	m.mainExpectation = nil

	expectation := &NodeNetworkMockGetOriginExpectation{}

	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *NodeNetworkMockGetOriginExpectation) Return(r core.Node) {
	e.result = &NodeNetworkMockGetOriginResult{r}
}

//Set uses given function f as a mock of NodeNetwork.GetOrigin method
func (m *mNodeNetworkMockGetOrigin) Set(f func() (r core.Node)) *NodeNetworkMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.GetOriginFunc = f
	return m.mock
}

//GetOrigin implements github.com/insolar/insolar/core.NodeNetwork interface
func (m *NodeNetworkMock) GetOrigin() (r core.Node) {
	counter := atomic.AddUint64(&m.GetOriginPreCounter, 1)
	defer atomic.AddUint64(&m.GetOriginCounter, 1)

	if len(m.GetOriginMock.expectationSeries) > 0 {
		if counter > uint64(len(m.GetOriginMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to NodeNetworkMock.GetOrigin.")
			return
		}

		result := m.GetOriginMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the NodeNetworkMock.GetOrigin")
			return
		}

		r = result.r

		return
	}

	if m.GetOriginMock.mainExpectation != nil {

		result := m.GetOriginMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the NodeNetworkMock.GetOrigin")
		}

		r = result.r

		return
	}

	if m.GetOriginFunc == nil {
		m.t.Fatalf("Unexpected call to NodeNetworkMock.GetOrigin.")
		return
	}

	return m.GetOriginFunc()
}

//GetOriginMinimockCounter returns a count of NodeNetworkMock.GetOriginFunc invocations
func (m *NodeNetworkMock) GetOriginMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.GetOriginCounter)
}

//GetOriginMinimockPreCounter returns the value of NodeNetworkMock.GetOrigin invocations
func (m *NodeNetworkMock) GetOriginMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.GetOriginPreCounter)
}

//GetOriginFinished returns true if mock invocations count is ok
func (m *NodeNetworkMock) GetOriginFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.GetOriginMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.GetOriginCounter) == uint64(len(m.GetOriginMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.GetOriginMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.GetOriginCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.GetOriginFunc != nil {
		return atomic.LoadUint64(&m.GetOriginCounter) > 0
	}

	return true
}

type mNodeNetworkMockGetState struct {
	mock              *NodeNetworkMock
	mainExpectation   *NodeNetworkMockGetStateExpectation
	expectationSeries []*NodeNetworkMockGetStateExpectation
}

type NodeNetworkMockGetStateExpectation struct {
	result *NodeNetworkMockGetStateResult
}

type NodeNetworkMockGetStateResult struct {
	r core.NodeNetworkState
}

//Expect specifies that invocation of NodeNetwork.GetState is expected from 1 to Infinity times
func (m *mNodeNetworkMockGetState) Expect() *mNodeNetworkMockGetState {
	m.mock.GetStateFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &NodeNetworkMockGetStateExpectation{}
	}

	return m
}

//Return specifies results of invocation of NodeNetwork.GetState
func (m *mNodeNetworkMockGetState) Return(r core.NodeNetworkState) *NodeNetworkMock {
	m.mock.GetStateFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &NodeNetworkMockGetStateExpectation{}
	}
	m.mainExpectation.result = &NodeNetworkMockGetStateResult{r}
	return m.mock
}

//ExpectOnce specifies that invocation of NodeNetwork.GetState is expected once
func (m *mNodeNetworkMockGetState) ExpectOnce() *NodeNetworkMockGetStateExpectation {
	m.mock.GetStateFunc = nil
	m.mainExpectation = nil

	expectation := &NodeNetworkMockGetStateExpectation{}

	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *NodeNetworkMockGetStateExpectation) Return(r core.NodeNetworkState) {
	e.result = &NodeNetworkMockGetStateResult{r}
}

//Set uses given function f as a mock of NodeNetwork.GetState method
func (m *mNodeNetworkMockGetState) Set(f func() (r core.NodeNetworkState)) *NodeNetworkMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.GetStateFunc = f
	return m.mock
}

//GetState implements github.com/insolar/insolar/core.NodeNetwork interface
func (m *NodeNetworkMock) GetState() (r core.NodeNetworkState) {
	counter := atomic.AddUint64(&m.GetStatePreCounter, 1)
	defer atomic.AddUint64(&m.GetStateCounter, 1)

	if len(m.GetStateMock.expectationSeries) > 0 {
		if counter > uint64(len(m.GetStateMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to NodeNetworkMock.GetState.")
			return
		}

		result := m.GetStateMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the NodeNetworkMock.GetState")
			return
		}

		r = result.r

		return
	}

	if m.GetStateMock.mainExpectation != nil {

		result := m.GetStateMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the NodeNetworkMock.GetState")
		}

		r = result.r

		return
	}

	if m.GetStateFunc == nil {
		m.t.Fatalf("Unexpected call to NodeNetworkMock.GetState.")
		return
	}

	return m.GetStateFunc()
}

//GetStateMinimockCounter returns a count of NodeNetworkMock.GetStateFunc invocations
func (m *NodeNetworkMock) GetStateMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.GetStateCounter)
}

//GetStateMinimockPreCounter returns the value of NodeNetworkMock.GetState invocations
func (m *NodeNetworkMock) GetStateMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.GetStatePreCounter)
}

//GetStateFinished returns true if mock invocations count is ok
func (m *NodeNetworkMock) GetStateFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.GetStateMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.GetStateCounter) == uint64(len(m.GetStateMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.GetStateMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.GetStateCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.GetStateFunc != nil {
		return atomic.LoadUint64(&m.GetStateCounter) > 0
	}

	return true
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *NodeNetworkMock) ValidateCallCounters() {

	if !m.GetActiveNodeFinished() {
		m.t.Fatal("Expected call to NodeNetworkMock.GetActiveNode")
	}

	if !m.GetActiveNodesFinished() {
		m.t.Fatal("Expected call to NodeNetworkMock.GetActiveNodes")
	}

	if !m.GetActiveNodesByRoleFinished() {
		m.t.Fatal("Expected call to NodeNetworkMock.GetActiveNodesByRole")
	}

	if !m.GetOriginFinished() {
		m.t.Fatal("Expected call to NodeNetworkMock.GetOrigin")
	}

	if !m.GetStateFinished() {
		m.t.Fatal("Expected call to NodeNetworkMock.GetState")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *NodeNetworkMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *NodeNetworkMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *NodeNetworkMock) MinimockFinish() {

	if !m.GetActiveNodeFinished() {
		m.t.Fatal("Expected call to NodeNetworkMock.GetActiveNode")
	}

	if !m.GetActiveNodesFinished() {
		m.t.Fatal("Expected call to NodeNetworkMock.GetActiveNodes")
	}

	if !m.GetActiveNodesByRoleFinished() {
		m.t.Fatal("Expected call to NodeNetworkMock.GetActiveNodesByRole")
	}

	if !m.GetOriginFinished() {
		m.t.Fatal("Expected call to NodeNetworkMock.GetOrigin")
	}

	if !m.GetStateFinished() {
		m.t.Fatal("Expected call to NodeNetworkMock.GetState")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *NodeNetworkMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *NodeNetworkMock) MinimockWait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	for {
		ok := true
		ok = ok && m.GetActiveNodeFinished()
		ok = ok && m.GetActiveNodesFinished()
		ok = ok && m.GetActiveNodesByRoleFinished()
		ok = ok && m.GetOriginFinished()
		ok = ok && m.GetStateFinished()

		if ok {
			return
		}

		select {
		case <-timeoutCh:

			if !m.GetActiveNodeFinished() {
				m.t.Error("Expected call to NodeNetworkMock.GetActiveNode")
			}

			if !m.GetActiveNodesFinished() {
				m.t.Error("Expected call to NodeNetworkMock.GetActiveNodes")
			}

			if !m.GetActiveNodesByRoleFinished() {
				m.t.Error("Expected call to NodeNetworkMock.GetActiveNodesByRole")
			}

			if !m.GetOriginFinished() {
				m.t.Error("Expected call to NodeNetworkMock.GetOrigin")
			}

			if !m.GetStateFinished() {
				m.t.Error("Expected call to NodeNetworkMock.GetState")
			}

			m.t.Fatalf("Some mocks were not called on time: %s", timeout)
			return
		default:
			time.Sleep(time.Millisecond)
		}
	}
}

//AllMocksCalled returns true if all mocked methods were called before the execution of AllMocksCalled,
//it can be used with assert/require, i.e. assert.True(mock.AllMocksCalled())
func (m *NodeNetworkMock) AllMocksCalled() bool {

	if !m.GetActiveNodeFinished() {
		return false
	}

	if !m.GetActiveNodesFinished() {
		return false
	}

	if !m.GetActiveNodesByRoleFinished() {
		return false
	}

	if !m.GetOriginFinished() {
		return false
	}

	if !m.GetStateFinished() {
		return false
	}

	return true
}
