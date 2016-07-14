// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLBufferRenderer represents a webglbufferrenderer.
type WebGLBufferRenderer struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WebGLBufferRenderer) JSObject() *js.Object { return w.p }

// WebGLBufferRenderer returns a WebGLBufferRenderer JavaScript class.
func (t *Three) WebGLBufferRenderer() *WebGLBufferRenderer {
	p := t.ctx.Get("WebGLBufferRenderer")
	return WebGLBufferRendererFromJSObject(p)
}

// WebGLBufferRendererFromJSObject returns a wrapped WebGLBufferRenderer JavaScript class.
func WebGLBufferRendererFromJSObject(p *js.Object) *WebGLBufferRenderer {
	return &WebGLBufferRenderer{p: p}
}

// NewWebGLBufferRenderer returns a new WebGLBufferRenderer object.
func (t *Three) NewWebGLBufferRenderer(_gl, extensions, _infoRender float64) *WebGLBufferRenderer {
	p := t.ctx.Get("WebGLBufferRenderer").New(_gl, extensions, _infoRender)
	return WebGLBufferRendererFromJSObject(p)
}
