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

// Sprite represents a plane in an 3d scene which faces always towards the camera.
//
// http://threejs.org/docs/index.html#Reference/Objects/Sprite
type Sprite struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (s *Sprite) JSObject() *js.Object { return s.p }

// Sprite returns a Sprite JavaScript class.
func (t *Three) Sprite() *Sprite {
	p := t.ctx.Get("Sprite")
	return SpriteFromJSObject(p)
}

// SpriteFromJSObject returns a wrapped Sprite JavaScript class.
func SpriteFromJSObject(p *js.Object) *Sprite {
	return &Sprite{p: p}
}

// NewSprite returns a new Sprite object.
//
//     material — An instance of Material (optional).
func (t *Three) NewSprite(material *js.Object) *Sprite {
	p := t.ctx.Get("Sprite").New(material)
	return SpriteFromJSObject(p)
}
