/*
 *    Copyright 2019 Insolar
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

package gen

import (
	fuzz "github.com/google/gofuzz"

	"github.com/insolar/insolar/core"
)

// ID generates random id.
func ID() (id core.RecordID) {
	fuzz.New().NilChance(0).Fuzz(&id)
	return
}

// JetID generates random jet id.
func JetID() (jetID core.JetID) {
	f := fuzz.New().Funcs(func(id *core.JetID, c fuzz.Continue) {
		c.Fuzz(id)
		// set special pulse number
		copy(id[:core.PulseNumberSize], core.PulseNumberJet.Bytes())
		// set depth
		// adds 1 because Intn returns [0,n)
		id[core.PulseNumberSize] = byte(c.Intn(core.JetMaximumDepth + 1))
	})
	f.Fuzz(&jetID)
	return
}

// Reference generates random reference.
func Reference() (ref core.RecordRef) {
	fuzz.New().NilChance(0).Fuzz(&ref)
	return
}
