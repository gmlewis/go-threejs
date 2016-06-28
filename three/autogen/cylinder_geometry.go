package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CylinderGeometry represents a cylindergeometry.
type CylinderGeometry struct{ p *js.Object }

// CylinderGeometry returns a CylinderGeometry object.
func (t *Three) CylinderGeometry() *CylinderGeometry {
	p := t.ctx.Get("CylinderGeometry")
	return &CylinderGeometry{p: p}
}

// New returns a new CylinderGeometry object.
func (t *CylinderGeometry) New(radiusTop, radiusBottom, height, radialSegments, heightSegments, openEnded, thetaStart, thetaLength float64) *CylinderGeometry {
	p := t.p.New(radiusTop, radiusBottom, height, radialSegments, heightSegments, openEnded, thetaStart, thetaLength)
	return &CylinderGeometry{p: p}
}

