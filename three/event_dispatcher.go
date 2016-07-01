// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

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
	return &EventDispatcher{p: p}
}

// NewEventDispatcher returns a new EventDispatcher object.
func (t *Three) NewEventDispatcher() *EventDispatcher {
	p := t.ctx.Get("EventDispatcher").New()
	return &EventDispatcher{p: p}
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
