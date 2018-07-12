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

// SpriteMaterial represents a material for a sprite.
//
// http://threejs.org/docs/index.html#Reference/Materials/SpriteMaterial
type SpriteMaterial struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (s *SpriteMaterial) JSObject() *js.Object { return s.p }

// SpriteMaterial returns a SpriteMaterial JavaScript class.
func (t *Three) SpriteMaterial() *SpriteMaterial {
	p := t.ctx.Get("SpriteMaterial")
	return SpriteMaterialFromJSObject(p)
}

// SpriteMaterialFromJSObject returns a wrapped SpriteMaterial JavaScript class.
func SpriteMaterialFromJSObject(p *js.Object) *SpriteMaterial {
	return &SpriteMaterial{p: p}
}

// NewSpriteMaterial returns a new SpriteMaterial object.
//
// parameters is an object defining the the default properties:
//     color - color of the sprite
//     map - the texture map
//     rotation - the rotation of the sprite
//     fog - whether or not to use the scene fog
func (t *Three) NewSpriteMaterial(parameters map[string]interface{}) *SpriteMaterial {
	p := t.ctx.Get("SpriteMaterial").New(parameters)
	return SpriteMaterialFromJSObject(p)
}

// Copy TODO description.
func (s *SpriteMaterial) Copy(source *SpriteMaterial) *SpriteMaterial {
	s.p.Call("copy", source.p)
	return s
}
