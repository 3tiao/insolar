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

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/flow"
	"github.com/insolar/insolar/insolar/message"
	"github.com/insolar/insolar/insolar/reply"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/ledger/recentstorage"
	"github.com/insolar/insolar/ledger/storage"
	"github.com/insolar/insolar/ledger/storage/object"
	"github.com/pkg/errors"
)

type GetIndex struct {
	Object insolar.Reference
	Jet    insolar.JetID

	Res struct {
		Index object.Lifeline
	}

	Dep struct {
		Recent      recentstorage.Provider
		Locker      storage.IDLocker
		Storage     object.IndexStorage
		Coordinator insolar.JetCoordinator
		Bus         insolar.MessageBus
	}
}

func (p *GetIndex) Proceed(ctx context.Context) error {
	objectID := *p.Object.Record()
	logger := inslogger.FromContext(ctx)
	p.Dep.Recent.GetIndexStorage(ctx, insolar.ID(p.Jet)).AddObject(ctx, objectID)

	p.Dep.Locker.Lock(&objectID)
	defer p.Dep.Locker.Unlock(&objectID)

	idx, err := p.Dep.Storage.ForID(ctx, objectID)
	if err == nil {
		p.Res.Index = idx
		return nil
	}
	if err != object.ErrIndexNotFound {
		return errors.Wrap(err, "failed to fetch index")
	}

	logger.Debug("failed to fetch index (fetching from heavy)")
	heavy, err := p.Dep.Coordinator.Heavy(ctx, flow.Pulse(ctx))
	if err != nil {
		return errors.Wrap(err, "failed to calculate heavy")
	}
	genericReply, err := p.Dep.Bus.Send(ctx, &message.GetObjectIndex{
		Object: p.Object,
	}, &insolar.MessageSendOptions{
		Receiver: heavy,
	})
	if err != nil {
		return errors.Wrap(err, "failed to fetch index from heavy")
	}
	rep, ok := genericReply.(*reply.ObjectIndex)
	if !ok {
		return fmt.Errorf("failed to fetch index from heavy: unexpected reply type %T", genericReply)
	}

	p.Res.Index, err = object.DecodeIndex(rep.Index)
	if err != nil {
		return errors.Wrap(err, "failed to decode index")
	}

	p.Res.Index.JetID = p.Jet
	err = p.Dep.Storage.Set(ctx, objectID, p.Res.Index)
	if err != nil {
		return errors.Wrap(err, "failed to save index")
	}

	return nil
}
