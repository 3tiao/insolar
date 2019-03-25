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

package object

var registry = map[TypeID]VirtualRecord{}

// Register makes provided record serializable. Should be called for each record in init().
func register(id TypeID, r VirtualRecord) {
	if _, ok := registry[id]; ok {
		panic("duplicate record type")
	}

	registry[id] = r
}

// Registered returns records by type.
func Registered() map[TypeID]VirtualRecord {
	res := map[TypeID]VirtualRecord{}
	for id, rec := range registry {
		res[id] = rec
	}
	return res
}
