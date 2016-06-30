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
func (t *WebGLBufferRenderer) JSObject() *js.Object { return t.p }

// WebGLBufferRenderer returns a WebGLBufferRenderer object.
func (t *Three) WebGLBufferRenderer() *WebGLBufferRenderer {
	p := t.ctx.Get("WebGLBufferRenderer")
	return &WebGLBufferRenderer{p: p}
}

// New returns a new WebGLBufferRenderer object.
func (t *WebGLBufferRenderer) New(_gl, extensions, _infoRender float64) *WebGLBufferRenderer {
	p := t.p.New(_gl, extensions, _infoRender)
	return &WebGLBufferRenderer{p: p}
}