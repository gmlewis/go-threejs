package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// SphereBufferGeometry represents a spherebuffergeometry.
type SphereBufferGeometry struct{ p *js.Object }

// SphereBufferGeometry returns a SphereBufferGeometry object.
func (t *Three) SphereBufferGeometry() *SphereBufferGeometry {
	p := t.ctx.Get("SphereBufferGeometry")
	return &SphereBufferGeometry{p: p}
}

// New returns a new SphereBufferGeometry object.
func (t *SphereBufferGeometry) New(radius, widthSegments, heightSegments, phiStart, phiLength, thetaStart, thetaLength float64) *SphereBufferGeometry {
	p := t.p.New(radius, widthSegments, heightSegments, phiStart, phiLength, thetaStart, thetaLength)
	return &SphereBufferGeometry{p: p}
}

