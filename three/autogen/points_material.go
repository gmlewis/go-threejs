package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// PointsMaterial represents a pointsmaterial.
type PointsMaterial struct{ p *js.Object }

// PointsMaterial returns a pointsmaterial object.
func (t *Three) PointsMaterial() *PointsMaterial {
	p := t.ctx.Get("PointsMaterial")
	return &PointsMaterial{p: p}
}

// NewPointsMaterial returns a new pointsmaterial object.
func (t *PointsMaterial) New(parameters float64) *PointsMaterial {
	p := t.p.New(parameters)
	return &PointsMaterial{p: p}
}

// Copy TODO description.
func (p *PointsMaterial) Copy(source float64) *PointsMaterial {
	p.p.Call("copy", source)
	return p
}

