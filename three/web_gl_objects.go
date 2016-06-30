// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLObjects represents a webglobjects.
type WebGLObjects struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WebGLObjects) JSObject() *js.Object { return w.p }

// WebGLObjects returns a WebGLObjects object.
func (t *Three) WebGLObjects() *WebGLObjects {
	p := t.ctx.Get("WebGLObjects")
	return &WebGLObjects{p: p}
}

// New returns a new WebGLObjects object.
func (w *WebGLObjects) New(gl, properties, info float64) *WebGLObjects {
	p := w.p.New(gl, properties, info)
	return &WebGLObjects{p: p}
}
