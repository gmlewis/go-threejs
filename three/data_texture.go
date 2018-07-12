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
