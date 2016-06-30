// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLRenderTarget represents a webglrendertarget.
type WebGLRenderTarget struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *WebGLRenderTarget) JSObject() *js.Object { return t.p }

// WebGLRenderTarget returns a WebGLRenderTarget object.
func (t *Three) WebGLRenderTarget() *WebGLRenderTarget {
	p := t.ctx.Get("WebGLRenderTarget")
	return &WebGLRenderTarget{p: p}
}

// New returns a new WebGLRenderTarget object.
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
