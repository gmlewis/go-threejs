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

// WebGLState represents a webglstate.
//
// http://threejs.org/docs/index.html#Reference/Renderers.WebGL/WebGLState
type WebGLState struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WebGLState) JSObject() *js.Object { return w.p }

// WebGLState returns a WebGLState JavaScript class.
func (t *Three) WebGLState() *WebGLState {
	p := t.ctx.Get("WebGLState")
	return WebGLStateFromJSObject(p)
}

// WebGLStateFromJSObject returns a wrapped WebGLState JavaScript class.
func WebGLStateFromJSObject(p *js.Object) *WebGLState {
	return &WebGLState{p: p}
}

// NewWebGLState returns a new WebGLState object.
func (t *Three) NewWebGLState(gl, extensions, paramThreeToGL float64) *WebGLState {
	p := t.ctx.Get("WebGLState").New(gl, extensions, paramThreeToGL)
	return WebGLStateFromJSObject(p)
}

// TODO: setDepthTest, setDepthWrite, setBlending, enable, disable
