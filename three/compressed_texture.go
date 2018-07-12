// Copyright 2016 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CompressedTexture represents a texture based on data in compressed form.
//
// http://threejs.org/docs/index.html#Reference/Textures/CompressedTexture
type CompressedTexture struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (c *CompressedTexture) JSObject() *js.Object { return c.p }

// CompressedTexture returns a CompressedTexture JavaScript class.
func (t *Three) CompressedTexture() *CompressedTexture {
	p := t.ctx.Get("CompressedTexture")
	return CompressedTextureFromJSObject(p)
}

// CompressedTextureFromJSObject returns a wrapped CompressedTexture JavaScript class.
func CompressedTextureFromJSObject(p *js.Object) *CompressedTexture {
	return &CompressedTexture{p: p}
}

// NewCompressedTexture returns a new CompressedTexture object.
func (t *Three) NewCompressedTexture(mipmaps, width, height, format, typ, mapping, wrapS, wrapT, magFilter, minFilter, anisotropy float64) *CompressedTexture {
	p := t.ctx.Get("CompressedTexture").New(mipmaps, width, height, format, typ, mapping, wrapS, wrapT, magFilter, minFilter, anisotropy)
	return CompressedTextureFromJSObject(p)
}
