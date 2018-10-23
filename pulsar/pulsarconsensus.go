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

package pulsar

import (
	"crypto/ecdsa"
	"sync"

	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/log"
	"github.com/pkg/errors"
)

// SetBftGridItem set item of the bftGrid in the thread-safe way
func (currentPulsar *Pulsar) SetBftGridItem(key string, value map[string]*BftCell) {
	currentPulsar.BftGridLock.Lock()
	currentPulsar.bftGrid[key] = value
	defer currentPulsar.BftGridLock.Unlock()
}

// GetBftGridItem returns a grid item i nthe thread-safe way
func (currentPulsar *Pulsar) GetBftGridItem(row string, column string) *BftCell {
	currentPulsar.BftGridLock.RLock()
	defer currentPulsar.BftGridLock.RUnlock()
	return currentPulsar.bftGrid[row][column]
}

// BftCell is a cell in NxN btf-grid
type BftCell struct {
	signLock              sync.RWMutex
	entropyLock           sync.RWMutex
	isEntropyReceivedLock sync.RWMutex

	Sign              []byte
	Entropy           core.Entropy
	IsEntropyReceived bool
}

// SetSign sets Sign in the thread-safe way
func (bftCell *BftCell) SetSign(sign []byte) {
	bftCell.signLock.Lock()
	defer bftCell.signLock.Unlock()
	bftCell.Sign = sign
}

// GetSign gets Sign in the thread-safe way
func (bftCell *BftCell) GetSign() []byte {
	bftCell.signLock.RLock()
	defer bftCell.signLock.RUnlock()
	return bftCell.Sign
}

// SetEntropy sets Entropy in the thread-safe way
func (bftCell *BftCell) SetEntropy(entropy core.Entropy) {
	bftCell.entropyLock.Lock()
	defer bftCell.entropyLock.Unlock()
	bftCell.Entropy = entropy
}

// GetEntropy gets Entropy in the thread-safe way
func (bftCell *BftCell) GetEntropy() core.Entropy {
	bftCell.entropyLock.RLock()
	defer bftCell.entropyLock.RUnlock()
	return bftCell.Entropy
}

// SetIsEntropyReceived sets IsEntropyReceived in the thread-safe way
func (bftCell *BftCell) SetIsEntropyReceived(isEntropyReceived bool) {
	bftCell.isEntropyReceivedLock.Lock()
	defer bftCell.isEntropyReceivedLock.Unlock()
	bftCell.IsEntropyReceived = isEntropyReceived
}

// GetIsEntropyReceived gets IsEntropyReceived in the thread-safe way
func (bftCell *BftCell) GetIsEntropyReceived() bool {
	bftCell.isEntropyReceivedLock.RLock()
	defer bftCell.isEntropyReceivedLock.RUnlock()
	return bftCell.IsEntropyReceived
}

func (currentPulsar *Pulsar) isVerificationNeeded() bool {
	if currentPulsar.IsStateFailed() {
		return false

	}
	if currentPulsar.isStandalone() {
		currentPulsar.CurrentSlotEntropy = currentPulsar.GeneratedEntropy
		currentPulsar.CurrentSlotPulseSender = currentPulsar.PublicKeyRaw
		currentPulsar.StateSwitcher.SwitchToState(SendingPulse, nil)
		return false

	}

	return true
}

func (currentPulsar *Pulsar) verify() {
	log.Debugf("[verify] - %v", currentPulsar.Config.MainListenerAddress)
	if !currentPulsar.isVerificationNeeded() {
		return
	}
	type bftMember struct {
		PubPem string
		PubKey *ecdsa.PublicKey
	}

	var finalEntropySet []core.Entropy

	keys := []string{currentPulsar.PublicKeyRaw}
	activePulsars := []*bftMember{{currentPulsar.PublicKeyRaw, &currentPulsar.PrivateKey.PublicKey}}
	for key, neighbour := range currentPulsar.Neighbours {
		activePulsars = append(activePulsars, &bftMember{key, neighbour.PublicKey})
		keys = append(keys, key)
	}

	// Check NxN consensus-matrix
	wrongVectors := 0
	for _, column := range activePulsars {
		currentColumnStat := map[string]int{}
		for _, row := range activePulsars {
			bftCell := currentPulsar.GetBftGridItem(row.PubPem, column.PubPem)

			if bftCell == nil {
				currentColumnStat["nil"]++
				continue
			}

			ok, err := checkSignature(bftCell.GetEntropy(), column.PubPem, bftCell.GetSign())
			if !ok || err != nil {
				currentColumnStat["nil"]++
				continue
			}

			entropy := bftCell.GetEntropy()
			currentColumnStat[string(entropy[:])]++
		}

		maxConfirmationsForEntropy := int(0)
		var chosenEntropy core.Entropy
		for key, value := range currentColumnStat {
			if value > maxConfirmationsForEntropy && key != "nil" {
				maxConfirmationsForEntropy = value
				copy(chosenEntropy[:], []byte(key)[:core.EntropySize])
			}
		}

		if maxConfirmationsForEntropy >= currentPulsar.getMinimumNonTraitorsCount() {
			finalEntropySet = append(finalEntropySet, chosenEntropy)
		} else {
			wrongVectors++
		}
	}

	if len(finalEntropySet) == 0 || wrongVectors > currentPulsar.getMaxTraitorsCount() {
		currentPulsar.StateSwitcher.SwitchToState(Failed, errors.New("bft is broken"))
		return
	}

	var finalEntropy core.Entropy

	for _, tempEntropy := range finalEntropySet {
		for byteIndex := 0; byteIndex < core.EntropySize; byteIndex++ {
			finalEntropy[byteIndex] ^= tempEntropy[byteIndex]
		}
	}
	currentPulsar.finalizeBft(finalEntropy, keys)
}

func (currentPulsar *Pulsar) finalizeBft(finalEntropy core.Entropy, activePulsars []string) {
	currentPulsar.CurrentSlotEntropy = finalEntropy
	chosenPulsar, err := selectByEntropy(finalEntropy, activePulsars, len(activePulsars))
	if err != nil {
		currentPulsar.StateSwitcher.SwitchToState(Failed, err)
	}
	currentPulsar.CurrentSlotPulseSender = chosenPulsar[0]
	if currentPulsar.CurrentSlotPulseSender == currentPulsar.PublicKeyRaw {
		//here confirmation myself
		signature, err := signData(currentPulsar.PrivateKey, core.PulseSenderConfirmation{
			ChosenPublicKey: currentPulsar.CurrentSlotPulseSender,
			Entropy:         currentPulsar.CurrentSlotEntropy,
			PulseNumber:     currentPulsar.ProcessingPulseNumber,
		})
		if err != nil {
			currentPulsar.StateSwitcher.SwitchToState(Failed, err)
			return
		}
		currentPulsar.currentSlotSenderConfirmationsLock.Lock()
		currentPulsar.CurrentSlotSenderConfirmations[currentPulsar.PublicKeyRaw] = core.PulseSenderConfirmation{
			ChosenPublicKey: currentPulsar.CurrentSlotPulseSender,
			Signature:       signature,
			Entropy:         currentPulsar.CurrentSlotEntropy,
			PulseNumber:     currentPulsar.ProcessingPulseNumber,
		}
		currentPulsar.currentSlotSenderConfirmationsLock.Unlock()

		currentPulsar.StateSwitcher.SwitchToState(WaitingForPulseSigns, nil)
	} else {
		currentPulsar.StateSwitcher.SwitchToState(SendingPulseSign, nil)
	}
}
