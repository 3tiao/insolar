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

package message

import (
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/platformpolicy"
)

// MethodReturnMode ENUM to set when method returns its result
type MethodReturnMode int

const (
	// ReturnResult - return result as soon as it is ready
	ReturnResult MethodReturnMode = iota
	// ReturnNoWait - call method and return without results
	ReturnNoWait
	// ReturnValidated (not yet) - return result only when it's validated
	// ReturnValidated
)

type PendingState int

const (
	PendingUnknown PendingState = iota
	NotPending
	InPending
)

type IBaseLogicMessage interface {
	insolar.Message
	GetBaseLogicMessage() *BaseLogicMessage
	GetReference() insolar.Reference
	GetCallerPrototype() *insolar.Reference
}

// BaseLogicMessage base of event class family, do not use it standalone
type BaseLogicMessage struct {
	Caller          insolar.Reference
	CallerPrototype insolar.Reference
	Nonce           uint64
	Sequence        uint64
}

func (m *BaseLogicMessage) GetBaseLogicMessage() *BaseLogicMessage {
	return m
}

func (m *BaseLogicMessage) Type() insolar.MessageType {
	panic("Virtual")
}

func (m *BaseLogicMessage) DefaultTarget() *insolar.Reference {
	panic("Virtual")
}

func (m *BaseLogicMessage) DefaultRole() insolar.DynamicRole {
	panic("implement me")
}

func (m *BaseLogicMessage) AllowedSenderObjectAndRole() (*insolar.Reference, insolar.DynamicRole) {
	panic("implement me")
}

func (m *BaseLogicMessage) GetReference() insolar.Reference {
	panic("implement me")
}

func (m *BaseLogicMessage) GetCaller() *insolar.Reference {
	return &m.Caller
}

func (m *BaseLogicMessage) GetCallerPrototype() *insolar.Reference {
	return &m.CallerPrototype
}

// ReturnResults - push results of methods
type ReturnResults struct {
	Target   insolar.Reference
	Caller   insolar.Reference
	Sequence uint64
	Reply    insolar.Reply
	Error    string
}

func (rr *ReturnResults) Type() insolar.MessageType {
	return insolar.TypeReturnResults
}

func (rr *ReturnResults) GetCaller() *insolar.Reference {
	return &rr.Caller
}

func (rr *ReturnResults) DefaultTarget() *insolar.Reference {
	return &rr.Target
}

func (rr *ReturnResults) DefaultRole() insolar.DynamicRole {
	return insolar.DynamicRoleVirtualExecutor
}

func (rr *ReturnResults) AllowedSenderObjectAndRole() (*insolar.Reference, insolar.DynamicRole) {
	return nil, insolar.DynamicRoleVirtualExecutor
}

// CallMethod - Simply call method and return result
type CallMethod struct {
	BaseLogicMessage
	ReturnMode     MethodReturnMode
	Immutable      bool
	ObjectRef      insolar.Reference
	Method         string
	Arguments      insolar.Arguments
	ProxyPrototype insolar.Reference
}

// AllowedSenderObjectAndRole implements interface method
func (cm *CallMethod) AllowedSenderObjectAndRole() (*insolar.Reference, insolar.DynamicRole) {
	c := cm.GetCaller()
	if c.IsEmpty() {
		return nil, 0
	}
	return c, insolar.DynamicRoleVirtualExecutor
}

// DefaultRole returns role for this event
func (*CallMethod) DefaultRole() insolar.DynamicRole {
	return insolar.DynamicRoleVirtualExecutor
}

// DefaultTarget returns of target of this event.
func (cm *CallMethod) DefaultTarget() *insolar.Reference {
	return &cm.ObjectRef
}

func (cm *CallMethod) GetReference() insolar.Reference {
	return cm.ObjectRef
}

// Type returns TypeCallMethod.
func (cm *CallMethod) Type() insolar.MessageType {
	return insolar.TypeCallMethod
}

type SaveAs int

const (
	Child SaveAs = iota
	Delegate
)

// CallConstructor is a message for calling constructor and obtain its reply
type CallConstructor struct {
	BaseLogicMessage
	ParentRef    insolar.Reference
	SaveAs       SaveAs
	PrototypeRef insolar.Reference
	Method       string
	Arguments    insolar.Arguments
	PulseNum     insolar.PulseNumber
}

//
func (cc *CallConstructor) AllowedSenderObjectAndRole() (*insolar.Reference, insolar.DynamicRole) {
	c := cc.GetCaller()
	if c.IsEmpty() {
		return nil, 0
	}
	return c, insolar.DynamicRoleVirtualExecutor
}

// DefaultRole returns role for this event
func (*CallConstructor) DefaultRole() insolar.DynamicRole {
	return insolar.DynamicRoleVirtualExecutor
}

// DefaultTarget returns of target of this event.
func (cc *CallConstructor) DefaultTarget() *insolar.Reference {
	if cc.SaveAs == Delegate {
		return &cc.ParentRef
	}
	return genRequest(cc.PulseNum, MustSerializeBytes(cc), &insolar.DomainID)
}

func (cc *CallConstructor) GetReference() insolar.Reference {
	return *genRequest(cc.PulseNum, MustSerializeBytes(cc), &insolar.DomainID)
}

// Type returns TypeCallConstructor.
func (cc *CallConstructor) Type() insolar.MessageType {
	return insolar.TypeCallConstructor
}

// TODO rename to executorObjectResult (results?)
type ExecutorResults struct {
	Caller                insolar.Reference
	RecordRef             insolar.Reference
	Requests              []CaseBindRequest
	Queue                 []ExecutionQueueElement
	LedgerHasMoreRequests bool
	Pending               PendingState
}

type ExecutionQueueElement struct {
	Parcel  insolar.Parcel
	Request *insolar.Reference
}

// AllowedSenderObjectAndRole implements interface method
func (er *ExecutorResults) AllowedSenderObjectAndRole() (*insolar.Reference, insolar.DynamicRole) {
	// TODO need to think - this message can send only Executor of Previous Pulse, this function
	return nil, 0
}

// DefaultRole returns role for this event
func (er *ExecutorResults) DefaultRole() insolar.DynamicRole {
	return insolar.DynamicRoleVirtualExecutor
}

// DefaultTarget returns of target of this event.
func (er *ExecutorResults) DefaultTarget() *insolar.Reference {
	return &er.RecordRef
}

func (er *ExecutorResults) Type() insolar.MessageType {
	return insolar.TypeExecutorResults
}

// TODO change after changing pulsar
func (er *ExecutorResults) GetCaller() *insolar.Reference {
	return &er.Caller
}

func (er *ExecutorResults) GetReference() insolar.Reference {
	return er.RecordRef
}

type ValidateCaseBind struct {
	Caller    insolar.Reference
	RecordRef insolar.Reference
	Requests  []CaseBindRequest
	Pulse     insolar.Pulse
}

type CaseBindRequest struct {
	Parcel         insolar.Parcel
	Request        insolar.Reference
	MessageBusTape []byte
	Reply          insolar.Reply
	Error          string
}

// AllowedSenderObjectAndRole implements interface method
func (vcb *ValidateCaseBind) AllowedSenderObjectAndRole() (*insolar.Reference, insolar.DynamicRole) {
	return &vcb.RecordRef, insolar.DynamicRoleVirtualExecutor
}

// DefaultRole returns role for this event
func (*ValidateCaseBind) DefaultRole() insolar.DynamicRole {
	return insolar.DynamicRoleVirtualValidator
}

// DefaultTarget returns of target of this event.
func (vcb *ValidateCaseBind) DefaultTarget() *insolar.Reference {
	return &vcb.RecordRef
}

func (vcb *ValidateCaseBind) Type() insolar.MessageType {
	return insolar.TypeValidateCaseBind
}

// TODO change after changing pulsar
func (vcb *ValidateCaseBind) GetCaller() *insolar.Reference {
	return &vcb.Caller // TODO actually it's not right. There is no caller.
}

func (vcb *ValidateCaseBind) GetReference() insolar.Reference {
	return vcb.RecordRef
}

func (vcb *ValidateCaseBind) GetPulse() insolar.Pulse {
	return vcb.Pulse
}

type ValidationResults struct {
	Caller           insolar.Reference
	RecordRef        insolar.Reference
	PassedStepsCount int
	Error            string
}

// AllowedSenderObjectAndRole implements interface method
func (vr *ValidationResults) AllowedSenderObjectAndRole() (*insolar.Reference, insolar.DynamicRole) {
	return &vr.RecordRef, insolar.DynamicRoleVirtualValidator
}

// DefaultRole returns role for this event
func (*ValidationResults) DefaultRole() insolar.DynamicRole {
	return insolar.DynamicRoleVirtualExecutor
}

// DefaultTarget returns of target of this event.
func (vr *ValidationResults) DefaultTarget() *insolar.Reference {
	return &vr.RecordRef
}

func (vr *ValidationResults) Type() insolar.MessageType {
	return insolar.TypeValidationResults
}

// TODO change after changing pulsar
func (vr *ValidationResults) GetCaller() *insolar.Reference {
	return &vr.Caller // TODO actually it's not right. There is no caller.
}

func (vr *ValidationResults) GetReference() insolar.Reference {
	return vr.RecordRef
}

var hasher = platformpolicy.NewPlatformCryptographyScheme().ReferenceHasher() // TODO: create message factory

// GenRequest calculates Reference for request message from pulse number and request's payload.
func genRequest(pn insolar.PulseNumber, payload []byte, domain *insolar.ID) *insolar.Reference {
	ref := insolar.NewReference(
		*domain,
		*insolar.NewID(pn, hasher.Hash(payload)),
	)
	return ref
}

// PendingFinished is sent by the old executor to the current executor
// when pending execution finishes.
type PendingFinished struct {
	Reference insolar.Reference // object pended in executor
}

func (pf *PendingFinished) GetCaller() *insolar.Reference {
	// Contract that initiated this call
	return &pf.Reference
}

func (pf *PendingFinished) AllowedSenderObjectAndRole() (*insolar.Reference, insolar.DynamicRole) {
	// This type of message currently can be send from any node todo: rethink it
	return nil, 0
}

func (pf *PendingFinished) DefaultRole() insolar.DynamicRole {
	return insolar.DynamicRoleVirtualExecutor
}

func (pf *PendingFinished) DefaultTarget() *insolar.Reference {
	return &pf.Reference
}

func (pf *PendingFinished) Type() insolar.MessageType {
	return insolar.TypePendingFinished
}

// StillExecuting
type StillExecuting struct {
	Reference insolar.Reference // object we still executing
}

func (se *StillExecuting) GetCaller() *insolar.Reference {
	return &se.Reference
}

func (se *StillExecuting) AllowedSenderObjectAndRole() (*insolar.Reference, insolar.DynamicRole) {
	return nil, 0
}

func (se *StillExecuting) DefaultRole() insolar.DynamicRole {
	return insolar.DynamicRoleVirtualExecutor
}

func (se *StillExecuting) DefaultTarget() *insolar.Reference {
	return &se.Reference
}

func (se *StillExecuting) Type() insolar.MessageType {
	return insolar.TypeStillExecuting
}
