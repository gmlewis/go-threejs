package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// SphereGeometry represents a sphere.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/SphereGeometry
type SphereGeometry struct{ p *js.Object }

// SphereGeometry returns a SphereGeometry object.
func (t *Three) SphereGeometry() *SphereGeometry {
	p := t.ctx.Get("SphereGeometry")
	return &SphereGeometry{p: p}
}

// New returns a new SphereGeometry object.
//
//     radius — sphere radius. Default is 50.
//     widthSegments — number of horizontal segments. Minimum value is 3, and the default is 8.
//     heightSegments — number of vertical segments. Minimum value is 2, and the default is 6.
//     phiStart — specify horizontal starting angle. Default is 0.
//     phiLength — specify horizontal sweep angle size. Default is Math.PI * 2.
//     thetaStart — specify vertical starting angle. Default is 0.
//     thetaLength — specify vertical sweep angle size. Default is Math.PI.
func (t *SphereGeometry) New(radius float64, widthSegments, heightSegments int, phiStart, phiLength, thetaStart, thetaLength float64) *SphereGeometry {
	p := t.p.New(radius, widthSegments, heightSegments, phiStart, phiLength, thetaStart, thetaLength)
	return &SphereGeometry{p: p}
}
