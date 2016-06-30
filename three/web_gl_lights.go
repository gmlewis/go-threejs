// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLLights represents a webgllights.
type WebGLLights struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *WebGLLights) JSObject() *js.Object { return t.p }

// WebGLLights returns a WebGLLights object.
func (t *Three) WebGLLights() *WebGLLights {
	p := t.ctx.Get("WebGLLights")
	return &WebGLLights{p: p}
}

// New returns a new WebGLLights object.
func (t *WebGLLights) New() *WebGLLights {
	p := t.p.New()
	return &WebGLLights{p: p}
}
