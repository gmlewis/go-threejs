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
func (t *WebGLExtensions) JSObject() *js.Object { return t.p }

// WebGLExtensions returns a WebGLExtensions object.
func (t *Three) WebGLExtensions() *WebGLExtensions {
	p := t.ctx.Get("WebGLExtensions")
	return &WebGLExtensions{p: p}
}

// New returns a new WebGLExtensions object.
func (t *WebGLExtensions) New(gl float64) *WebGLExtensions {
	p := t.p.New(gl)
	return &WebGLExtensions{p: p}
}
