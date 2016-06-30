// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLShader represents a webglshader.
type WebGLShader struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *WebGLShader) JSObject() *js.Object { return t.p }

// WebGLShader returns a WebGLShader object.
func (t *Three) WebGLShader() *WebGLShader {
	p := t.ctx.Get("WebGLShader")
	return &WebGLShader{p: p}
}

// New returns a new WebGLShader object.
func (t *WebGLShader) New() *WebGLShader {
	p := t.p.New()
	return &WebGLShader{p: p}
}
