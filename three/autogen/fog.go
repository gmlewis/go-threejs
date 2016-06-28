package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Fog represents a fog.
type Fog struct{ p *js.Object }

// Fog returns a Fog object.
func (t *Three) Fog() *Fog {
	p := t.ctx.Get("Fog")
	return &Fog{p: p}
}

// New returns a new Fog object.
func (t *Fog) New(color, near, far float64) *Fog {
	p := t.p.New(color, near, far)
	return &Fog{p: p}
}

