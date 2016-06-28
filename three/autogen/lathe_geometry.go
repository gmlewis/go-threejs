package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// LatheGeometry represents a lathegeometry.
type LatheGeometry struct{ p *js.Object }

// LatheGeometry returns a LatheGeometry object.
func (t *Three) LatheGeometry() *LatheGeometry {
	p := t.ctx.Get("LatheGeometry")
	return &LatheGeometry{p: p}
}

// New returns a new LatheGeometry object.
func (t *LatheGeometry) New(points, segments, phiStart, phiLength float64) *LatheGeometry {
	p := t.p.New(points, segments, phiStart, phiLength)
	return &LatheGeometry{p: p}
}

