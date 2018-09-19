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

package event

import (
	"io"

	"github.com/insolar/insolar/core"
)

// CallConstructorEvent is a event for calling constructor and obtain its reaction
type CallConstructorEvent struct {
	baseEvent
	ClassRef  core.RecordRef
	Name      string
	Arguments core.Arguments
}

// React handles event and returns associated reaction.
func (e *CallConstructorEvent) React(core.Components) (core.Reaction, error) {
	panic("implement me")
}

// GetOperatingRole returns operating jet role for given event type.
func (e *CallConstructorEvent) GetOperatingRole() core.JetRole {
	return core.RoleVirtualExecutor
}

// Get reference returns referenced object.
func (e *CallConstructorEvent) GetReference() core.RecordRef {
	return e.ClassRef
}

// Serialize serializes event.
func (e *CallConstructorEvent) Serialize() (io.Reader, error) {
	return serialize(e, CallConstructorEventType)
}
