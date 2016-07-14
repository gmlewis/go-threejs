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

// WebGLGeometries returns a WebGLGeometries JavaScript class.
func (t *Three) WebGLGeometries() *WebGLGeometries {
	p := t.ctx.Get("WebGLGeometries")
	return WebGLGeometriesFromJSObject(p)
}

// WebGLGeometriesFromJSObject returns a wrapped WebGLGeometries JavaScript class.
func WebGLGeometriesFromJSObject(p *js.Object) *WebGLGeometries {
	return &WebGLGeometries{p: p}
}

// NewWebGLGeometries returns a new WebGLGeometries object.
func (t *Three) NewWebGLGeometries(gl, properties, info float64) *WebGLGeometries {
	p := t.ctx.Get("WebGLGeometries").New(gl, properties, info)
	return WebGLGeometriesFromJSObject(p)
}
