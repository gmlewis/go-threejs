package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Points represents a points.
type Points struct{ p *js.Object }

// Points returns a Points object.
func (t *Three) Points() *Points {
	p := t.ctx.Get("Points")
	return &Points{p: p}
}

// New returns a new Points object.
func (t *Points) New(geometry, material float64) *Points {
	p := t.p.New(geometry, material)
	return &Points{p: p}
}

