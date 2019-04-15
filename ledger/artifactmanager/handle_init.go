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

package artifactmanager

import (
	"context"
	"fmt"

	"github.com/insolar/insolar/insolar/flow"
	"github.com/insolar/insolar/insolar/flow/bus"
	"github.com/pkg/errors"
)

type Init struct {
	dep *Dependencies

	Message bus.Message
}

func (s *Init) Future(ctx context.Context, f flow.Flow) error {
	return f.Migrate(ctx, s.Present)
}

func (s *Init) Present(ctx context.Context, f flow.Flow) error {
	// switch s.Message.Parcel.Message().Type() {
	// case insolar.TypeGetObject:
	// 	h := &GetObject{
	// 		dep:     s.dep,
	// 		Message: s.Message,
	// 	}
	// 	return f.Handle(ctx, h.Present)
	// default:
	// 	return fmt.Errorf("no handler for message type %s", s.Message.Parcel.Message().Type().String())
	// }
	fmt.Println("sorry love, type is", s.Message.Msg.Metadata.Get("Type"))
	switch s.Message.Msg.Metadata.Get("Type") {
	case "TypeGetObject":
		h := &GetObject{
			dep:     s.dep,
			Message: s.Message,
		}
		fmt.Println("all well love")
		return f.Handle(ctx, h.Present)
	default:
		fmt.Println("sorry love, msg is", s.Message)
		return fmt.Errorf("no handler for message type %s", s.Message.Parcel.Message().Type().String())
	}
}

func (s *Init) Past(ctx context.Context, f flow.Flow) error {
	return f.Procedure(ctx, &ReturnReply{ReplyTo: s.Message.ReplyTo, Err: errors.New("no past handler")})
}
