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

// ShaderLib represents a shaderlib.
//
// http://threejs.org/docs/index.html#Reference/Renderers.Shaders/ShaderLib
type ShaderLib struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (s *ShaderLib) JSObject() *js.Object { return s.p }

// ShaderLib returns a ShaderLib JavaScript class.
func (t *Three) ShaderLib() *ShaderLib {
	p := t.ctx.Get("ShaderLib")
	return ShaderLibFromJSObject(p)
}

// ShaderLibFromJSObject returns a wrapped ShaderLib JavaScript class.
func ShaderLibFromJSObject(p *js.Object) *ShaderLib {
	return &ShaderLib{p: p}
}

// NewShaderLib returns a new ShaderLib object.
func (t *Three) NewShaderLib() *ShaderLib {
	p := t.ctx.Get("ShaderLib").New()
	return ShaderLibFromJSObject(p)
}
