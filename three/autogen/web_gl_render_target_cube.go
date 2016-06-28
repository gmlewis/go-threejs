package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLRenderTargetCube represents a webglrendertargetcube.
type WebGLRenderTargetCube struct{ p *js.Object }

// WebGLRenderTargetCube returns a webglrendertargetcube object.
func (t *Three) WebGLRenderTargetCube() *WebGLRenderTargetCube {
	p := t.ctx.Get("WebGLRenderTargetCube")
	return &WebGLRenderTargetCube{p: p}
}

// NewWebGLRenderTargetCube returns a new webglrendertargetcube object.
func (t *WebGLRenderTargetCube) New(width, height, options float64) *WebGLRenderTargetCube {
	p := t.p.New(width, height, options)
	return &WebGLRenderTargetCube{p: p}
}

