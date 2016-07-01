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

// WebGLRenderTargetCube returns a WebGLRenderTargetCube JavaScript class.
func (t *Three) WebGLRenderTargetCube() *WebGLRenderTargetCube {
	p := t.ctx.Get("WebGLRenderTargetCube")
	return &WebGLRenderTargetCube{p: p}
}

// NewWebGLRenderTargetCube returns a new WebGLRenderTargetCube object.
func (t *Three) NewWebGLRenderTargetCube(width, height, options float64) *WebGLRenderTargetCube {
	p := t.ctx.Get("WebGLRenderTargetCube").New(width, height, options)
	return &WebGLRenderTargetCube{p: p}
}
