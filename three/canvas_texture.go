// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CanvasTexture represents a canvastexture.
type CanvasTexture struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (c *CanvasTexture) JSObject() *js.Object { return c.p }

// CanvasTexture returns a CanvasTexture JavaScript class.
func (t *Three) CanvasTexture() *CanvasTexture {
	p := t.ctx.Get("CanvasTexture")
	return CanvasTextureFromJSObject(p)
}

// CanvasTextureFromJSObject returns a wrapped CanvasTexture JavaScript class.
func CanvasTextureFromJSObject(p *js.Object) *CanvasTexture {
	return &CanvasTexture{p: p}
}

// NewCanvasTexture returns a new CanvasTexture object.
func (t *Three) NewCanvasTexture(canvas, mapping, wrapS, wrapT, magFilter, minFilter, format, typ, anisotropy float64) *CanvasTexture {
	p := t.ctx.Get("CanvasTexture").New(canvas, mapping, wrapS, wrapT, magFilter, minFilter, format, typ, anisotropy)
	return CanvasTextureFromJSObject(p)
}
