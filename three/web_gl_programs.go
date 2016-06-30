// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLPrograms represents a webglprograms.
type WebGLPrograms struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WebGLPrograms) JSObject() *js.Object { return w.p }

// WebGLPrograms returns a WebGLPrograms object.
func (t *Three) WebGLPrograms() *WebGLPrograms {
	p := t.ctx.Get("WebGLPrograms")
	return &WebGLPrograms{p: p}
}

// New returns a new WebGLPrograms object.
func (w *WebGLPrograms) New(renderer, capabilities float64) *WebGLPrograms {
	p := w.p.New(renderer, capabilities)
	return &WebGLPrograms{p: p}
}
