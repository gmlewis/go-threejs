// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLProperties represents a webglproperties.
type WebGLProperties struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WebGLProperties) JSObject() *js.Object { return w.p }

// WebGLProperties returns a WebGLProperties JavaScript class.
func (t *Three) WebGLProperties() *WebGLProperties {
	p := t.ctx.Get("WebGLProperties")
	return &WebGLProperties{p: p}
}

// NewWebGLProperties returns a new WebGLProperties object.
func (t *Three) NewWebGLProperties() *WebGLProperties {
	p := t.ctx.Get("WebGLProperties").New()
	return &WebGLProperties{p: p}
}
