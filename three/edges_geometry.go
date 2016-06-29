package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// EdgesGeometry represents an edgesgeometry.
type EdgesGeometry struct{ p *js.Object }

// EdgesGeometry returns an EdgesGeometry object.
func (t *Three) EdgesGeometry() *EdgesGeometry {
	p := t.ctx.Get("EdgesGeometry")
	return &EdgesGeometry{p: p}
}

// New returns a new EdgesGeometry object.
func (t *EdgesGeometry) New(geometry, thresholdAngle float64) *EdgesGeometry {
	p := t.p.New(geometry, thresholdAngle)
	return &EdgesGeometry{p: p}
}

