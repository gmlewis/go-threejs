// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLRenderTargetCube represents a webglrendertargetcube.
type WebGLRenderTargetCube struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WebGLRenderTargetCube) JSObject() *js.Object { return w.p }

// WebGLRenderTargetCube returns a WebGLRenderTargetCube object.
func (t *Three) WebGLRenderTargetCube() *WebGLRenderTargetCube {
	p := t.ctx.Get("WebGLRenderTargetCube")
	return &WebGLRenderTargetCube{p: p}
}

// New returns a new WebGLRenderTargetCube object.
func (w *WebGLRenderTargetCube) New(width, height, options float64) *WebGLRenderTargetCube {
	p := w.p.New(width, height, options)
	return &WebGLRenderTargetCube{p: p}
}
