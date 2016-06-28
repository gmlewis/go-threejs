package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// LineSegments represents a linesegments.
type LineSegments struct{ p *js.Object }

// LineSegments returns a LineSegments object.
func (t *Three) LineSegments() *LineSegments {
	p := t.ctx.Get("LineSegments")
	return &LineSegments{p: p}
}

// New returns a new LineSegments object.
func (t *LineSegments) New(geometry, material float64) *LineSegments {
	p := t.p.New(geometry, material)
	return &LineSegments{p: p}
}

