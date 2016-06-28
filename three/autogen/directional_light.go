package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// DirectionalLight represents a directionallight.
type DirectionalLight struct{ p *js.Object }

// DirectionalLight returns a DirectionalLight object.
func (t *Three) DirectionalLight() *DirectionalLight {
	p := t.ctx.Get("DirectionalLight")
	return &DirectionalLight{p: p}
}

// New returns a new DirectionalLight object.
func (t *DirectionalLight) New(color, intensity float64) *DirectionalLight {
	p := t.p.New(color, intensity)
	return &DirectionalLight{p: p}
}

// Copy TODO description.
func (d *DirectionalLight) Copy(source float64) *DirectionalLight {
	d.p.Call("copy", source)
	return d
}

