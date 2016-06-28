package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLRenderTarget represents a webglrendertarget.
type WebGLRenderTarget struct{ p *js.Object }

// WebGLRenderTarget returns a webglrendertarget object.
func (t *Three) WebGLRenderTarget() *WebGLRenderTarget {
	p := t.ctx.Get("WebGLRenderTarget")
	return &WebGLRenderTarget{p: p}
}

// NewWebGLRenderTarget returns a new webglrendertarget object.
func (t *WebGLRenderTarget) New(width, height, options float64) *WebGLRenderTarget {
	p := t.p.New(width, height, options)
	return &WebGLRenderTarget{p: p}
}

// SetSize TODO description.
func (w *WebGLRenderTarget) SetSize(width, height float64) *WebGLRenderTarget {
	w.p.Call("setSize", width, height)
	return w
}

// Clone TODO description.
func (w *WebGLRenderTarget) Clone() *WebGLRenderTarget {
	w.p.Call("clone")
	return w
}

// Copy TODO description.
func (w *WebGLRenderTarget) Copy(source float64) *WebGLRenderTarget {
	w.p.Call("copy", source)
	return w
}

// Dispose TODO description.
func (w *WebGLRenderTarget) Dispose() *WebGLRenderTarget {
	w.p.Call("dispose")
	return w
}

