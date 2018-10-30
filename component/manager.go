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

package component

import (
	"context"
	"reflect"

	"github.com/insolar/insolar/log"
	"github.com/pkg/errors"
)

// Manager provide methods to manage components lifecycle
type Manager struct {
	components []interface{}
}

// Register components in Manager and inject dependencies
// TODO: write about inject
func (m *Manager) Register(components ...interface{}) {
	m.components = components
	// todo: fix logs
	for _, c := range components {
		componentValue := reflect.ValueOf(c).Elem()
		componentType := componentValue.Type()
		log.Infof("~~~~ NEW ComponentManager: Register component: %s", componentType.String())

		for i := 0; i < componentType.NumField(); i++ {
			f := componentType.Field(i)
			if _, ok := f.Tag.Lookup("inject"); ok {
				log.Debugf("~~~~ NEW ComponentManager: Component %s need inject: ", componentType.String(), f.Name)

				// try to inject
				isInjected := false
				for _, cc := range m.components {
					fieldValue := componentValue.Field(i)
					if reflect.ValueOf(cc).Type().Implements(fieldValue.Type()) {
						fieldValue.Set(reflect.ValueOf(cc))
						log.Infof("~~~~ NEW ComponentManager: Inject interface %s with %s: ", fieldValue.Type().String(), reflect.ValueOf(cc).Type().String())
						isInjected = true
						break
					}
				}

				if !isInjected {
					panic("Not fount component with interface: " + f.Type.String())
				}

			}

		}
	}
}

// Start invokes Start method of all components which implements Starter interface
func (m *Manager) Start(ctx context.Context) error {
	for _, c := range m.components {
		name := reflect.TypeOf(c).Elem().String()
		if s, ok := c.(Starter); ok {
			log.Infoln("~~~~ NEW ComponentManager: Start component: ", name)
			err := s.Start(ctx)
			if err != nil {
				return errors.Wrap(err, "Failed to start components.")
			}
		} else {
			log.Warnf("~~~~ NEW ComponentManager: Component %s has no Stop method", name)
		}
	}
	return nil
}

// Stop invokes Stop method of all components which implements Starter interface
func (m *Manager) Stop(ctx context.Context) error {

	for i := len(m.components) - 1; i >= 0; i-- {
		name := reflect.TypeOf(m.components[i]).Elem().String()
		if s, ok := m.components[i].(Stopper); ok {
			log.Infoln("~~~~ NEW ComponentManager: Stop component: ", name)

			err := s.Stop(ctx)
			if err != nil {
				return errors.Wrap(err, "Failed to stop components.")
			}
		} else {
			log.Warnf("~~~~ NEW ComponentManager: Component %s has no Stop method", name)
		}
	}
	return nil
}
