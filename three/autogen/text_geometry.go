package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// TextGeometry represents a textgeometry.
type TextGeometry struct{ p *js.Object }

// TextGeometry returns a TextGeometry object.
func (t *Three) TextGeometry() *TextGeometry {
	p := t.ctx.Get("TextGeometry")
	return &TextGeometry{p: p}
}

// New returns a new TextGeometry object.
func (t *TextGeometry) New(text, parameters float64) *TextGeometry {
	p := t.p.New(text, parameters)
	return &TextGeometry{p: p}
}

