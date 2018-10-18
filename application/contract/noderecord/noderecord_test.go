/*
 *    Copyright 2018 Insolar
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package noderecord

import (
	"testing"

	"github.com/insolar/insolar/core"
	"github.com/stretchr/testify/assert"
)

const TestPubKey = "test"
const TestIP = "127.0.0.1"

var TestRoles = []string{"virtual"}

func rolesToStrings(roles []string) []core.NodeRole {
	var result []core.NodeRole
	for _, role := range roles {
		result = append(result, core.GetRoleFromString(role))
	}

	return result
}

func TestNewNodeRecord(t *testing.T) {

	r := rolesToStrings(TestRoles)
	assert.NotEqual(t, core.RoleUnknown, r)
	record := NewNodeRecord(TestPubKey, TestRoles, TestIP)
	assert.Len(t, record.Record.Roles, 1)
	assert.Equal(t, r, record.Record.Roles)
	assert.Equal(t, TestPubKey, record.Record.PublicKey)
}

func TestFromString(t *testing.T) {
	role := core.GetRoleFromString("ZZZ")
	assert.Equal(t, core.RoleUnknown, role)
}

func TestNodeRecord_GetPublicKey(t *testing.T) {
	record := NewNodeRecord(TestPubKey, TestRoles, TestIP)
	assert.Equal(t, TestPubKey, record.GetPublicKey())
}

func TestNodeRecord_GetNodeInfo(t *testing.T) {
	record := NewNodeRecord(TestPubKey, TestRoles, TestIP)
	assert.Equal(t, TestIP, record.Record.IP)
	r := rolesToStrings(TestRoles)
	assert.Equal(t, r, record.Record.Roles)
	assert.Equal(t, TestPubKey, record.Record.PublicKey)
}
