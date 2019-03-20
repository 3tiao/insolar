/*
 *    Copyright 2019 Insolar Technologies
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

package adapter

import (
	"fmt"
	"time"

	"github.com/insolar/insolar/log"
	"github.com/pkg/errors"
)

// NewWaitAdapter creates new instance of SimpleWaitAdapter with Waiter as worker
func NewWaitAdapter() PulseConveyorAdapterTaskSink {
	return NewAdapterWithQueue(NewWaiter())
}

// WaiterTask is task for adapter for waiting
type WaiterTask struct {
	waitPeriodMilliseconds int
}

// Waiter is worker for adapter for waiting
type Waiter struct{}

// NewWaiter returns new instance of worker which waiting
func NewWaiter() Processor {
	return &Waiter{}
}

// Process implements Processor interface
func (w *Waiter) Process(adapterID uint32, task AdapterTask, cancelInfo CancelInfo) Events {
	log.Info("[ Waiter.Process ] Start. cancelInfo.id: ", cancelInfo.ID())

	payload, ok := task.taskPayload.(WaiterTask)
	var msg interface{}

	if !ok {
		msg = errors.Errorf("[ Waiter.Process ] Incorrect payload type: %T", task.taskPayload)
		return Events{RespPayload: msg}
	}

	select {
	case <-cancelInfo.Cancel():
		log.Info("[ Waiter.Process ] Cancel. Return Nil as Response")
		msg = nil
	case <-cancelInfo.Flush():
		log.Info("[ Waiter.Process ] Flush. DON'T Return Response")
		return Events{Flushed: true}
	case <-time.After(time.Duration(payload.waitPeriodMilliseconds) * time.Millisecond):
		msg = fmt.Sprintf("Work completed successfully. Waited %d millisecond", payload.waitPeriodMilliseconds)
	}

	log.Info("[ Waiter.Process ] ", msg)

	return Events{RespPayload: msg}
	// TODO: remove cancelInfo from swa.taskHolder
}