package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CircleGeometry represents a circlegeometry.
type CircleGeometry struct{ p *js.Object }

// CircleGeometry returns a circlegeometry object.
func (t *Three) CircleGeometry() *CircleGeometry {
	p := t.ctx.Get("CircleGeometry")
	return &CircleGeometry{p: p}
}

// NewCircleGeometry returns a new circlegeometry object.
func (t *CircleGeometry) New(radius, segments, thetaStart, thetaLength float64) *CircleGeometry {
	p := t.p.New(radius, segments, thetaStart, thetaLength)
	return &CircleGeometry{p: p}
}

