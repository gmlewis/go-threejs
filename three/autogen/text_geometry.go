package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// TextGeometry represents a 3D object of text as a single object.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/TextGeometry
type TextGeometry struct{ p *js.Object }

// TextGeometry returns a TextGeometry object.
func (t *Three) TextGeometry() *TextGeometry {
	p := t.ctx.Get("TextGeometry")
	return &TextGeometry{p: p}
}

// New returns a new TextGeometry object.
//
//     text — The text to be shown.
//     parameters — Object that can contains the following parameters.  TODO.
func (t *TextGeometry) New(text string, parameters map[string]interface{}) *TextGeometry {
	p := t.p.New(text, parameters)
	return &TextGeometry{p: p}
}
