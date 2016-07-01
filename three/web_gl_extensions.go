// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLExtensions represents a webglextensions.
type WebGLExtensions struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WebGLExtensions) JSObject() *js.Object { return w.p }

// WebGLExtensions returns a WebGLExtensions JavaScript class.
func (t *Three) WebGLExtensions() *WebGLExtensions {
	p := t.ctx.Get("WebGLExtensions")
	return &WebGLExtensions{p: p}
}

// NewWebGLExtensions returns a new WebGLExtensions object.
func (t *Three) NewWebGLExtensions(gl float64) *WebGLExtensions {
	p := t.ctx.Get("WebGLExtensions").New(gl)
	return &WebGLExtensions{p: p}
}
