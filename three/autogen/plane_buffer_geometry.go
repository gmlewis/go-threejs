package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// PlaneBufferGeometry represents a planebuffergeometry.
type PlaneBufferGeometry struct{ p *js.Object }

// PlaneBufferGeometry returns a planebuffergeometry object.
func (t *Three) PlaneBufferGeometry() *PlaneBufferGeometry {
	p := t.ctx.Get("PlaneBufferGeometry")
	return &PlaneBufferGeometry{p: p}
}

// NewPlaneBufferGeometry returns a new planebuffergeometry object.
func (t *PlaneBufferGeometry) New(width, height, widthSegments, heightSegments float64) *PlaneBufferGeometry {
	p := t.p.New(width, height, widthSegments, heightSegments)
	return &PlaneBufferGeometry{p: p}
}

