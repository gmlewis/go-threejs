package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// LineDashedMaterial represents a linedashedmaterial.
type LineDashedMaterial struct{ p *js.Object }

// LineDashedMaterial returns a LineDashedMaterial object.
func (t *Three) LineDashedMaterial() *LineDashedMaterial {
	p := t.ctx.Get("LineDashedMaterial")
	return &LineDashedMaterial{p: p}
}

// New returns a new LineDashedMaterial object.
func (t *LineDashedMaterial) New(parameters float64) *LineDashedMaterial {
	p := t.p.New(parameters)
	return &LineDashedMaterial{p: p}
}

// Copy TODO description.
func (l *LineDashedMaterial) Copy(source float64) *LineDashedMaterial {
	l.p.Call("copy", source)
	return l
}

