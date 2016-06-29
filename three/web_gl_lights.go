package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLLights represents a webgllights.
type WebGLLights struct{ p *js.Object }

// WebGLLights returns a WebGLLights object.
func (t *Three) WebGLLights() *WebGLLights {
	p := t.ctx.Get("WebGLLights")
	return &WebGLLights{p: p}
}

// New returns a new WebGLLights object.
func (t *WebGLLights) New() *WebGLLights {
	p := t.p.New()
	return &WebGLLights{p: p}
}

