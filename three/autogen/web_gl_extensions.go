package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLExtensions represents a webglextensions.
type WebGLExtensions struct{ p *js.Object }

// WebGLExtensions returns a webglextensions object.
func (t *Three) WebGLExtensions() *WebGLExtensions {
	p := t.ctx.Get("WebGLExtensions")
	return &WebGLExtensions{p: p}
}

// NewWebGLExtensions returns a new webglextensions object.
func (t *WebGLExtensions) New(gl float64) *WebGLExtensions {
	p := t.p.New(gl)
	return &WebGLExtensions{p: p}
}

