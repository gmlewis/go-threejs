package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLShader represents a webglshader.
type WebGLShader struct{ p *js.Object }

// WebGLShader returns a webglshader object.
func (t *Three) WebGLShader() *WebGLShader {
	p := t.ctx.Get("WebGLShader")
	return &WebGLShader{p: p}
}

// NewWebGLShader returns a new webglshader object.
func (t *WebGLShader) New() *WebGLShader {
	p := t.p.New()
	return &WebGLShader{p: p}
}

