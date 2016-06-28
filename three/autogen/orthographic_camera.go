package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// OrthographicCamera represents an orthographiccamera.
type OrthographicCamera struct{ p *js.Object }

// OrthographicCamera returns an orthographiccamera object.
func (t *Three) OrthographicCamera() *OrthographicCamera {
	p := t.ctx.Get("OrthographicCamera")
	return &OrthographicCamera{p: p}
}

// NewOrthographicCamera returns a new orthographiccamera object.
func (t *OrthographicCamera) New(left, right, top, bottom, near, far float64) *OrthographicCamera {
	p := t.p.New(left, right, top, bottom, near, far)
	return &OrthographicCamera{p: p}
}

// Copy TODO description.
func (o *OrthographicCamera) Copy(source float64) *OrthographicCamera {
	o.p.Call("copy", source)
	return o
}

// ToJSON TODO description.
func (o *OrthographicCamera) ToJSON(meta float64) *OrthographicCamera {
	o.p.Call("toJSON", meta)
	return o
}

