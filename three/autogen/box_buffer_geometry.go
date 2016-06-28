package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// BoxBufferGeometry represents a boxbuffergeometry.
type BoxBufferGeometry struct{ p *js.Object }

// BoxBufferGeometry returns a boxbuffergeometry object.
func (t *Three) BoxBufferGeometry() *BoxBufferGeometry {
	p := t.ctx.Get("BoxBufferGeometry")
	return &BoxBufferGeometry{p: p}
}

// NewBoxBufferGeometry returns a new boxbuffergeometry object.
func (t *BoxBufferGeometry) New(width, height, depth, widthSegments, heightSegments, depthSegments float64) *BoxBufferGeometry {
	p := t.p.New(width, height, depth, widthSegments, heightSegments, depthSegments)
	return &BoxBufferGeometry{p: p}
}

