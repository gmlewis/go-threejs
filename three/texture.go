// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Texture creates a texture to apply to a surface or as a reflection or refraction map.
//
// http://threejs.org/docs/index.html#Reference/Textures/Texture
type Texture struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *Texture) JSObject() *js.Object { return t.p }

// Texture returns a Texture JavaScript class.
func (t *Three) Texture() *Texture {
	p := t.ctx.Get("Texture")
	return &Texture{p: p}
}

// NewTexture returns a new Texture object.
func (t *Three) NewTexture(image, mapping, wrapS, wrapT, magFilter, minFilter, format, typ, anisotropy float64) *Texture {
	p := t.ctx.Get("Texture").New(image, mapping, wrapS, wrapT, magFilter, minFilter, format, typ, anisotropy)
	return &Texture{p: p}
}

// SetNeedsUpdate sets the needsUpdate-component of the Texture.
func (t *Texture) SetNeedsUpdate(value float64) *Texture {
	t.p.Set("needsUpdate", value)
	return t
}

// Clone makes a copy of texture. Note this is not a "deep copy", the image is shared.
func (t *Texture) Clone() *Texture {
	t.p.Call("clone")
	return t
}

// Copy TODO description.
func (t *Texture) Copy(source *Texture) *Texture {
	t.p.Call("copy", source.p)
	return t
}

// ToJSON TODO description.
func (t *Texture) ToJSON(meta float64) *Texture {
	t.p.Call("toJSON", meta)
	return t
}

// Dispose calls EventDispatcher.dispatchEvent with a 'dispose' event type.
func (t *Texture) Dispose() *Texture {
	t.p.Call("dispose")
	return t
}

// TransformUv TODO description.
func (t *Texture) TransformUv(uv float64) *Texture {
	t.p.Call("transformUv", uv)
	return t
}
