package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLCapabilities represents a webglcapabilities.
type WebGLCapabilities struct{ p *js.Object }

// WebGLCapabilities returns a WebGLCapabilities object.
func (t *Three) WebGLCapabilities() *WebGLCapabilities {
	p := t.ctx.Get("WebGLCapabilities")
	return &WebGLCapabilities{p: p}
}

// New returns a new WebGLCapabilities object.
func (t *WebGLCapabilities) New(gl, extensions, parameters float64) *WebGLCapabilities {
	p := t.p.New(gl, extensions, parameters)
	return &WebGLCapabilities{p: p}
}
