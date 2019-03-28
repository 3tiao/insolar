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

package generator

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"text/template"
)

const (
	insolarRep           = "github.com/insolar/insolar"
	stateMachineTemplate = "conveyor/generator/generator/templates/state_machine.go.tpl"
	matrixTemplate       = "conveyor/generator/generator/templates/matrix.go.tpl"
	generatedMatrix      = "conveyor/generator/matrix/matrix.go"
)

// todo remove this type
type TemporaryCustomAdapterHelper struct{}

type Generator struct {
	stateMachines     []*StateMachine
	fullPathToInsolar string
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func exitWithError(errMsg string, a ...interface{}) {
	log.Fatal(fmt.Sprintf(errMsg, a...))
}

func NewGenerator() *Generator {
	_, me, _, ok := runtime.Caller(0)
	if !ok {
		exitWithError("couldn't get self full path")
	}
	idx := strings.LastIndex(string(me), insolarRep)
	return &Generator{
		fullPathToInsolar: string(me)[0 : idx+len(insolarRep)],
	}
}

func (g *Generator) CheckAllMachines() {
	for _, machine := range g.stateMachines {
		if machine.States[0].handlers[Present][Transition] == nil {
			log.Fatal("Present Init handler should be defined")
		}
		if machine.States[0].handlers[Future][Transition] == nil {
			log.Fatal("Future Init handler should be defined")
		}

		for i, s := range machine.States {
			for _, ps := range []PulseState{Future, Present, Past} {
				for _, ht := range []handlerType{Transition, Migration, AdapterResponse} {
					if s.handlers[ps][ht] == nil {
						continue
					}

					if len(s.handlers[ps][ht].params) < 3 || s.handlers[ps][ht].params[2] != *machine.InputEventType {
						log.Fatalf("[%s] Forth parameter should be %s\n", s.handlers[ps][ht].Name, *machine.InputEventType)
					}

					if i == 0 && ht == Transition {
						if s.handlers[ps][ht].params[3] != "interface {}" {
							log.Fatalf("[%s] Init handlers should have interface{} as payload parameter\n", s.handlers[ps][ht].Name)
						}
						if s.handlers[ps][ht].results[1] != *machine.PayloadType {
							log.Fatalf("[%s] Init handlers should return payload as %s\n", s.handlers[ps][ht].Name, *machine.PayloadType)
						}
					} else {
						if s.handlers[ps][ht].params[3] != *machine.PayloadType {
							log.Fatalf("[%s] Handlers payload should be %s not %s\n", s.handlers[ps][ht].Name, *machine.PayloadType, s.handlers[ps][ht].params[3])
						}
						if len(s.handlers[ps][ht].results) != 1 {
							log.Fatalf("[%s] Handlers should return only fsm.ElementState\n", s.handlers[ps][ht].Name)
						}
					}

					if i != 0 && ht == Transition {
						if len(s.handlers[ps][ht].params) != 4 && len(s.handlers[ps][ht].params) != 5 {
							log.Fatalf("[%s] Transition handlers should have 4 or 5 (with adapher helper) parameters\n", s.handlers[ps][ht].Name)
						}
					}

					if ht == AdapterResponse {
						if len(s.handlers[ps][ht].params) != 5 {
							log.Fatalf("[%s] AdapterResponse handlers should have 5 parameters\n", s.handlers[ps][ht].Name)
						}
					}
				}
			}
		}
	}
}

type stateMachineWithID struct {
	StateMachine
	ID int
}

func (g *Generator) GenerateStateMachines() {
	for i, machine := range g.stateMachines {
		tplBody, err := ioutil.ReadFile(path.Join(g.fullPathToInsolar, stateMachineTemplate))
		checkErr(err)

		file, err := os.Create(machine.File[:len(machine.File)-3] + "_generated.go")
		checkErr(err)

		defer file.Close()
		out := bufio.NewWriter(file)

		err = template.Must(template.New("smTmpl").Funcs(templateFuncs).
			Parse(string(tplBody))).
			Execute(out, stateMachineWithID{StateMachine: *machine, ID: i + 1})
		checkErr(err)

		err = out.Flush()
		checkErr(err)
	}

}
func (g *Generator) GenerateMatrix() {
	tplBody, err := ioutil.ReadFile(path.Join(g.fullPathToInsolar, matrixTemplate))
	checkErr(err)

	file, err := os.Create(path.Join(g.fullPathToInsolar, generatedMatrix))
	checkErr(err)

	defer file.Close()
	out := bufio.NewWriter(file)

	err = template.Must(template.New("MtTmpl").Funcs(templateFuncs).
		Parse(string(tplBody))).
		Execute(out, g.stateMachines)
	checkErr(err)

	err = out.Flush()
	checkErr(err)
}
