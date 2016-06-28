package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// RingGeometry represents a ringgeometry.
type RingGeometry struct{ p *js.Object }

// RingGeometry returns a RingGeometry object.
func (t *Three) RingGeometry() *RingGeometry {
	p := t.ctx.Get("RingGeometry")
	return &RingGeometry{p: p}
}

// New returns a new RingGeometry object.
func (t *RingGeometry) New(innerRadius, outerRadius, thetaSegments, phiSegments, thetaStart, thetaLength float64) *RingGeometry {
	p := t.p.New(innerRadius, outerRadius, thetaSegments, phiSegments, thetaStart, thetaLength)
	return &RingGeometry{p: p}
}

