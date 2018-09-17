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

package configuration

import (
	"github.com/insolar/insolar/core"
)

type Storage struct {
	// DataDirectory is a directory where database's files live.
	DataDirectory string
	// TxRetriesOnConflict defines how many retries on transaction conflicts
	// storage update methods should do.
	TxRetriesOnConflict int
}

// JetCoordinator holds configuration for JetCoordinator.
type JetCoordinator struct {
	RoleCandidates map[int][]string
	RoleCounts     map[int]int
}

// Ledger holds configuration for ledger.
type Ledger struct {
	// Storage defines storage configuration.
	Storage Storage
	// JetCoordinator defines jet coordinator configuration.
	JetCoordinator JetCoordinator
}

// NewLedger creates new default Ledger configuration.
func NewLedger() Ledger {
	return Ledger{
		Storage: Storage{
			DataDirectory:       "./data",
			TxRetriesOnConflict: 3,
		},

		JetCoordinator: JetCoordinator{
			RoleCandidates: map[int][]string{
				int(core.RoleVirtualExecutor):  {"ve1", "ve2"},
				int(core.RoleHeavyExecutor):    {"he1", "he2"},
				int(core.RoleLightExecutor):    {"le1", "le2"},
				int(core.RoleVirtualValidator): {"vv1", "vv2"},
				int(core.RoleLightValidator):   {"lv1", "lv2"},
			},
			RoleCounts: map[int]int{
				int(core.RoleVirtualExecutor):  1,
				int(core.RoleHeavyExecutor):    1,
				int(core.RoleLightExecutor):    1,
				int(core.RoleVirtualValidator): 3,
				int(core.RoleLightValidator):   3,
			},
		},
	}
}
