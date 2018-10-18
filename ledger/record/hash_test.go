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

package record

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type recordgen func() Record

var emptyRecordsGens = []recordgen{
	// request records
	func() Record { return &CallRequest{} },
	// result records
	func() Record { return &ClassActivateRecord{} },
	func() Record { return &ObjectActivateRecord{} },
	func() Record { return &CodeRecord{} },
	func() Record { return &ClassAmendRecord{} },
	func() Record { return &DeactivationRecord{} },
	func() Record { return &ObjectAmendRecord{} },
	func() Record { return &TypeRecord{} },
	func() Record { return &ChildRecord{} },
	func() Record { return &GenesisRecord{} },
}

func getRecordHash(rec Record) []byte {
	buff := bytes.NewBuffer(nil)
	rec.WriteHashData(buff)
	return buff.Bytes()
}

func Test_HashesNotTheSameOnDifferentTypes(t *testing.T) {
	found := make(map[string]string)
	for _, recFn := range emptyRecordsGens {
		rec := recFn()
		recType := fmt.Sprintf("%T", rec)
		hashHex := fmt.Sprintf("%x", getRecordHash(rec))
		// fmt.Println(recType, "=>", hashHex)
		typename, ok := found[hashHex]
		if !ok {
			found[hashHex] = recType
			continue
		}
		t.Errorf("same hashes for %s and %s types, empty struct with different types should not be the same", recType, typename)
	}
}

func Test_HashesTheSame(t *testing.T) {
	hashes := make([]string, len(emptyRecordsGens))
	for i, recFn := range emptyRecordsGens {
		rec := recFn()
		hashHex := fmt.Sprintf("%x", getRecordHash(rec))
		hashes[i] = hashHex
	}

	// same struct with different should produce the same hashes
	for i, recFn := range emptyRecordsGens {
		rec := recFn()
		hashHex := fmt.Sprintf("%x", getRecordHash(rec))
		assert.Equal(t, hashes[i], hashHex)
	}
}

var hashtestsRecordsMutate = []struct {
	typ     string
	records []Record
}{
	{
		"CodeRecord",
		[]Record{
			&CodeRecord{},
			&CodeRecord{Code: []byte{1, 2, 3}},
			&CodeRecord{
				Code: []byte{1, 2, 3},
				ResultRecord: ResultRecord{
					Domain: Reference{
						Record: str2ID("0A"),
					},
				},
			},
		},
	},
}

func Test_CBORhashesMutation(t *testing.T) {
	for _, tt := range hashtestsRecordsMutate {
		found := make(map[string]string)
		for _, rec := range tt.records {
			h := getRecordHash(rec)
			hHex := fmt.Sprintf("%x", h)

			typ, ok := found[hHex]
			if !ok {
				found[hHex] = tt.typ
				continue
			}
			t.Errorf("%s failed: found %s hash for \"%s\" test, should not repeats in sha3hash224tests", tt.typ, hHex, typ)
		}
	}
}
