// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CompressedTexture represents a compressedtexture.
type CompressedTexture struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *CompressedTexture) JSObject() *js.Object { return t.p }

// CompressedTexture returns a CompressedTexture object.
func (t *Three) CompressedTexture() *CompressedTexture {
	p := t.ctx.Get("CompressedTexture")
	return &CompressedTexture{p: p}
}

// New returns a new CompressedTexture object.
func (t *CompressedTexture) New(mipmaps, width, height, format, typ, mapping, wrapS, wrapT, magFilter, minFilter, anisotropy float64) *CompressedTexture {
	p := t.p.New(mipmaps, width, height, format, typ, mapping, wrapS, wrapT, magFilter, minFilter, anisotropy)
	return &CompressedTexture{p: p}
}
