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

// /go:generate protoc -I./vendor -I./ --gogoslick_out=./ insolar/record/record.proto
package record

import (
	"io"

	"github.com/insolar/insolar/insolar"
)

type Record interface{}

type VirtualRecord interface {
	// WriteHashData writes record data to provided writer. This data is used to calculate record's hash.
	WriteHashData(w io.Writer) (int, error)
}

type MaterialRecord struct {
	Record VirtualRecord

	JetID insolar.JetID
}
