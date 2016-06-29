package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// FogExp2 represents exponential fog.
//
// http://threejs.org/docs/index.html#Reference/Scenes/FogExp2
type FogExp2 struct{ p *js.Object }

// FogExp2 returns a FogExp2 object.
func (t *Three) FogExp2() *FogExp2 {
	p := t.ctx.Get("FogExp2")
	return &FogExp2{p: p}
}

// New returns a new FogExp2 object.
func (t *FogExp2) New(color, density float64) *FogExp2 {
	p := t.p.New(color, density)
	return &FogExp2{p: p}
}
