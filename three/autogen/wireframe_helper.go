package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WireframeHelper represents a wireframehelper.
type WireframeHelper struct{ p *js.Object }

// WireframeHelper returns a wireframehelper object.
func (t *Three) WireframeHelper() *WireframeHelper {
	p := t.ctx.Get("WireframeHelper")
	return &WireframeHelper{p: p}
}

// NewWireframeHelper returns a new wireframehelper object.
func (t *WireframeHelper) New(object, hex float64) *WireframeHelper {
	p := t.p.New(object, hex)
	return &WireframeHelper{p: p}
}

