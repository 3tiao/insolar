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

package messagebus

import (
	"errors"
	"sync/atomic"
	"time"

	"github.com/insolar/insolar/core"
)

var (
	// ErrFutureTimeout is returned when the operation timeout is exceeded.
	ErrFutureTimeout = errors.New("can't wait for result: timeout")
	// ErrFutureChannelClosed is returned when the input channel is closed.
	ErrFutureChannelClosed = errors.New("can't wait for result: channel closed")
)

type future struct {
	result   chan core.Reply
	finished uint64
}

// NewFuture creates new ConveyorFuture.
func NewFuture() core.ConveyorFuture {
	return &future{
		result: make(chan core.Reply, 1),
	}
}

// Result returns result packet channel.
func (future *future) Result() <-chan core.Reply {
	return future.result
}

// SetResult write packet to the result channel.
func (future *future) SetResult(res core.Reply) {
	if atomic.CompareAndSwapUint64(&future.finished, 0, 1) {
		future.result <- res
		future.finish()
	}
}

// GetResult gets the future result from Result() channel with a timeout set to `duration`.
func (future *future) GetResult(duration time.Duration) (core.Reply, error) {
	select {
	case result, ok := <-future.Result():
		if !ok {
			return nil, ErrFutureChannelClosed
		}
		return result, nil
	case <-time.After(duration):
		future.Cancel()
		return nil, ErrFutureTimeout
	}
}

// Cancel allows to cancel ConveyorFuture processing.
func (future *future) Cancel() {
	if atomic.CompareAndSwapUint64(&future.finished, 0, 1) {
		future.finish()
	}
}

func (future *future) finish() {
	close(future.result)
}

// ConveyorPendingMessage is message for conveyor which can pend for response
type ConveyorPendingMessage struct {
	Msg    core.Parcel
	Future core.ConveyorFuture
}