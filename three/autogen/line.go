package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Line represents a line.
type Line struct{ p *js.Object }

// Line returns a Line object.
func (t *Three) Line() *Line {
	p := t.ctx.Get("Line")
	return &Line{p: p}
}

// New returns a new Line object.
func (t *Line) New(geometry, material, mode float64) *Line {
	p := t.p.New(geometry, material, mode)
	return &Line{p: p}
}

