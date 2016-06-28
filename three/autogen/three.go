package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// MOUSE represents a mouse.
type MOUSE struct{ p *js.Object }

// MOUSE returns a mouse object.
func (t *Three) MOUSE() *MOUSE {
	p := t.ctx.Get("MOUSE")
	return &MOUSE{p: p}
}

// NewMOUSE returns a new mouse object.
func (t *MOUSE) New() *MOUSE {
	p := t.p.New()
	return &MOUSE{p: p}
}

// Get TODO description.
func (m *MOUSE) Get() *MOUSE {
	m.p.Call("get")
	return m
}

// Value TODO description.
func (m *MOUSE) Value(target float64) *MOUSE {
	m.p.Call("value", target)
	return m
}

