package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLExtensions represents a webglextensions.
type WebGLExtensions struct{ p *js.Object }

// WebGLExtensions returns a WebGLExtensions object.
func (t *Three) WebGLExtensions() *WebGLExtensions {
	p := t.ctx.Get("WebGLExtensions")
	return &WebGLExtensions{p: p}
}

// New returns a new WebGLExtensions object.
func (t *WebGLExtensions) New(gl float64) *WebGLExtensions {
	p := t.p.New(gl)
	return &WebGLExtensions{p: p}
}

