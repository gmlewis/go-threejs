// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLState represents a webglstate.
type WebGLState struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WebGLState) JSObject() *js.Object { return w.p }

// WebGLState returns a WebGLState object.
func (t *Three) WebGLState() *WebGLState {
	p := t.ctx.Get("WebGLState")
	return &WebGLState{p: p}
}

// New returns a new WebGLState object.
func (w *WebGLState) New(gl, extensions, paramThreeToGL float64) *WebGLState {
	p := w.p.New(gl, extensions, paramThreeToGL)
	return &WebGLState{p: p}
}
