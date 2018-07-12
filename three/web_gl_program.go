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

// WebGLProgram represents a webglprogram. Constructor for the GLSL program sent to vertex
// and fragment shaders, including default uniforms and attributes.
//
// http://threejs.org/docs/index.html#Reference/Renderers.WebGL/WebGLProgram
type WebGLProgram struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WebGLProgram) JSObject() *js.Object { return w.p }

// WebGLProgram returns a WebGLProgram JavaScript class.
func (t *Three) WebGLProgram() *WebGLProgram {
	p := t.ctx.Get("WebGLProgram")
	return WebGLProgramFromJSObject(p)
}

// WebGLProgramFromJSObject returns a wrapped WebGLProgram JavaScript class.
func WebGLProgramFromJSObject(p *js.Object) *WebGLProgram {
	return &WebGLProgram{p: p}
}

// NewWebGLProgram returns a new WebGLProgram object.
func (t *Three) NewWebGLProgram() *WebGLProgram {
	p := t.ctx.Get("WebGLProgram").New()
	return WebGLProgramFromJSObject(p)
}

// Get TODO description.
func (w *WebGLProgram) Get() *WebGLProgram {
	w.p.Call("get")
	return w
}

/* TODO:
Id
Code
UsedTimes
Program
VertexShader
FragmentShader
GetUniforms returns a name-value mapping of all active uniform locations.
GetAttributes returns a name-value mapping of all active vertex attribute locations.
*/
