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

package heavyclient

import (
	"context"
	"sync"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/ledger/recentstorage"
	"github.com/insolar/insolar/ledger/storage"
	"github.com/insolar/insolar/ledger/storage/blob"
	"github.com/insolar/insolar/ledger/storage/drop"
	"github.com/insolar/insolar/ledger/storage/object"
	"github.com/insolar/insolar/ledger/storage/pulse"
	"golang.org/x/sync/singleflight"
)

// Pool manages state of heavy sync clients (one client per jet id).
type Pool struct {
	bus                   insolar.MessageBus
	pulseAccessor         pulse.Accessor
	pulseCalculator       pulse.Calculator
	dropAccessor          drop.Accessor
	blobSyncAccessor      blob.CollectionAccessor
	recSyncAccessor       object.RecordCollectionAccessor
	replicaStorage        storage.ReplicaStorage
	idxCollectionAccessor object.IndexCollectionAccessor
	idxCleaner            object.IndexCleaner
	db                    storage.DBContext

	clientDefaults Options

	sync.Mutex
	clients map[insolar.ID]*JetClient

	cleanupGroup singleflight.Group
}

// NewPool constructor of new pool.
func NewPool(
	bus insolar.MessageBus,
	pulseAccessor pulse.Accessor,
	pulseCalculator pulse.Calculator,
	replicaStorage storage.ReplicaStorage,
	dropAccessor drop.Accessor,
	blobSyncAccessor blob.CollectionAccessor,
	recSyncAccessor object.RecordCollectionAccessor,
	idxCollectionAccessor object.IndexCollectionAccessor,
	idxCleaner object.IndexCleaner,
	db storage.DBContext,
	clientDefaults Options,
) *Pool {
	return &Pool{
		bus:                   bus,
		dropAccessor:          dropAccessor,
		blobSyncAccessor:      blobSyncAccessor,
		recSyncAccessor:       recSyncAccessor,
		pulseAccessor:         pulseAccessor,
		pulseCalculator:       pulseCalculator,
		replicaStorage:        replicaStorage,
		clientDefaults:        clientDefaults,
		idxCollectionAccessor: idxCollectionAccessor,
		idxCleaner:            idxCleaner,
		db:                    db,
		clients:               map[insolar.ID]*JetClient{},
	}
}

// Stop send stop signals to all managed heavy clients and waits when until all of them will stop.
func (scp *Pool) Stop(ctx context.Context) {
	scp.Lock()
	defer scp.Unlock()

	var wg sync.WaitGroup
	wg.Add(len(scp.clients))
	for _, c := range scp.clients {
		c := c
		go func() {
			c.Stop(ctx)
			wg.Done()
		}()
	}
	wg.Wait()
}

// AddPulsesToSyncClient add pulse numbers to the end of jet's heavy client queue.
//
// Bool flag 'shouldrun' controls should heavy client be started (if not already) or not.
func (scp *Pool) AddPulsesToSyncClient(
	ctx context.Context,
	jetID insolar.ID,
	shouldrun bool,
	pns ...insolar.PulseNumber,
) *JetClient {
	scp.Lock()
	client, ok := scp.clients[jetID]
	if !ok {
		client = NewJetClient(
			scp.replicaStorage,
			scp.bus,
			scp.pulseAccessor,
			scp.pulseCalculator,
			scp.dropAccessor,
			scp.blobSyncAccessor,
			scp.recSyncAccessor,
			scp.idxCollectionAccessor,
			scp.idxCleaner,
			scp.db,
			jetID,
			scp.clientDefaults,
		)

		scp.clients[jetID] = client
	}
	scp.Unlock()

	client.addPulses(ctx, pns)

	if shouldrun {
		client.runOnce(ctx)
		if len(client.signal) == 0 {
			// send signal we have new pulse
			client.signal <- struct{}{}
		}
	}
	return client
}

// AllClients returns slice with all clients in Pool.
func (scp *Pool) AllClients(ctx context.Context) []*JetClient {
	scp.Lock()
	defer scp.Unlock()
	clients := make([]*JetClient, 0, len(scp.clients))
	for _, c := range scp.clients {
		clients = append(clients, c)
	}
	return clients
}

// LightCleanup starts async cleanup on all heavy synchronization clients (per jet cleanup).
//
// Waits until all cleanup will done and mesaures time.
//
// Under hood it uses singleflight on Jet prefix to avoid clashing on the same key space.
func (scp *Pool) LightCleanup(
	ctx context.Context,
	untilPN insolar.PulseNumber,
	rsp recentstorage.Provider,
	jetIndexesRemoved map[insolar.ID][]insolar.ID,
) error {
	// inslog := inslogger.FromContext(ctx)
	// start := time.Now()
	// defer func() {
	// 	latency := time.Since(start)
	// 	inslog.Infof("cleanLightData db clean phase time spend=%v", latency)
	// 	stats.Record(ctx, statCleanLatencyDB.M(latency.Nanoseconds()/1e6))
	// }()
	//
	// func() {
	// 	startCleanup := time.Now()
	// 	defer func() {
	// 		latency := time.Since(startCleanup)
	// 		inslog.Infof("cleanLightData db clean phase job time spend=%v", latency)
	// 	}()
	//
	// 	// This is how we can get all jets served on light during it storage lifetime.
	// 	// jets, err := scp.db.GetAllSyncClientJets(ctx)
	//
	// 	allClients := scp.AllClients(ctx)
	// 	var wg sync.WaitGroup
	//
	// 	cleanupConcurrency := 8
	// 	sem := make(chan struct{}, cleanupConcurrency)
	//
	// 	jetPrefixSeen := map[string]struct{}{}
	//
	// 	for _, c := range allClients {
	// 		jetID := c.jetID
	// 		jetPrefix := insolar.JetID(jetID).Prefix()
	// 		prefixKey := string(jetPrefix)
	//
	// 		_, skipRecordsCleanup := jetPrefixSeen[prefixKey]
	// 		jetPrefixSeen[prefixKey] = struct{}{}
	//
	// 		candidates := jetIndexesRemoved[insolar.ID(jetID)]
	//
	// 		if (len(candidates) == 0) && skipRecordsCleanup {
	// 			continue
	// 		}
	//
	// 		wg.Add(1)
	// 		sem <- struct{}{}
	// 		go func() {
	// 			defer func() {
	// 				wg.Done()
	// 				<-sem
	// 			}()
	// 			_, _, _ = scp.cleanupGroup.Do(string(jetPrefix), func() (interface{}, error) {
	//
	// 				inslogger.FromContext(ctx).Debugf("Start light cleanup, pulse < %v, jet = %v",
	// 					untilPN, jetID.DebugString())
	//
	// 				if len(candidates) > 0 {
	// 					jetRecentStore := rsp.GetIndexStorage(ctx, insolar.ID(jetID))
	// 				}
	//
	// 				if skipRecordsCleanup {
	// 					return nil, nil
	// 				}
	//
	// 				recsRmStat, err := scp.cleaner.CleanJetRecordsUntilPulse(ctx, insolar.ID(jetID), untilPN)
	// 				if err != nil {
	// 					inslogger.FromContext(ctx).Errorf("Error on light cleanup (pulse < %v, jet = %v): %v",
	// 						untilPN, jetID.DebugString(), err)
	// 					return nil, nil
	// 				}
	// 				inslogger.FromContext(ctx).Infof(
	// 					"Records light cleanup, records stat=%#v (pulse < %v, jet = %v)", recsRmStat, untilPN, jetID.DebugString())
	// 				return nil, nil
	// 			})
	// 		}()
	// 	}
	// 	wg.Wait()
	// }()
	return nil
}
