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

package proc

import (
	"context"
	"fmt"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/flow"
	"github.com/insolar/insolar/insolar/flow/bus"
	"github.com/insolar/insolar/insolar/record"
	"github.com/insolar/insolar/insolar/reply"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/ledger/light/recentstorage"
	"github.com/insolar/insolar/ledger/object"
	"github.com/pkg/errors"
)

type SetRecord struct {
	replyTo chan<- bus.Reply
	record  []byte
	jet     insolar.JetID

	Dep struct {
		RecentStorageProvider recentstorage.Provider
		PCS                   insolar.PlatformCryptographyScheme
		RecordModifier        object.RecordModifier
		PendingRequestsLimit  int
	}
}

func NewSetRecord(jetID insolar.JetID, replyTo chan<- bus.Reply, record []byte) *SetRecord {
	return &SetRecord{
		record:  record,
		replyTo: replyTo,
		jet:     jetID,
	}
}

func (p *SetRecord) Proceed(ctx context.Context) error {
	p.replyTo <- p.reply(ctx)
	return nil
}

func (p *SetRecord) reply(ctx context.Context) bus.Reply {
	virtRec := record.Virtual{}
	err := virtRec.Unmarshal(p.record)
	if err != nil {
		return bus.Reply{Err: errors.Wrap(err, "can't deserialize record")}
	}

	hash := record.HashVirtual(p.Dep.PCS.ReferenceHasher(), virtRec)
	calculatedID := insolar.NewID(flow.Pulse(ctx), hash)

	concrete := record.Unwrap(&virtRec)
	switch r := concrete.(type) {
	case *record.Request:
		if p.Dep.RecentStorageProvider.Count() > p.Dep.PendingRequestsLimit {
			return bus.Reply{Reply: &reply.Error{ErrType: reply.ErrTooManyPendingRequests}}
		}
		recentStorage := p.Dep.RecentStorageProvider.GetPendingStorage(ctx, insolar.ID(p.jet))
		recentStorage.AddPendingRequest(ctx, r.GetObject(), *calculatedID)
	case *record.Result:
		recentStorage := p.Dep.RecentStorageProvider.GetPendingStorage(ctx, insolar.ID(p.jet))
		recentStorage.RemovePendingRequest(ctx, r.Object, *r.Request.Record())
	}

	hash = record.HashVirtual(p.Dep.PCS.ReferenceHasher(), virtRec)
	id := insolar.NewID(flow.Pulse(ctx), hash)
	rec := record.Material{
		Virtual: &virtRec,
		JetID:   p.jet,
	}

	err = p.Dep.RecordModifier.Set(ctx, *id, rec)

	if err == object.ErrOverride {
		inslogger.FromContext(ctx).WithField("type", fmt.Sprintf("%T", virtRec)).Warn("set record override")
		id = calculatedID
	} else if err != nil {
		return bus.Reply{Err: errors.Wrap(err, "can't save record into storage")}
	}

	return bus.Reply{Reply: &reply.ID{ID: *id}}
}
