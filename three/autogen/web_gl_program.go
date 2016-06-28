package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLProgram represents a webglprogram.
type WebGLProgram struct{ p *js.Object }

// WebGLProgram returns a WebGLProgram object.
func (t *Three) WebGLProgram() *WebGLProgram {
	p := t.ctx.Get("WebGLProgram")
	return &WebGLProgram{p: p}
}

// New returns a new WebGLProgram object.
func (t *WebGLProgram) New() *WebGLProgram {
	p := t.p.New()
	return &WebGLProgram{p: p}
}

// Get TODO description.
func (w *WebGLProgram) Get() *WebGLProgram {
	w.p.Call("get")
	return w
}
