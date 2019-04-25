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

	watermillMsg "github.com/ThreeDotsLabs/watermill/message"
	"github.com/insolar/insolar/insolar/flow"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/instrumentation/instracer"
)

type GetLedgerPendingRequest struct {
	dep *Dependencies

	Message *watermillMsg.Message
}

func (p *GetLedgerPendingRequest) Present(ctx context.Context, f flow.Flow) error {
	ctx, span := instracer.StartSpan(ctx, "LogicRunner.getLedgerPendingRequest")
	defer span.End()

	lr := p.dep.lr
	es := lr.getExecStateFromRef(ctx, p.Message.Payload)
	if es == nil {
		return nil
	}

	es.getLedgerPendingMutex.Lock()
	defer es.getLedgerPendingMutex.Unlock()

	ref := lr.unsafeGetLedgerPendingRequest(ctx, es)
	if ref == nil {
		return nil
	}

	insolarRef := Ref{}.FromSlice(p.Message.Payload)
	h := StartQueueProcessorIfNeeded{
		es:  es,
		dep: p.dep,
		ref: &insolarRef,
	}
	if err := f.Handle(ctx, h.Present); err != nil {
		inslogger.FromContext(ctx).Warn("[ GetLedgerPendingRequest.Present ] Couldn't start queue: ", err)
	}

	return nil
}