package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WireframeGeometry represents a wireframegeometry.
type WireframeGeometry struct{ p *js.Object }

// WireframeGeometry returns a wireframegeometry object.
func (t *Three) WireframeGeometry() *WireframeGeometry {
	p := t.ctx.Get("WireframeGeometry")
	return &WireframeGeometry{p: p}
}

// NewWireframeGeometry returns a new wireframegeometry object.
func (t *WireframeGeometry) New(geometry float64) *WireframeGeometry {
	p := t.p.New(geometry)
	return &WireframeGeometry{p: p}
}

