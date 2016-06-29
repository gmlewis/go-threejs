package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WireframeHelper represents a wireframehelper.
type WireframeHelper struct{ p *js.Object }

// WireframeHelper returns a WireframeHelper object.
func (t *Three) WireframeHelper() *WireframeHelper {
	p := t.ctx.Get("WireframeHelper")
	return &WireframeHelper{p: p}
}

// New returns a new WireframeHelper object.
func (t *WireframeHelper) New(object, hex float64) *WireframeHelper {
	p := t.p.New(object, hex)
	return &WireframeHelper{p: p}
}

