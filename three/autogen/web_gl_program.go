package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLProgram represents a webglprogram.
type WebGLProgram struct{ p *js.Object }

// WebGLProgram returns a webglprogram object.
func (t *Three) WebGLProgram() *WebGLProgram {
	p := t.ctx.Get("WebGLProgram")
	return &WebGLProgram{p: p}
}

// NewWebGLProgram returns a new webglprogram object.
func (t *WebGLProgram) New() *WebGLProgram {
	p := t.p.New()
	return &WebGLProgram{p: p}
}

// Get TODO description.
func (w *WebGLProgram) Get() *WebGLProgram {
	w.p.Call("get")
	return w
}

// Get TODO description.
func (w *WebGLProgram) Get() *WebGLProgram {
	w.p.Call("get")
	return w
}

