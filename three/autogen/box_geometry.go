package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// BoxGeometry represents a boxgeometry.
type BoxGeometry struct{ p *js.Object }

// BoxGeometry returns a BoxGeometry object.
func (t *Three) BoxGeometry() *BoxGeometry {
	p := t.ctx.Get("BoxGeometry")
	return &BoxGeometry{p: p}
}

// New returns a new BoxGeometry object.
func (t *BoxGeometry) New(width, height, depth, widthSegments, heightSegments, depthSegments float64) *BoxGeometry {
	p := t.p.New(width, height, depth, widthSegments, heightSegments, depthSegments)
	return &BoxGeometry{p: p}
}

