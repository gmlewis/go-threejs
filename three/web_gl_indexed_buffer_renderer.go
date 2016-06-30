// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLIndexedBufferRenderer represents a webglindexedbufferrenderer.
type WebGLIndexedBufferRenderer struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *WebGLIndexedBufferRenderer) JSObject() *js.Object { return t.p }

// WebGLIndexedBufferRenderer returns a WebGLIndexedBufferRenderer object.
func (t *Three) WebGLIndexedBufferRenderer() *WebGLIndexedBufferRenderer {
	p := t.ctx.Get("WebGLIndexedBufferRenderer")
	return &WebGLIndexedBufferRenderer{p: p}
}

// New returns a new WebGLIndexedBufferRenderer object.
func (t *WebGLIndexedBufferRenderer) New(_gl, extensions, _infoRender float64) *WebGLIndexedBufferRenderer {
	p := t.p.New(_gl, extensions, _infoRender)
	return &WebGLIndexedBufferRenderer{p: p}
}
