package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// FogExp2 represents a fogexp2.
type FogExp2 struct{ p *js.Object }

// FogExp2 returns a fogexp2 object.
func (t *Three) FogExp2() *FogExp2 {
	p := t.ctx.Get("FogExp2")
	return &FogExp2{p: p}
}

// NewFogExp2 returns a new fogexp2 object.
func (t *FogExp2) New(color, density float64) *FogExp2 {
	p := t.p.New(color, density)
	return &FogExp2{p: p}
}

