package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLState represents a webglstate.
type WebGLState struct{ p *js.Object }

// WebGLState returns a webglstate object.
func (t *Three) WebGLState() *WebGLState {
	p := t.ctx.Get("WebGLState")
	return &WebGLState{p: p}
}

// NewWebGLState returns a new webglstate object.
func (t *WebGLState) New(gl, extensions, paramThreeToGL float64) *WebGLState {
	p := t.p.New(gl, extensions, paramThreeToGL)
	return &WebGLState{p: p}
}

