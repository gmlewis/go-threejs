// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLState represents a webglstate.
//
// http://threejs.org/docs/index.html#Reference/Renderers.WebGL/WebGLState
type WebGLState struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WebGLState) JSObject() *js.Object { return w.p }

// WebGLState returns a WebGLState JavaScript class.
func (t *Three) WebGLState() *WebGLState {
	p := t.ctx.Get("WebGLState")
	return WebGLStateFromJSObject(p)
}

// WebGLStateFromJSObject returns a wrapped WebGLState JavaScript class.
func WebGLStateFromJSObject(p *js.Object) *WebGLState {
	return &WebGLState{p: p}
}

// NewWebGLState returns a new WebGLState object.
func (t *Three) NewWebGLState(gl, extensions, paramThreeToGL float64) *WebGLState {
	p := t.ctx.Get("WebGLState").New(gl, extensions, paramThreeToGL)
	return WebGLStateFromJSObject(p)
}

// TODO: setDepthTest, setDepthWrite, setBlending, enable, disable
