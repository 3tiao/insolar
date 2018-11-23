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

package packets

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"
	"math/bits"

	"github.com/pkg/errors"
)

const lastBitMask = 0x01

// TriStateBitSet bitset implementation.
type TriStateBitSet struct {
	CompressedSet bool

	cells  []BitSetCell
	mapper BitSetMapper
}

// NewTriStateBitSet creates and returns a tristatebitset.
func NewTriStateBitSet(cells []BitSetCell, mapper BitSetMapper) (*TriStateBitSet, error) {
	if (mapper == nil) || (cells == nil) {
		return nil, errors.New("[ NewTriStateBitSet ] failed to create tristatebitset")
	}
	bitset := &TriStateBitSet{
		cells:  make([]BitSetCell, mapper.Length()),
		mapper: mapper,
	}
	bitset.ApplyChanges(cells)
	return bitset, nil
}

func (dbs *TriStateBitSet) GetCells() []BitSetCell {
	return dbs.cells
}

func (dbs *TriStateBitSet) ApplyChanges(changes []BitSetCell) {
	for _, cell := range changes {
		err := dbs.changeBucketState(&cell)
		if err != nil {
			panic(err)
		}
	}
}

func (dbs *TriStateBitSet) Serialize() ([]byte, error) {
	firstByte := uint8(0)
	secondByte := uint8(0)
	if dbs.CompressedSet {
		firstByte = 0x01
	} else {
		firstByte = 0x00
	}

	array, err := dbs.cellsToBitArray()
	if err != nil {
		return nil, errors.Wrap(err, "[ Serialize ] failed to get bitarray from cells")
	}

	tmpLen := array.Len()
	totalSize := int(math.Round(float64((tmpLen*2)/8))+0.5) + 1 // first byte
	var result *bytes.Buffer
	firstByte = firstByte << 1
	if bits.Len(tmpLen) > 6 {
		firstByte++
		firstByte = firstByte << 6
		secondByte = uint8(tmpLen)
		totalSize += 1 // secondbyte is optional
		result = allocateBuffer(totalSize)
		err = binary.Write(result, defaultByteOrder, firstByte)
		if err != nil {
			return nil, errors.Wrap(err, "[ Serialize ] failed to write binary")
		}
		err = binary.Write(result, defaultByteOrder, secondByte)
		if err != nil {
			return nil, errors.Wrap(err, "[ Serialize ] failed to write binary")
		}
	} else {
		result = allocateBuffer(totalSize)
		firstByte = firstByte << 1
		firstByte += uint8(tmpLen)
		err = binary.Write(result, defaultByteOrder, firstByte)
		if err != nil {
			return nil, errors.Wrap(err, "[ Serialize ] failed to write binary")
		}
	}

	data, err := array.serialize()
	if err != nil {
		return nil, errors.Wrap(err, "[ Serialize ] failed to serialize a bitarray")
	}
	err = binary.Write(result, defaultByteOrder, data)
	if err != nil {
		return nil, errors.Wrap(err, "[ Serialize ] failed to write binary")
	}

	return result.Bytes(), nil
}

func (dbs *TriStateBitSet) Deserialize(data io.Reader) error {
	return nil
}

func (dbs *TriStateBitSet) changeBucketState(cell *BitSetCell) error {
	n, err := dbs.mapper.RefToIndex(cell.NodeID)
	if err != nil {
		return errors.Wrap(err, "[ changeBucketState ] failed to get index from ref")
	}
	dbs.cells[n] = *cell
	return nil
}

func putLastBit(array *bitArray, state TriState, index int) error {
	bit := int(state & lastBitMask)
	err := array.put(bit, index)
	return err
	// return errors.Wrap(err, "[ putLastBit ] failed to put a bit ti bitset")
}

func changeBitState(array *bitArray, i int, state TriState) error {
	err := putLastBit(array, state>>1, 2*i)
	if err != nil {
		return errors.Wrap(err, "[ changeBitState ] failed to put last bit")
	}
	err = putLastBit(array, state, 2*i+1)
	if err != nil {
		return errors.Wrap(err, "[ changeBitState ] failed to put last bit")
	}
	return nil
}

func (dbs *TriStateBitSet) cellsToBitArray() (*bitArray, error) {
	array := newBitArray(uint(dbs.mapper.Length() * 2))
	for i := 0; i < len(dbs.cells); i++ {
		cell := dbs.cells[i]
		index, err := dbs.mapper.RefToIndex(cell.NodeID)
		if err != nil {
			return nil, errors.Wrap(err, "[ cellsToBitArray ] failed to get index from ref")
		}
		err = changeBitState(array, index, cell.State)
		if err != nil {
			return nil, errors.Wrap(err, "[ cellsToBitArray ] failed to change bit state")
		}
	}
	return array, nil
}
