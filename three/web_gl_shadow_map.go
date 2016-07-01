// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLShadowMap represents a webglshadowmap.
type WebGLShadowMap struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WebGLShadowMap) JSObject() *js.Object { return w.p }

// WebGLShadowMap returns a WebGLShadowMap JavaScript class.
func (t *Three) WebGLShadowMap() *WebGLShadowMap {
	p := t.ctx.Get("WebGLShadowMap")
	return &WebGLShadowMap{p: p}
}

// NewWebGLShadowMap returns a new WebGLShadowMap object.
func (t *Three) NewWebGLShadowMap(_renderer, _lights, _objects float64) *WebGLShadowMap {
	p := t.ctx.Get("WebGLShadowMap").New(_renderer, _lights, _objects)
	return &WebGLShadowMap{p: p}
}
