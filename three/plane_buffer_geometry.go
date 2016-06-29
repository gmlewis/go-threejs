package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// PlaneBufferGeometry represents a planebuffergeometry.
type PlaneBufferGeometry struct{ p *js.Object }

// PlaneBufferGeometry returns a PlaneBufferGeometry object.
func (t *Three) PlaneBufferGeometry() *PlaneBufferGeometry {
	p := t.ctx.Get("PlaneBufferGeometry")
	return &PlaneBufferGeometry{p: p}
}

// New returns a new PlaneBufferGeometry object.
func (t *PlaneBufferGeometry) New(width, height, widthSegments, heightSegments float64) *PlaneBufferGeometry {
	p := t.p.New(width, height, widthSegments, heightSegments)
	return &PlaneBufferGeometry{p: p}
}

