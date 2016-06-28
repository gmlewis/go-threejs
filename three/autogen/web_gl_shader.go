package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLShader represents a webglshader.
type WebGLShader struct{ p *js.Object }

// WebGLShader returns a WebGLShader object.
func (t *Three) WebGLShader() *WebGLShader {
	p := t.ctx.Get("WebGLShader")
	return &WebGLShader{p: p}
}

// New returns a new WebGLShader object.
func (t *WebGLShader) New() *WebGLShader {
	p := t.p.New()
	return &WebGLShader{p: p}
}

