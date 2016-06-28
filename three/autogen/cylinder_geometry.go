package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CylinderGeometry represents a cylindergeometry.
type CylinderGeometry struct{ p *js.Object }

// CylinderGeometry returns a cylindergeometry object.
func (t *Three) CylinderGeometry() *CylinderGeometry {
	p := t.ctx.Get("CylinderGeometry")
	return &CylinderGeometry{p: p}
}

// NewCylinderGeometry returns a new cylindergeometry object.
func (t *CylinderGeometry) New(radiusTop, radiusBottom, height, radialSegments, heightSegments, openEnded, thetaStart, thetaLength float64) *CylinderGeometry {
	p := t.p.New(radiusTop, radiusBottom, height, radialSegments, heightSegments, openEnded, thetaStart, thetaLength)
	return &CylinderGeometry{p: p}
}

