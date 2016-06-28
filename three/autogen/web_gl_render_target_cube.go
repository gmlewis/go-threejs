package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLRenderTargetCube represents a webglrendertargetcube.
type WebGLRenderTargetCube struct{ p *js.Object }

// WebGLRenderTargetCube returns a WebGLRenderTargetCube object.
func (t *Three) WebGLRenderTargetCube() *WebGLRenderTargetCube {
	p := t.ctx.Get("WebGLRenderTargetCube")
	return &WebGLRenderTargetCube{p: p}
}

// New returns a new WebGLRenderTargetCube object.
func (t *WebGLRenderTargetCube) New(width, height, options float64) *WebGLRenderTargetCube {
	p := t.p.New(width, height, options)
	return &WebGLRenderTargetCube{p: p}
}

