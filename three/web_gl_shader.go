// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLShader is a lower level function to compile either a vertex or fragment shader.
//
// http://threejs.org/docs/index.html#Reference/Renderers.WebGL/WebGLShader
type WebGLShader struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WebGLShader) JSObject() *js.Object { return w.p }

// WebGLShader returns a WebGLShader JavaScript class.
func (t *Three) WebGLShader() *WebGLShader {
	p := t.ctx.Get("WebGLShader")
	return &WebGLShader{p: p}
}

// NewWebGLShader returns a new WebGLShader object.
func (t *Three) NewWebGLShader() *WebGLShader {
	p := t.ctx.Get("WebGLShader").New()
	return &WebGLShader{p: p}
}

// TODO Objects
