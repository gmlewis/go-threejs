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

// UniformsLib is a library for shared webgl shaders.
//
// http://threejs.org/docs/index.html#Reference/Renderers.Shaders/UniformsLib
type UniformsLib struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (u *UniformsLib) JSObject() *js.Object { return u.p }

// UniformsLib returns an UniformsLib JavaScript class.
func (t *Three) UniformsLib() *UniformsLib {
	p := t.ctx.Get("UniformsLib")
	return UniformsLibFromJSObject(p)
}

// UniformsLibFromJSObject returns a wrapped UniformsLib JavaScript class.
func UniformsLibFromJSObject(p *js.Object) *UniformsLib {
	return &UniformsLib{p: p}
}

// NewUniformsLib returns a new UniformsLib object.
func (t *Three) NewUniformsLib() *UniformsLib {
	p := t.ctx.Get("UniformsLib").New()
	return UniformsLibFromJSObject(p)
}
