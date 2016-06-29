package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// SphereGeometry represents a spheregeometry.
type SphereGeometry struct{ p *js.Object }

// SphereGeometry returns a SphereGeometry object.
func (t *Three) SphereGeometry() *SphereGeometry {
	p := t.ctx.Get("SphereGeometry")
	return &SphereGeometry{p: p}
}

// New returns a new SphereGeometry object.
func (t *SphereGeometry) New(radius, widthSegments, heightSegments, phiStart, phiLength, thetaStart, thetaLength float64) *SphereGeometry {
	p := t.p.New(radius, widthSegments, heightSegments, phiStart, phiLength, thetaStart, thetaLength)
	return &SphereGeometry{p: p}
}
