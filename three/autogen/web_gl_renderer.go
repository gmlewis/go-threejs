package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLRenderer represents a webglrenderer.
type WebGLRenderer struct{ p *js.Object }

// WebGLRenderer returns a WebGLRenderer object.
func (t *Three) WebGLRenderer() *WebGLRenderer {
	p := t.ctx.Get("WebGLRenderer")
	return &WebGLRenderer{p: p}
}

// New returns a new WebGLRenderer object.
func (t *WebGLRenderer) New(parameters float64) *WebGLRenderer {
	p := t.p.New(parameters)
	return &WebGLRenderer{p: p}
}

