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

package logicrunner

import (
	"context"
	"fmt"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/flow"
	"github.com/insolar/insolar/insolar/flow/bus"
)

const InnerMsgTopic = "InnerMsg"

type Dependencies struct {
	Publisher message.Publisher
}

type Init struct {
	dep *Dependencies

	Message bus.Message
}

func (s *Init) Present(ctx context.Context, f flow.Flow) error {
	switch s.Message.Parcel.Message().Type() {
	case insolar.TypeCallMethod:
		return nil
	default:
		return fmt.Errorf("no handler for message type %s", s.Message.Parcel.Message().Type().String())
	}
}

type InitInner struct {
	dep *Dependencies

	Message message.Message
}

func (s *InitInner) Present(ctx context.Context, f flow.Flow) error {
	switch s.Message.Metadata.Get("Type") {
	default:
		return fmt.Errorf("no handler for message type %s", s.Message.Metadata.Get("Type"))
	}
}
