// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLProgram represents a webglprogram.
type WebGLProgram struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WebGLProgram) JSObject() *js.Object { return w.p }

// WebGLProgram returns a WebGLProgram JavaScript class.
func (t *Three) WebGLProgram() *WebGLProgram {
	p := t.ctx.Get("WebGLProgram")
	return WebGLProgramFromJSObject(p)
}

// WebGLProgramFromJSObject returns a wrapped WebGLProgram JavaScript class.
func WebGLProgramFromJSObject(p *js.Object) *WebGLProgram {
	return &WebGLProgram{p: p}
}

// NewWebGLProgram returns a new WebGLProgram object.
func (t *Three) NewWebGLProgram() *WebGLProgram {
	p := t.ctx.Get("WebGLProgram").New()
	return WebGLProgramFromJSObject(p)
}

// Get TODO description.
func (w *WebGLProgram) Get() *WebGLProgram {
	w.p.Call("get")
	return w
}
