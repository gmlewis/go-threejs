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
