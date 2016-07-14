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
func (w *WebGLLights) JSObject() *js.Object { return w.p }

// WebGLLights returns a WebGLLights JavaScript class.
func (t *Three) WebGLLights() *WebGLLights {
	p := t.ctx.Get("WebGLLights")
	return WebGLLightsFromJSObject(p)
}

// WebGLLightsFromJSObject returns a wrapped WebGLLights JavaScript class.
func WebGLLightsFromJSObject(p *js.Object) *WebGLLights {
	return &WebGLLights{p: p}
}

// NewWebGLLights returns a new WebGLLights object.
func (t *Three) NewWebGLLights() *WebGLLights {
	p := t.ctx.Get("WebGLLights").New()
	return WebGLLightsFromJSObject(p)
}
