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

// CubeTexture represents a cube texture made up of six images.
//
// http://threejs.org/docs/index.html#Reference/Textures/CubeTexture
type CubeTexture struct{ *Texture }

// JSObject returns the underlying *js.Object.
func (c *CubeTexture) JSObject() *js.Object { return c.p }

// CubeTexture returns a CubeTexture JavaScript class.
func (t *Three) CubeTexture() *CubeTexture {
	p := t.ctx.Get("CubeTexture")
	return CubeTextureFromJSObject(p)
}

// CubeTextureFromJSObject returns a wrapped CubeTexture JavaScript class.
func CubeTextureFromJSObject(p *js.Object) *CubeTexture {
	return &CubeTexture{TextureFromJSObject(p)}
}

// NewCubeTexture returns a new CubeTexture object.
func (t *Three) NewCubeTexture(images, mapping, wrapS, wrapT, magFilter, minFilter, format, typ, anisotropy float64) *CubeTexture {
	p := t.ctx.Get("CubeTexture").New(images, mapping, wrapS, wrapT, magFilter, minFilter, format, typ, anisotropy)
	return CubeTextureFromJSObject(p)
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
