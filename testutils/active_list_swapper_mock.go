package testutils

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.9
The original interface "ActiveListSwapper" can be found in github.com/insolar/insolar/ledger/pulsemanager
*/
import (
	context "context"
	"sync/atomic"
	"time"

	"github.com/gojuno/minimock"
	insolar "github.com/insolar/insolar/insolar"

	testify_assert "github.com/stretchr/testify/assert"
)

//ActiveListSwapperMock implements github.com/insolar/insolar/ledger/pulsemanager.ActiveListSwapper
type ActiveListSwapperMock struct {
	t minimock.Tester

	MoveSyncToActiveFunc       func(p context.Context, p1 insolar.PulseNumber) (r error)
	MoveSyncToActiveCounter    uint64
	MoveSyncToActivePreCounter uint64
	MoveSyncToActiveMock       mActiveListSwapperMockMoveSyncToActive
}

//NewActiveListSwapperMock returns a mock for github.com/insolar/insolar/ledger/pulsemanager.ActiveListSwapper
func NewActiveListSwapperMock(t minimock.Tester) *ActiveListSwapperMock {
	m := &ActiveListSwapperMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.MoveSyncToActiveMock = mActiveListSwapperMockMoveSyncToActive{mock: m}

	return m
}

type mActiveListSwapperMockMoveSyncToActive struct {
	mock              *ActiveListSwapperMock
	mainExpectation   *ActiveListSwapperMockMoveSyncToActiveExpectation
	expectationSeries []*ActiveListSwapperMockMoveSyncToActiveExpectation
}

type ActiveListSwapperMockMoveSyncToActiveExpectation struct {
	input  *ActiveListSwapperMockMoveSyncToActiveInput
	result *ActiveListSwapperMockMoveSyncToActiveResult
}

type ActiveListSwapperMockMoveSyncToActiveInput struct {
	p  context.Context
	p1 insolar.PulseNumber
}

type ActiveListSwapperMockMoveSyncToActiveResult struct {
	r error
}

//Expect specifies that invocation of ActiveListSwapper.MoveSyncToActive is expected from 1 to Infinity times
func (m *mActiveListSwapperMockMoveSyncToActive) Expect(p context.Context, p1 insolar.PulseNumber) *mActiveListSwapperMockMoveSyncToActive {
	m.mock.MoveSyncToActiveFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &ActiveListSwapperMockMoveSyncToActiveExpectation{}
	}
	m.mainExpectation.input = &ActiveListSwapperMockMoveSyncToActiveInput{p, p1}
	return m
}

//Return specifies results of invocation of ActiveListSwapper.MoveSyncToActive
func (m *mActiveListSwapperMockMoveSyncToActive) Return(r error) *ActiveListSwapperMock {
	m.mock.MoveSyncToActiveFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &ActiveListSwapperMockMoveSyncToActiveExpectation{}
	}
	m.mainExpectation.result = &ActiveListSwapperMockMoveSyncToActiveResult{r}
	return m.mock
}

//ExpectOnce specifies that invocation of ActiveListSwapper.MoveSyncToActive is expected once
func (m *mActiveListSwapperMockMoveSyncToActive) ExpectOnce(p context.Context, p1 insolar.PulseNumber) *ActiveListSwapperMockMoveSyncToActiveExpectation {
	m.mock.MoveSyncToActiveFunc = nil
	m.mainExpectation = nil

	expectation := &ActiveListSwapperMockMoveSyncToActiveExpectation{}
	expectation.input = &ActiveListSwapperMockMoveSyncToActiveInput{p, p1}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *ActiveListSwapperMockMoveSyncToActiveExpectation) Return(r error) {
	e.result = &ActiveListSwapperMockMoveSyncToActiveResult{r}
}

//Set uses given function f as a mock of ActiveListSwapper.MoveSyncToActive method
func (m *mActiveListSwapperMockMoveSyncToActive) Set(f func(p context.Context, p1 insolar.PulseNumber) (r error)) *ActiveListSwapperMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.MoveSyncToActiveFunc = f
	return m.mock
}

//MoveSyncToActive implements github.com/insolar/insolar/ledger/pulsemanager.ActiveListSwapper interface
func (m *ActiveListSwapperMock) MoveSyncToActive(p context.Context, p1 insolar.PulseNumber) (r error) {
	counter := atomic.AddUint64(&m.MoveSyncToActivePreCounter, 1)
	defer atomic.AddUint64(&m.MoveSyncToActiveCounter, 1)

	if len(m.MoveSyncToActiveMock.expectationSeries) > 0 {
		if counter > uint64(len(m.MoveSyncToActiveMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to ActiveListSwapperMock.MoveSyncToActive. %v %v", p, p1)
			return
		}

		input := m.MoveSyncToActiveMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, ActiveListSwapperMockMoveSyncToActiveInput{p, p1}, "ActiveListSwapper.MoveSyncToActive got unexpected parameters")

		result := m.MoveSyncToActiveMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the ActiveListSwapperMock.MoveSyncToActive")
			return
		}

		r = result.r

		return
	}

	if m.MoveSyncToActiveMock.mainExpectation != nil {

		input := m.MoveSyncToActiveMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, ActiveListSwapperMockMoveSyncToActiveInput{p, p1}, "ActiveListSwapper.MoveSyncToActive got unexpected parameters")
		}

		result := m.MoveSyncToActiveMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the ActiveListSwapperMock.MoveSyncToActive")
		}

		r = result.r

		return
	}

	if m.MoveSyncToActiveFunc == nil {
		m.t.Fatalf("Unexpected call to ActiveListSwapperMock.MoveSyncToActive. %v %v", p, p1)
		return
	}

	return m.MoveSyncToActiveFunc(p, p1)
}

//MoveSyncToActiveMinimockCounter returns a count of ActiveListSwapperMock.MoveSyncToActiveFunc invocations
func (m *ActiveListSwapperMock) MoveSyncToActiveMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.MoveSyncToActiveCounter)
}

//MoveSyncToActiveMinimockPreCounter returns the value of ActiveListSwapperMock.MoveSyncToActive invocations
func (m *ActiveListSwapperMock) MoveSyncToActiveMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.MoveSyncToActivePreCounter)
}

//MoveSyncToActiveFinished returns true if mock invocations count is ok
func (m *ActiveListSwapperMock) MoveSyncToActiveFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.MoveSyncToActiveMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.MoveSyncToActiveCounter) == uint64(len(m.MoveSyncToActiveMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.MoveSyncToActiveMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.MoveSyncToActiveCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.MoveSyncToActiveFunc != nil {
		return atomic.LoadUint64(&m.MoveSyncToActiveCounter) > 0
	}

	return true
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *ActiveListSwapperMock) ValidateCallCounters() {

	if !m.MoveSyncToActiveFinished() {
		m.t.Fatal("Expected call to ActiveListSwapperMock.MoveSyncToActive")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *ActiveListSwapperMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *ActiveListSwapperMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *ActiveListSwapperMock) MinimockFinish() {

	if !m.MoveSyncToActiveFinished() {
		m.t.Fatal("Expected call to ActiveListSwapperMock.MoveSyncToActive")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *ActiveListSwapperMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *ActiveListSwapperMock) MinimockWait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	for {
		ok := true
		ok = ok && m.MoveSyncToActiveFinished()

		if ok {
			return
		}

		select {
		case <-timeoutCh:

			if !m.MoveSyncToActiveFinished() {
				m.t.Error("Expected call to ActiveListSwapperMock.MoveSyncToActive")
			}

			m.t.Fatalf("Some mocks were not called on time: %s", timeout)
			return
		default:
			time.Sleep(time.Millisecond)
		}
	}
}

//AllMocksCalled returns true if all mocked methods were called before the execution of AllMocksCalled,
//it can be used with assert/require, i.e. assert.True(mock.AllMocksCalled())
func (m *ActiveListSwapperMock) AllMocksCalled() bool {

	if !m.MoveSyncToActiveFinished() {
		return false
	}

	return true
}
