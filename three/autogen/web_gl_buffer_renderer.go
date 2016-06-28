package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLBufferRenderer represents a webglbufferrenderer.
type WebGLBufferRenderer struct{ p *js.Object }

// WebGLBufferRenderer returns a webglbufferrenderer object.
func (t *Three) WebGLBufferRenderer() *WebGLBufferRenderer {
	p := t.ctx.Get("WebGLBufferRenderer")
	return &WebGLBufferRenderer{p: p}
}

// NewWebGLBufferRenderer returns a new webglbufferrenderer object.
func (t *WebGLBufferRenderer) New(_gl, extensions, _infoRender float64) *WebGLBufferRenderer {
	p := t.p.New(_gl, extensions, _infoRender)
	return &WebGLBufferRenderer{p: p}
}

