package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CylinderBufferGeometry represents a cylinderbuffergeometry.
type CylinderBufferGeometry struct{ p *js.Object }

// CylinderBufferGeometry returns a cylinderbuffergeometry object.
func (t *Three) CylinderBufferGeometry() *CylinderBufferGeometry {
	p := t.ctx.Get("CylinderBufferGeometry")
	return &CylinderBufferGeometry{p: p}
}

// NewCylinderBufferGeometry returns a new cylinderbuffergeometry object.
func (t *CylinderBufferGeometry) New(radiusTop, radiusBottom, height, radialSegments, heightSegments, openEnded, thetaStart, thetaLength float64) *CylinderBufferGeometry {
	p := t.p.New(radiusTop, radiusBottom, height, radialSegments, heightSegments, openEnded, thetaStart, thetaLength)
	return &CylinderBufferGeometry{p: p}
}

