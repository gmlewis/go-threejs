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

// RawShaderMaterial works just like ShaderMaterial, except that definitions of built-in uniforms
// and attributes are not automatically prepended to the GLSL shader code.
//
// http://threejs.org/docs/index.html#Reference/Materials/RawShaderMaterial
type RawShaderMaterial struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (r *RawShaderMaterial) JSObject() *js.Object { return r.p }

// RawShaderMaterial returns a RawShaderMaterial JavaScript class.
func (t *Three) RawShaderMaterial() *RawShaderMaterial {
	p := t.ctx.Get("RawShaderMaterial")
	return RawShaderMaterialFromJSObject(p)
}

// RawShaderMaterialFromJSObject returns a wrapped RawShaderMaterial JavaScript class.
func RawShaderMaterialFromJSObject(p *js.Object) *RawShaderMaterial {
	return &RawShaderMaterial{p: p}
}

// NewRawShaderMaterial returns a new RawShaderMaterial object.
func (t *Three) NewRawShaderMaterial(parameters float64) *RawShaderMaterial {
	p := t.ctx.Get("RawShaderMaterial").New(parameters)
	return RawShaderMaterialFromJSObject(p)
}
