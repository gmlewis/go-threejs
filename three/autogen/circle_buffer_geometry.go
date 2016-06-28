package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CircleBufferGeometry represents a circlebuffergeometry.
type CircleBufferGeometry struct{ p *js.Object }

// CircleBufferGeometry returns a CircleBufferGeometry object.
func (t *Three) CircleBufferGeometry() *CircleBufferGeometry {
	p := t.ctx.Get("CircleBufferGeometry")
	return &CircleBufferGeometry{p: p}
}

// New returns a new CircleBufferGeometry object.
func (t *CircleBufferGeometry) New(radius, segments, thetaStart, thetaLength float64) *CircleBufferGeometry {
	p := t.p.New(radius, segments, thetaStart, thetaLength)
	return &CircleBufferGeometry{p: p}
}

