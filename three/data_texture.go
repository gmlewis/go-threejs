// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// DataTexture represents a a texture directly from raw data, width and height.
//
// http://threejs.org/docs/index.html#Reference/Textures/DataTexture
type DataTexture struct{ *Texture }

// JSObject returns the underlying *js.Object.
func (d *DataTexture) JSObject() *js.Object { return d.p }

// DataTexture returns a DataTexture JavaScript class.
func (t *Three) DataTexture() *DataTexture {
	p := t.ctx.Get("DataTexture")
	return DataTextureFromJSObject(p)
}

// DataTextureFromJSObject returns a wrapped DataTexture JavaScript class.
func DataTextureFromJSObject(p *js.Object) *DataTexture {
	return &DataTexture{TextureFromJSObject(p)}
}

// NewDataTexture returns a new DataTexture object.
func (t *Three) NewDataTexture(data, width, height, format, typ, mapping, wrapS, wrapT, magFilter, minFilter, anisotropy float64) *DataTexture {
	p := t.ctx.Get("DataTexture").New(data, width, height, format, typ, mapping, wrapS, wrapT, magFilter, minFilter, anisotropy)
	return DataTextureFromJSObject(p)
}
