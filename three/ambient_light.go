package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// AmbientLight represents an ambientlight.
type AmbientLight struct{ p *js.Object }

// AmbientLight returns an AmbientLight object.
func (t *Three) AmbientLight() *AmbientLight {
	p := t.ctx.Get("AmbientLight")
	return &AmbientLight{p: p}
}

// New returns a new AmbientLight object.
func (t *AmbientLight) New(color, intensity float64) *AmbientLight {
	p := t.p.New(color, intensity)
	return &AmbientLight{p: p}
}

