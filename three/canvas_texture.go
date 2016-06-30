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

// CanvasTexture returns a CanvasTexture object.
func (t *Three) CanvasTexture() *CanvasTexture {
	p := t.ctx.Get("CanvasTexture")
	return &CanvasTexture{p: p}
}

// New returns a new CanvasTexture object.
func (c *CanvasTexture) New(canvas, mapping, wrapS, wrapT, magFilter, minFilter, format, typ, anisotropy float64) *CanvasTexture {
	p := c.p.New(canvas, mapping, wrapS, wrapT, magFilter, minFilter, format, typ, anisotropy)
	return &CanvasTexture{p: p}
}
