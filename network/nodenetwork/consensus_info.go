/*
 * The Clear BSD License
 *
 * Copyright (c) 2019 Insolar Technologies
 *
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without modification,
 * are permitted (subject to the limitations in the disclaimer below) provided that
 * the following conditions are met:
 *
 *  * Redistributions of source code must retain the above copyright notice,
 *    this list of conditions and the following disclaimer.
 *  * Redistributions in binary form must reproduce the above copyright notice,
 *    this list of conditions and the following disclaimer in the documentation
 *    and/or other materials provided with the distribution.
 *  * Neither the name of Insolar Technologies nor the names of its contributors
 *    may be used to endorse or promote products derived from this software
 *    without specific prior written permission.
 *
 * NO EXPRESS OR IMPLIED LICENSES TO ANY PARTY'S PATENT RIGHTS ARE GRANTED
 * BY THIS LICENSE. THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND
 * CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING,
 * BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS
 * FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 * HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 * SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 * LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 * DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 * THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

package nodenetwork

import (
	"github.com/insolar/insolar/network/node"
	"sync"

	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/log"
	"github.com/insolar/insolar/network/transport/host"
	"github.com/pkg/errors"
)

type consensusInfo struct {
	lock                       sync.RWMutex
	tempMapR                   map[core.RecordRef]*host.Host
	tempMapS                   map[core.ShortNodeID]*host.Host
	nodesJoinedDuringPrevPulse bool
	isJoiner                   bool
}

func (ci *consensusInfo) SetIsJoiner(isJoiner bool) {
	ci.lock.Lock()
	defer ci.lock.Unlock()

	ci.isJoiner = isJoiner
}

func (ci *consensusInfo) IsJoiner() bool {
	ci.lock.RLock()
	defer ci.lock.RUnlock()

	return ci.isJoiner
}

func (ci *consensusInfo) NodesJoinedDuringPreviousPulse() bool {
	ci.lock.RLock()
	defer ci.lock.RUnlock()

	return ci.nodesJoinedDuringPrevPulse
}

func (ci *consensusInfo) AddTemporaryMapping(nodeID core.RecordRef, shortID core.ShortNodeID, address string) error {
	consensusAddress, err := node.IncrementPort(address)
	if err != nil {
		return errors.Wrapf(err, "Failed to increment port for address %s", address)
	}
	h, err := host.NewHostNS(consensusAddress, nodeID, shortID)
	if err != nil {
		return errors.Wrapf(err, "Failed to generate address (%s, %s, %d)", consensusAddress, nodeID, shortID)
	}
	ci.lock.Lock()
	ci.tempMapR[nodeID] = h
	ci.tempMapS[shortID] = h
	ci.lock.Unlock()
	log.Infof("Added temporary mapping: %s -> (%s, %d)", consensusAddress, nodeID, shortID)
	return nil
}

func (ci *consensusInfo) ResolveConsensus(shortID core.ShortNodeID) *host.Host {
	ci.lock.RLock()
	defer ci.lock.RUnlock()

	return ci.tempMapS[shortID]
}

func (ci *consensusInfo) ResolveConsensusRef(nodeID core.RecordRef) *host.Host {
	ci.lock.RLock()
	defer ci.lock.RUnlock()

	return ci.tempMapR[nodeID]
}

func (ci *consensusInfo) flush(nodesJoinedDuringPrevPulse bool) {
	ci.lock.Lock()
	defer ci.lock.Unlock()

	ci.tempMapR = make(map[core.RecordRef]*host.Host)
	ci.tempMapS = make(map[core.ShortNodeID]*host.Host)
	ci.nodesJoinedDuringPrevPulse = nodesJoinedDuringPrevPulse
}

func newConsensusInfo() *consensusInfo {
	return &consensusInfo{
		tempMapR: make(map[core.RecordRef]*host.Host),
		tempMapS: make(map[core.ShortNodeID]*host.Host),
	}
}
