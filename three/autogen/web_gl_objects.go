package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLObjects represents a webglobjects.
type WebGLObjects struct{ p *js.Object }

// WebGLObjects returns a WebGLObjects object.
func (t *Three) WebGLObjects() *WebGLObjects {
	p := t.ctx.Get("WebGLObjects")
	return &WebGLObjects{p: p}
}

// New returns a new WebGLObjects object.
func (t *WebGLObjects) New(gl, properties, info float64) *WebGLObjects {
	p := t.p.New(gl, properties, info)
	return &WebGLObjects{p: p}
}

