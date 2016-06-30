// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLGeometries represents a webglgeometries.
type WebGLGeometries struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WebGLGeometries) JSObject() *js.Object { return w.p }

// WebGLGeometries returns a WebGLGeometries object.
func (t *Three) WebGLGeometries() *WebGLGeometries {
	p := t.ctx.Get("WebGLGeometries")
	return &WebGLGeometries{p: p}
}

// New returns a new WebGLGeometries object.
func (w *WebGLGeometries) New(gl, properties, info float64) *WebGLGeometries {
	p := w.p.New(gl, properties, info)
	return &WebGLGeometries{p: p}
}
