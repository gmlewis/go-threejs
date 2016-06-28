package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Uniform represents an uniform.
type Uniform struct{ p *js.Object }

// Uniform returns an uniform object.
func (t *Three) Uniform() *Uniform {
	p := t.ctx.Get("Uniform")
	return &Uniform{p: p}
}

// NewUniform returns a new uniform object.
func (t *Uniform) New(typ, value float64) *Uniform {
	p := t.p.New(typ, value)
	return &Uniform{p: p}
}

// OnUpdate TODO description.
func (u *Uniform) OnUpdate(callback float64) *Uniform {
	u.p.Call("onUpdate", callback)
	return u
}
