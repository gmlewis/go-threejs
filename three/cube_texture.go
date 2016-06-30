// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CubeTexture represents a cubetexture.
type CubeTexture struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (c *CubeTexture) JSObject() *js.Object { return c.p }

// CubeTexture returns a CubeTexture object.
func (t *Three) CubeTexture() *CubeTexture {
	p := t.ctx.Get("CubeTexture")
	return &CubeTexture{p: p}
}

// New returns a new CubeTexture object.
func (c *CubeTexture) New(images, mapping, wrapS, wrapT, magFilter, minFilter, format, typ, anisotropy float64) *CubeTexture {
	p := c.p.New(images, mapping, wrapS, wrapT, magFilter, minFilter, format, typ, anisotropy)
	return &CubeTexture{p: p}
}

// Get TODO description.
func (c *CubeTexture) Get() *CubeTexture {
	c.p.Call("get")
	return c
}

// Set TODO description.
func (c *CubeTexture) Set(value float64) *CubeTexture {
	c.p.Call("set", value)
	return c
}
