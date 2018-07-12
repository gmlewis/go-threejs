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

// WebGLShader is a lower level function to compile either a vertex or fragment shader.
//
// http://threejs.org/docs/index.html#Reference/Renderers.WebGL/WebGLShader
type WebGLShader struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WebGLShader) JSObject() *js.Object { return w.p }

// WebGLShader returns a WebGLShader JavaScript class.
func (t *Three) WebGLShader() *WebGLShader {
	p := t.ctx.Get("WebGLShader")
	return WebGLShaderFromJSObject(p)
}

// WebGLShaderFromJSObject returns a wrapped WebGLShader JavaScript class.
func WebGLShaderFromJSObject(p *js.Object) *WebGLShader {
	return &WebGLShader{p: p}
}

// NewWebGLShader returns a new WebGLShader object.
func (t *Three) NewWebGLShader() *WebGLShader {
	p := t.ctx.Get("WebGLShader").New()
	return WebGLShaderFromJSObject(p)
}

// TODO Objects
