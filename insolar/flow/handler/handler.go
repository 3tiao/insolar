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

package handler

import (
	"context"
	"errors"
	"sync/atomic"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/insolar/insolar/instrumentation/inslogger"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/flow"
	"github.com/insolar/insolar/insolar/flow/bus"
	"github.com/insolar/insolar/insolar/flow/internal/pulse"
	"github.com/insolar/insolar/insolar/flow/internal/thread"
)

type Handler struct {
	handles struct {
		present flow.MakeHandle
		future  flow.MakeHandle
	}
	controller         *thread.Controller
	currentPulseNumber uint32
}

func NewHandler(present flow.MakeHandle, future flow.MakeHandle) *Handler {
	h := &Handler{
		controller: thread.NewController(),
	}
	h.handles.present = present
	h.handles.future = future
	h.currentPulseNumber = insolar.FirstPulseNumber
	return h
}

// ChangePulse is a handle for pulse change vent.
func (h *Handler) ChangePulse(ctx context.Context, pulse insolar.Pulse) {
	h.controller.Pulse()
	atomic.StoreUint32(&h.currentPulseNumber, uint32(pulse.PulseNumber))
}

func (h *Handler) getHandleByPulse(msgPulseNumber insolar.PulseNumber) flow.MakeHandle {
	if uint32(msgPulseNumber) > atomic.LoadUint32(&h.currentPulseNumber) {
		return h.handles.future
	}
	return h.handles.present
}

func (h *Handler) WrapBusHandle(ctx context.Context, parcel insolar.Parcel) (insolar.Reply, error) {
	msg := bus.Message{
		ReplyTo: make(chan bus.Reply, 1),
		Parcel:  parcel,
	}

	ctx = pulse.ContextWith(ctx, parcel.Pulse())

	f := thread.NewThread(msg, h.controller)
	handle := h.getHandleByPulse(parcel.Pulse())

	err := f.Run(ctx, handle(msg))
	var rep bus.Reply
	select {
	case rep = <-msg.ReplyTo:
		return rep.Reply, rep.Err
	default:
	}

	if err != nil {
		return nil, err
	}

	return nil, errors.New("no reply from handler")
}

func (h *Handler) InnerSubscriber(watermillMsg *message.Message) ([]*message.Message, error) {
	msg := bus.Message{
		WatermillMsg: watermillMsg,
	}

	ctx := context.Background()
	ctx = inslogger.ContextWithTrace(ctx, watermillMsg.Metadata.Get("TraceID"))
	logger := inslogger.FromContext(ctx)
	go func() {
		f := thread.NewThread(msg, h.controller)
		err := f.Run(ctx, h.handles.present(msg))
		if err != nil {
			logger.Error("Handling failed", err)
		}
	}()
	return nil, nil
}
