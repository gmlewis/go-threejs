// Copyright 2016 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// EventDispatcher handles JavaScript events for custom objects.
//
// http://threejs.org/docs/index.html#Reference/Core/EventDispatcher
type EventDispatcher struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (e *EventDispatcher) JSObject() *js.Object { return e.p }

// EventDispatcher returns an EventDispatcher JavaScript class.
func (t *Three) EventDispatcher() *EventDispatcher {
	p := t.ctx.Get("EventDispatcher")
	return EventDispatcherFromJSObject(p)
}

// EventDispatcherFromJSObject returns a wrapped EventDispatcher JavaScript class.
func EventDispatcherFromJSObject(p *js.Object) *EventDispatcher {
	return &EventDispatcher{p: p}
}

// NewEventDispatcher returns a new EventDispatcher object.
func (t *Three) NewEventDispatcher() *EventDispatcher {
	p := t.ctx.Get("EventDispatcher").New()
	return EventDispatcherFromJSObject(p)
}

// Apply TODO description.
func (e *EventDispatcher) Apply(object float64) *EventDispatcher {
	e.p.Call("apply", object)
	return e
}

// AddEventListener TODO description.
func (e *EventDispatcher) AddEventListener(typ, listener float64) *EventDispatcher {
	e.p.Call("addEventListener", typ, listener)
	return e
}

// HasEventListener TODO description.
func (e *EventDispatcher) HasEventListener(typ, listener float64) *EventDispatcher {
	e.p.Call("hasEventListener", typ, listener)
	return e
}

// RemoveEventListener TODO description.
func (e *EventDispatcher) RemoveEventListener(typ, listener float64) *EventDispatcher {
	e.p.Call("removeEventListener", typ, listener)
	return e
}

// DispatchEvent TODO description.
func (e *EventDispatcher) DispatchEvent(event float64) *EventDispatcher {
	e.p.Call("dispatchEvent", event)
	return e
}
