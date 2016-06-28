package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// EventDispatcher represents an eventdispatcher.
type EventDispatcher struct{ p *js.Object }

// EventDispatcher returns an EventDispatcher object.
func (t *Three) EventDispatcher() *EventDispatcher {
	p := t.ctx.Get("EventDispatcher")
	return &EventDispatcher{p: p}
}

// New returns a new EventDispatcher object.
func (t *EventDispatcher) New() *EventDispatcher {
	p := t.p.New()
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

