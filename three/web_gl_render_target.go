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

// WebGLRenderTarget is a buffer where the video card draws pixels for a scene that is being
// rendered in the background. It is used in different effects.
//
// http://threejs.org/docs/index.html#Reference/Renderers/WebGLRenderTarget
type WebGLRenderTarget struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WebGLRenderTarget) JSObject() *js.Object { return w.p }

// WebGLRenderTarget returns a WebGLRenderTarget JavaScript class.
func (t *Three) WebGLRenderTarget() *WebGLRenderTarget {
	p := t.ctx.Get("WebGLRenderTarget")
	return WebGLRenderTargetFromJSObject(p)
}

// WebGLRenderTargetFromJSObject returns a wrapped WebGLRenderTarget JavaScript class.
func WebGLRenderTargetFromJSObject(p *js.Object) *WebGLRenderTarget {
	return &WebGLRenderTarget{p: p}
}

// NewWebGLRenderTarget creates a new render target with a certain width and height.
//
//     width -- The width of the renderTarget.
//     height -- The height of the renderTarget.
//
// options is an optional object that holds texture parameters for an auto-generated target
// texture and depthBuffer/stencilBuffer booleans. For an explanation of the texture parameters see Texture.
//
//     wrapS — Number default is THREE.ClampToEdgeWrapping.
//     wrapT — Number default is THREE.ClampToEdgeWrapping.
//     magFilter — Number, default is THREE.LinearFilter.
//     minFilter — Number, default is THREE.LinearFilter.
//     format — Number, default is THREE.RGBAFormat.
//     type — Number, default is THREE.UnsignedByteType.
//     anisotropy — Number, default is 1.
//     encoding — Number, default is THREE.LinearEncoding.
//     depthBuffer — Boolean, default is true. Set this to false if you don't need it.
//     stencilBuffer — Boolean, default is true. Set this to false if you don't need it.
func (t *Three) NewWebGLRenderTarget(width, height, options float64) *WebGLRenderTarget {
	p := t.ctx.Get("WebGLRenderTarget").New(width, height, options)
	return WebGLRenderTargetFromJSObject(p)
}

// SetSize sets the size of the render target.
func (w *WebGLRenderTarget) SetSize(width, height float64) *WebGLRenderTarget {
	w.p.Call("setSize", width, height)
	return w
}

// Clone creates a copy of this render target.
func (w *WebGLRenderTarget) Clone() *WebGLRenderTarget {
	w.p.Call("clone")
	return w
}

// Copy adopts the settings of the given render target.
func (w *WebGLRenderTarget) Copy(source *WebGLRenderTarget) *WebGLRenderTarget {
	w.p.Call("copy", source.p)
	return w
}

// Dispose dispatches a dispose event..
func (w *WebGLRenderTarget) Dispose() *WebGLRenderTarget {
	w.p.Call("dispose")
	return w
}

/* TODO
Uuid is a unique number for this render target instance.
Width is the width of the render target.
Height is the height of the render target.
Scissor is a rectangular area inside the render target's viewport. Fragments that are outside the area will be discarded.
ScissorTest indicates whether the scissor test is active or not.
Viewport is the viewport of this render target.
Texture holds the rendered pixels. Use it as input for further processing.
DepthBuffer renders to the depth buffer. Default is true.
StencilBuffer renders to the stencil buffer. Default is true.
DepthTexture if set, the scene depth will be rendered to this texture. Default is null.
*/
