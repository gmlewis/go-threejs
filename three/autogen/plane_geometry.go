package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// PlaneGeometry represents a planegeometry.
type PlaneGeometry struct{ p *js.Object }

// PlaneGeometry returns a PlaneGeometry object.
func (t *Three) PlaneGeometry() *PlaneGeometry {
	p := t.ctx.Get("PlaneGeometry")
	return &PlaneGeometry{p: p}
}

// New returns a new PlaneGeometry object.
func (t *PlaneGeometry) New(width, height, widthSegments, heightSegments float64) *PlaneGeometry {
	p := t.p.New(width, height, widthSegments, heightSegments)
	return &PlaneGeometry{p: p}
}

