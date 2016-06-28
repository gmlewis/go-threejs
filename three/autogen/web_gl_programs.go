package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLPrograms represents a webglprograms.
type WebGLPrograms struct{ p *js.Object }

// WebGLPrograms returns a WebGLPrograms object.
func (t *Three) WebGLPrograms() *WebGLPrograms {
	p := t.ctx.Get("WebGLPrograms")
	return &WebGLPrograms{p: p}
}

// New returns a new WebGLPrograms object.
func (t *WebGLPrograms) New(renderer, capabilities float64) *WebGLPrograms {
	p := t.p.New(renderer, capabilities)
	return &WebGLPrograms{p: p}
}

