package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// LightShadow represents a lightshadow.
type LightShadow struct{ p *js.Object }

// LightShadow returns a LightShadow object.
func (t *Three) LightShadow() *LightShadow {
	p := t.ctx.Get("LightShadow")
	return &LightShadow{p: p}
}

// New returns a new LightShadow object.
func (t *LightShadow) New(camera float64) *LightShadow {
	p := t.p.New(camera)
	return &LightShadow{p: p}
}

// Copy TODO description.
func (l *LightShadow) Copy(source float64) *LightShadow {
	l.p.Call("copy", source)
	return l
}

// Clone TODO description.
func (l *LightShadow) Clone() *LightShadow {
	l.p.Call("clone")
	return l
}

