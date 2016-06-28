package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// RingBufferGeometry represents a ringbuffergeometry.
type RingBufferGeometry struct{ p *js.Object }

// RingBufferGeometry returns a RingBufferGeometry object.
func (t *Three) RingBufferGeometry() *RingBufferGeometry {
	p := t.ctx.Get("RingBufferGeometry")
	return &RingBufferGeometry{p: p}
}

// New returns a new RingBufferGeometry object.
func (t *RingBufferGeometry) New(innerRadius, outerRadius, thetaSegments, phiSegments, thetaStart, thetaLength float64) *RingBufferGeometry {
	p := t.p.New(innerRadius, outerRadius, thetaSegments, phiSegments, thetaStart, thetaLength)
	return &RingBufferGeometry{p: p}
}

