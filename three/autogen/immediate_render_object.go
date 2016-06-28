package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// ImmediateRenderObject represents an immediaterenderobject.
type ImmediateRenderObject struct{ p *js.Object }

// ImmediateRenderObject returns an ImmediateRenderObject object.
func (t *Three) ImmediateRenderObject() *ImmediateRenderObject {
	p := t.ctx.Get("ImmediateRenderObject")
	return &ImmediateRenderObject{p: p}
}

// New returns a new ImmediateRenderObject object.
func (t *ImmediateRenderObject) New(material float64) *ImmediateRenderObject {
	p := t.p.New(material)
	return &ImmediateRenderObject{p: p}
}

