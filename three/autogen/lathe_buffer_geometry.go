package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// LatheBufferGeometry represents a lathebuffergeometry.
type LatheBufferGeometry struct{ p *js.Object }

// LatheBufferGeometry returns a lathebuffergeometry object.
func (t *Three) LatheBufferGeometry() *LatheBufferGeometry {
	p := t.ctx.Get("LatheBufferGeometry")
	return &LatheBufferGeometry{p: p}
}

// NewLatheBufferGeometry returns a new lathebuffergeometry object.
func (t *LatheBufferGeometry) New(points, segments, phiStart, phiLength float64) *LatheBufferGeometry {
	p := t.p.New(points, segments, phiStart, phiLength)
	return &LatheBufferGeometry{p: p}
}

