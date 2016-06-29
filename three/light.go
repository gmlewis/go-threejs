package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Light represents a light.
type Light struct{ p *js.Object }

// Light returns a Light object.
func (t *Three) Light() *Light {
	p := t.ctx.Get("Light")
	return &Light{p: p}
}

// New returns a new Light object.
func (t *Light) New(color, intensity float64) *Light {
	p := t.p.New(color, intensity)
	return &Light{p: p}
}

// Copy TODO description.
func (l *Light) Copy(source float64) *Light {
	l.p.Call("copy", source)
	return l
}

// ToJSON TODO description.
func (l *Light) ToJSON(meta float64) *Light {
	l.p.Call("toJSON", meta)
	return l
}

