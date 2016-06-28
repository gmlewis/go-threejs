package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// IcosahedronGeometry represents an icosahedrongeometry.
type IcosahedronGeometry struct{ p *js.Object }

// IcosahedronGeometry returns an icosahedrongeometry object.
func (t *Three) IcosahedronGeometry() *IcosahedronGeometry {
	p := t.ctx.Get("IcosahedronGeometry")
	return &IcosahedronGeometry{p: p}
}

// NewIcosahedronGeometry returns a new icosahedrongeometry object.
func (t *IcosahedronGeometry) New(radius, detail float64) *IcosahedronGeometry {
	p := t.p.New(radius, detail)
	return &IcosahedronGeometry{p: p}
}

