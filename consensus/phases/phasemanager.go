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

package phases

import (
	"context"
	"time"

	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/network"
	"github.com/insolar/insolar/network/merkle"
	"github.com/pkg/errors"
)

type PhaseManager interface {
	OnPulse(ctx context.Context, pulse *core.Pulse) error
}

type Phases struct {
	FirstPhase  FirstPhase  `inject:"subcomponent"`
	SecondPhase SecondPhase `inject:"subcomponent"`
	ThirdPhase  ThirdPhase  `inject:"subcomponent"`

	PulseManager core.PulseManager  `inject:""`
	NodeKeeper   network.NodeKeeper `inject:""`
	Calculator   merkle.Calculator  `inject:""`
}

// NewPhaseManager creates and returns a new phase manager.
func NewPhaseManager() PhaseManager {
	return &Phases{}
}

// OnPulse starts calculate args on phases.
func (pm *Phases) OnPulse(ctx context.Context, pulse *core.Pulse) error {
	var err error

	pulseDuration, err := getPulseDuration(pulse)
	if err != nil {
		return errors.Wrap(err, "[ OnPulse ] Failed to get pulse duration")
	}

	var tctx context.Context
	var cancel context.CancelFunc

	tctx, cancel = contextTimeout(ctx, *pulseDuration, 0.2)
	defer cancel()

	firstPhaseState, err := pm.FirstPhase.Execute(tctx, pulse)
	if err != nil {
		return errors.Wrap(err, "Network consensus: error executing phase 1")
	}

	tctx, cancel = contextTimeout(ctx, *pulseDuration, 0.2)
	defer cancel()

	secondPhaseState, err := pm.SecondPhase.Execute(tctx, firstPhaseState)
	if err != nil {
		return errors.Wrap(err, "Network consensus: error executing phase 2.0")
	}

	tctx, cancel = contextTimeout(ctx, *pulseDuration, 0.2)
	defer cancel()

	secondPhaseState, err = pm.SecondPhase.Execute21(tctx, secondPhaseState)
	if err != nil {
		return errors.Wrap(err, "Network consensus: error executing phase 2.1")
	}

	tctx, cancel = contextTimeout(ctx, *pulseDuration, 0.2)
	defer cancel()

	thirdPhaseState, err := pm.ThirdPhase.Execute(tctx, secondPhaseState)
	if err != nil {
		return errors.Wrap(err, "Network consensus: error executing phase 3")
	}

	state := thirdPhaseState
	cloud := &merkle.CloudEntry{
		ProofSet:      []*merkle.GlobuleProof{state.GlobuleProof},
		PrevCloudHash: pm.NodeKeeper.GetCloudHash(),
	}
	hash, _, err := pm.Calculator.GetCloudProof(cloud)
	if err != nil {
		return errors.Wrap(err, "Network consensus: error calculating cloud hash")
	}
	pm.NodeKeeper.SetCloudHash(hash)

	state.UnsyncList.ApproveSync(state.ActiveNodes)
	pm.NodeKeeper.Sync(state.UnsyncList)
	return nil
}

func getPulseDuration(pulse *core.Pulse) (*time.Duration, error) {
	duration := time.Duration(pulse.PulseNumber-pulse.PrevPulseNumber) * time.Second
	return &duration, nil
}

func contextTimeout(ctx context.Context, duration time.Duration, k float64) (context.Context, context.CancelFunc) {
	timeout := time.Duration(k * float64(duration))
	timedCtx, cancelFund := context.WithTimeout(ctx, timeout)
	return timedCtx, cancelFund
}
