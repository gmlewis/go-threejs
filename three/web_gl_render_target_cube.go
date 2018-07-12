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

// WebGLRenderTargetCube represents a webglrendertargetcube. CubeCamera uses this as its WebGLRenderTarget.
//
// http://threejs.org/docs/index.html#Reference/Renderers/WebGLRenderTargetCube
type WebGLRenderTargetCube struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WebGLRenderTargetCube) JSObject() *js.Object { return w.p }

// WebGLRenderTargetCube returns a WebGLRenderTargetCube JavaScript class.
func (t *Three) WebGLRenderTargetCube() *WebGLRenderTargetCube {
	p := t.ctx.Get("WebGLRenderTargetCube")
	return WebGLRenderTargetCubeFromJSObject(p)
}

// WebGLRenderTargetCubeFromJSObject returns a wrapped WebGLRenderTargetCube JavaScript class.
func WebGLRenderTargetCubeFromJSObject(p *js.Object) *WebGLRenderTargetCube {
	return &WebGLRenderTargetCube{p: p}
}

// NewWebGLRenderTargetCube returns a new WebGLRenderTargetCube object.
//
//     width -- The width of the renderTarget.
//     height -- The height of the renderTarget.
//     options -- The options sets the properties of the render target.
func (t *Three) NewWebGLRenderTargetCube(width, height, options float64) *WebGLRenderTargetCube {
	p := t.ctx.Get("WebGLRenderTargetCube").New(width, height, options)
	return WebGLRenderTargetCubeFromJSObject(p)
}

// TODO: ActiveCubeFace
