package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// EdgesGeometry represents an edgesgeometry.
type EdgesGeometry struct{ p *js.Object }

// EdgesGeometry returns an edgesgeometry object.
func (t *Three) EdgesGeometry() *EdgesGeometry {
	p := t.ctx.Get("EdgesGeometry")
	return &EdgesGeometry{p: p}
}

// NewEdgesGeometry returns a new edgesgeometry object.
func (t *EdgesGeometry) New(geometry, thresholdAngle float64) *EdgesGeometry {
	p := t.p.New(geometry, thresholdAngle)
	return &EdgesGeometry{p: p}
}

