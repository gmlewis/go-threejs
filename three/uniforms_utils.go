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

// UniformsUtils support the merging and cloning of uniform variables.
//
// http://threejs.org/docs/index.html#Reference/Renderers.Shaders/UniformsUtils
type UniformsUtils struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (u *UniformsUtils) JSObject() *js.Object { return u.p }

// UniformsUtils returns an UniformsUtils JavaScript class.
func (t *Three) UniformsUtils() *UniformsUtils {
	p := t.ctx.Get("UniformsUtils")
	return UniformsUtilsFromJSObject(p)
}

// UniformsUtilsFromJSObject returns a wrapped UniformsUtils JavaScript class.
func UniformsUtilsFromJSObject(p *js.Object) *UniformsUtils {
	return &UniformsUtils{p: p}
}

// NewUniformsUtils returns a new UniformsUtils object.
func (t *Three) NewUniformsUtils() *UniformsUtils {
	p := t.ctx.Get("UniformsUtils").New()
	return UniformsUtilsFromJSObject(p)
}

// Merge TODO description.
func (u *UniformsUtils) Merge(uniforms float64) *UniformsUtils {
	u.p.Call("merge", uniforms)
	return u
}

// Clone TODO description.
func (u *UniformsUtils) Clone(uniformsSrc JSObject) *UniformsUtils {
	u.p.Call("clone", uniformsSrc.JSObject())
	return u
}
