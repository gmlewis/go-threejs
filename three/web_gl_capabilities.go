// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLCapabilities represents a webglcapabilities.
type WebGLCapabilities struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WebGLCapabilities) JSObject() *js.Object { return w.p }

// WebGLCapabilities returns a WebGLCapabilities JavaScript class.
func (t *Three) WebGLCapabilities() *WebGLCapabilities {
	p := t.ctx.Get("WebGLCapabilities")
	return &WebGLCapabilities{p: p}
}

// NewWebGLCapabilities returns a new WebGLCapabilities object.
func (t *Three) NewWebGLCapabilities(gl, extensions, parameters float64) *WebGLCapabilities {
	p := t.ctx.Get("WebGLCapabilities").New(gl, extensions, parameters)
	return &WebGLCapabilities{p: p}
}
