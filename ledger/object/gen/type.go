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

// +build ignore

// This program generates type_gen.go. It can be invoked by running
// go generate

package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"sort"

	"github.com/insolar/insolar/ledger/object"
)

type typedRecord struct {
	TypeID   object.TypeID
	TypeName string
	HasInit  bool
}

var (
	stdout      = flag.Bool("stdout", false, "print result to stdout")
	verboseFlag = flag.Bool("verbose", false, "debug output to stderr")
)

var verbose = struct {
	Printfln func(format string, a ...interface{})
	Println  func(a ...interface{})
}{
	Printfln: func(format string, a ...interface{}) {},
	Println:  func(a ...interface{}) {},
}

func main() {
	flag.Parse()

	if *verboseFlag {
		verbose.Println = verbosePrintln
		verbose.Printfln = verbosePrintfln
	}

	out := os.Stdout
	finalize := func() {}
	if !*stdout {
		f, err := ioutil.TempFile(os.TempDir(), "")
		fatal(err)
		finalize = func() {
			f.Close()
			err = os.Rename(f.Name(), "type_gen.go")
			fatal(err)
		}
		out = f
	}

	reg := object.Registered()
	records := make([]typedRecord, 0, len(reg))
	for id, elem := range reg {
		elemType := reflect.TypeOf(elem)
		elemName := elemType.Elem().Name()
		record := typedRecord{
			TypeID:   id,
			TypeName: elemName,
		}
		verbose.Printfln("process elem name=%v, type=%v", elemName, elemType)

		m, ok := reflect.TypeOf(elem).MethodByName("Init")
		if ok {
			numIn, numOut := m.Type.NumIn(), m.Type.NumOut()
			if numIn == 1 && numOut == 1 {
				inType, outType := m.Type.In(0), m.Type.Out(0)
				if inType.Name() == elemType.Name() && outType.Name() == elemType.Name() {
					record.HasInit = true
					verbose.Println("        ", elemType, "has special Init() method")
				}
			}
		}
		records = append(records, record)
	}
	sort.Slice(records, func(i, j int) bool {
		return records[i].TypeID < records[j].TypeID
	})

	err := typeTemplate.Execute(out, struct {
		Records []typedRecord
	}{
		Records: records,
	})
	fatal(err)

	finalize()
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var typeTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
// This file was generated by: go run gen/type.go

package object

import (
	"fmt"

	record "github.com/insolar/insolar/insolar/record"
)

func TypeFromRecord(generic record.VirtualRecord) TypeID {
	switch generic.(type) {
	{{- range .Records }}
	case *{{ .TypeName }}:
		return {{ printf "%d" .TypeID }}
	{{- end }}
	default:
		panic(fmt.Sprintf("%T record is not registered", generic))
	}
}

func RecordFromType(i TypeID) record.VirtualRecord {
	switch i {
	{{- range .Records }}
	case {{ printf "%d" .TypeID }}:
		return new({{ .TypeName }}){{ if .HasInit }}.Init(){{ end }}
	{{- end }}
	default:
		panic(fmt.Sprintf("identificator %d is not registered", i))
	}
}

func (i TypeID) String() string {
	switch i {
	{{- range .Records }}
	case {{ printf "%d" .TypeID }}:
		return "{{ .TypeName }}"
	{{- end }}
	default:
		panic(fmt.Sprintf("identificator %d is not registered", i))
	}
}
`))

func verbosePrintfln(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format+"\n", a...)
}

func verbosePrintln(a ...interface{}) {
	_, _ = fmt.Fprintln(os.Stderr, a...)
}