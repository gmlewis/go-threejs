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

// WebGLRenderer displays your beautifully crafted scenes using WebGL, if your device supports it.
//
// http://threejs.org/docs/index.html#Reference/Renderers/WebGLRenderer
type WebGLRenderer struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WebGLRenderer) JSObject() *js.Object { return w.p }

// WebGLRenderer returns a WebGLRenderer JavaScript class.
func (t *Three) WebGLRenderer() *WebGLRenderer {
	p := t.ctx.Get("WebGLRenderer")
	return WebGLRendererFromJSObject(p)
}

// WebGLRendererOpts provides optional parameters to WebGLRenderer.
type WebGLRendererOpts struct {
	// canvas — A Canvas where the renderer draws its output.
	// context — The RenderingContext context to use.
	Precision              *string `json:"precision,omitempty"`             // Shader precision. Can be "highp", "mediump" or "lowp". Defaults to "highp" if supported by the device.
	Alpha                  *bool   `json:"alpha,omitempty"`                 // default is false.
	PremultipliedAlpha     *bool   `json:"premultipliedAlpha,omitempty"`    // default is true.
	Antialias              *bool   `json:"antialias,omitempty"`             // default is false.
	Stencil                *bool   `json:"stencil,omitempty"`               // default is true.
	PreserveDrawingBuffer  *bool   `json:"preserveDrawingBuffer,omitempty"` // default is false.
	Depth                  *bool   `json:"depth,omitempty"`                 // default is true.
	LogarithmicDepthBuffer *bool   `json:"logatihmicDepthBuffer,omitempty"` // default is false.
}

// WebGLRendererFromJSObject returns a wrapped WebGLRenderer JavaScript class.
func WebGLRendererFromJSObject(p *js.Object) *WebGLRenderer {
	return &WebGLRenderer{p: p}
}

// NewWebGLRenderer returns a new WebGLRenderer object.
func (t *Three) NewWebGLRenderer(opts *WebGLRendererOpts) *WebGLRenderer {
	params := map[string]interface{}{}
	if opts != nil {
		if opts.Precision != nil {
			params["precision"] = *opts.Precision
		}
		if opts.Alpha != nil {
			params["alpha"] = *opts.Alpha
		}
		if opts.PremultipliedAlpha != nil {
			params["premultipliedAlpha"] = *opts.PremultipliedAlpha
		}
		if opts.Antialias != nil {
			params["antialias"] = *opts.Antialias
		}
		if opts.Stencil != nil {
			params["stencil"] = *opts.Stencil
		}
		if opts.PreserveDrawingBuffer != nil {
			params["preserveDrawingBuffer"] = *opts.PreserveDrawingBuffer
		}
		if opts.Depth != nil {
			params["depth"] = *opts.Depth
		}
		if opts.LogarithmicDepthBuffer != nil {
			params["logarithmicDepthBuffer"] = *opts.LogarithmicDepthBuffer
		}
	}
	p := t.ctx.Get("WebGLRenderer").New(params)
	return WebGLRendererFromJSObject(p)
}

// DOMElement returns the DOM element associated with this renderer.
func (w *WebGLRenderer) DOMElement() *js.Object {
	return w.p.Get("domElement")
}

// GetContext returns the WebGL context.
func (w *WebGLRenderer) GetContext() *js.Object {
	return w.p.Call("getContext")
}

// GetContextAttributes returns an object that describes the attributes set on the WebGL context when it was created.
func (w *WebGLRenderer) GetContextAttributes() *js.Object {
	return w.p.Call("getContextAttributes")
}

// ForceContextLoss TODO description.
func (w *WebGLRenderer) ForceContextLoss() *WebGLRenderer {
	w.p.Call("forceContextLoss")
	return w
}

// GetMaxAnisotropy returns the anisotropy level of the textures.
func (w *WebGLRenderer) GetMaxAnisotropy() float64 {
	return w.p.Call("getMaxAnisotropy").Float()
}

// GetPrecision gets the precision used by the shaders. It returns "highp","mediump" or "lowp".
func (w *WebGLRenderer) GetPrecision() float64 {
	return w.p.Call("getPrecision").Float()
}

// GetPixelRatio returns current device pixel ratio used.
func (w *WebGLRenderer) GetPixelRatio() float64 {
	return w.p.Call("getPixelRatio").Float()
}

// SetPixelRatio sets device pixel ratio. This is usually used for HiDPI device to prevent blurring output canvas.
func (w *WebGLRenderer) SetPixelRatio(value float64) *WebGLRenderer {
	w.p.Call("setPixelRatio", value)
	return w
}

// GetSize returns an object containing the width and height of the renderer's output canvas, in pixels.
func (w *WebGLRenderer) GetSize() (width, height float64) {
	result := w.p.Call("getSize")
	return result.Get("width").Float(), result.Get("height").Float()
}

// SetSize resizes the output canvas to (width, height) with device pixel ratio taken into account, and
// also sets the viewport to fit that size, starting in (0, 0). Setting updateStyle to true adds explicit
// pixel units to the output canvas style.
func (w *WebGLRenderer) SetSize(width, height float64, updateStyle bool) *WebGLRenderer {
	w.p.Call("setSize", width, height, updateStyle)
	return w
}

// SetViewport sets the viewport to render from (x, y) to (x + width, y + height).
func (w *WebGLRenderer) SetViewport(x, y, width, height float64) *WebGLRenderer {
	w.p.Call("setViewport", x, y, width, height)
	return w
}

// SetScissor Sets the scissor area from (x, y) to (x + width, y + height).
// NOTE: The point (x, y) is the lower left corner of the area to be set for both of these methods.
// The area is defined from left to right in width but bottom to top in height. The sense of the
// vertical definition is opposite to the fill direction of an HTML canvas element.
func (w *WebGLRenderer) SetScissor(x, y, width, height float64) *WebGLRenderer {
	w.p.Call("setScissor", x, y, width, height)
	return w
}

// SetScissorTest Enable or disable the scissor test. When this is enabled, only the pixels within the
// defined scissor area will be affected by further renderer actions.
func (w *WebGLRenderer) SetScissorTest(value bool) *WebGLRenderer {
	w.p.Call("setScissorTest", value)
	return w
}

// GetClearColor returns a THREE.Color instance with the current clear color.
func (w *WebGLRenderer) GetClearColor() *Color {
	p := w.p.Call("getClearColor")
	return &Color{p: p}
}

// SetClearColor sets the clear color and opacity.
func (w *WebGLRenderer) SetClearColor(color *Color, alpha float64) *WebGLRenderer {
	w.p.Call("setClearColor", color.p, alpha)
	return w
}

// GetClearAlpha returns a float with the current clear alpha. Ranges from 0 to 1.
func (w *WebGLRenderer) GetClearAlpha() float64 {
	return w.p.Call("getClearAlpha").Float()
}

// SetClearAlpha TODO description.
func (w *WebGLRenderer) SetClearAlpha(alpha float64) *WebGLRenderer {
	w.p.Call("setClearAlpha", alpha)
	return w
}

// Clear Tells the renderer to clear its color, depth or stencil drawing buffer(s). This method initializes
// the color buffer to the current clear color value. Arguments default to true.
func (w *WebGLRenderer) Clear(color, depth, stencil bool) *WebGLRenderer {
	w.p.Call("clear", color, depth, stencil)
	return w
}

// ClearColor TODO description.
func (w *WebGLRenderer) ClearColor() *WebGLRenderer {
	w.p.Call("clearColor")
	return w
}

// ClearDepth TODO description.
func (w *WebGLRenderer) ClearDepth() *WebGLRenderer {
	w.p.Call("clearDepth")
	return w
}

// ClearStencil TODO description.
func (w *WebGLRenderer) ClearStencil() *WebGLRenderer {
	w.p.Call("clearStencil")
	return w
}

// ClearTarget clears a rendertarget. To do this, it activates the rendertarget.
//
//     renderTarget -- The renderTarget that needs to be cleared.
//     color -- If set, then the color gets cleared.
//     depth -- If set, then the depth gets cleared.
//     stencil -- If set, then the stencil gets cleared.
//     This method clears a rendertarget. To do this, it activates the rendertarget
func (w *WebGLRenderer) ClearTarget(renderTarget JSObject, color, depth, stencil bool) *WebGLRenderer {
	w.p.Call("clearStencil", renderTarget.JSObject(), color, depth, stencil)
	return w
}

// ResetGLState TODO description.
func (w *WebGLRenderer) ResetGLState() *WebGLRenderer {
	w.p.Call("resetGLState")
	return w
}

// RenderBufferImmediate renders an immediate buffer. Gets called by renderImmediateObject.
//
//     object — an instance of Object3D
//     program — an instance of shaderProgram
//     shading — an instance of Material
func (w *WebGLRenderer) RenderBufferImmediate(object, program, material JSObject) *WebGLRenderer {
	w.p.Call("renderBufferImmediate", object.JSObject(), program.JSObject(), material.JSObject())
	return w
}

// RenderBufferDirect renders a buffer geometry group using the camera and with the correct material.
func (w *WebGLRenderer) RenderBufferDirect(camera, fog, geometry, material, object, group JSObject) *WebGLRenderer {
	w.p.Call("renderBufferDirect", camera.JSObject(), fog.JSObject(), geometry.JSObject(), material.JSObject(), object.JSObject(), group.JSObject())
	return w
}

// RenderOpts represents the optional arguments that can be passed to Render.
type RenderOpts struct {
	renderTarget JSObject
	forceClear   bool
}

// Render Render a scene using a camera. The render is done to the renderTarget (if specified) or to the canvas
// as usual. If forceClear is true, the depth, stencil and color buffers will be cleared before rendering even
// if the renderer's autoClear property is false. Even with forceClear set to true you can prevent certain
// buffers being cleared by setting either the .autoClearColor, .autoClearStencil or .autoClearDepth properties to false.
func (w *WebGLRenderer) Render(scene, camera JSObject, opts *RenderOpts) *WebGLRenderer {
	if opts != nil {
		w.p.Call("render", scene.JSObject(), camera.JSObject(), opts.renderTarget.JSObject(), opts.forceClear)
	} else {
		w.p.Call("render", scene.JSObject(), camera.JSObject())
	}
	return w
}

// SetFaceCulling is used for setting the gl frontFace, cullFace states in the GPU, thus enabling/disabling
// face culling when rendering. If cullFace is false, culling will be disabled.
//
//     cullFace —- "back", "front", "front_and_back", or false.
//     frontFace —- "ccw" or "cw
func (w *WebGLRenderer) SetFaceCulling(cullFace, frontFaceDirection int) *WebGLRenderer {
	w.p.Call("setFaceCulling", cullFace, frontFaceDirection)
	return w
}

// SetTexture sets the correct texture to the correct slot for the wegl shader. The slot number can be found
// as a value of the uniform of the sampler.
//
//     texture -- The texture that needs to be set.
//     slot -- The number indicating which slot should be used by the texture.
func (w *WebGLRenderer) SetTexture(texture, slot JSObject) *WebGLRenderer {
	w.p.Call("setTexture", texture.JSObject(), slot.JSObject())
	return w
}

// GetCurrentRenderTarget TODO description.
func (w *WebGLRenderer) GetCurrentRenderTarget() *js.Object {
	return w.p.Call("getCurrentRenderTarget")
}

// SetRenderTarget sets the active rendertarget. If the parameter is omitted the canvas is set as the active rendertarget.
//
//     renderTarget -- The renderTarget that needs to be activated (optional).
func (w *WebGLRenderer) SetRenderTarget(renderTarget JSObject) *WebGLRenderer {
	w.p.Call("setRenderTarget", renderTarget.JSObject())
	return w
}

// ReadRenderTargetPixels reads the pixel data from the renderTarget into the buffer you pass in. Buffer should be a
// Javascript Uint8Array instantiated with new Uint8Array( renderTargetWidth * renderTargetWidth * 4 ) to account for
// size and color information. This is a wrapper around gl.readPixels.
func (w *WebGLRenderer) ReadRenderTargetPixels(renderTarget JSObject, x, y, width, height int, buffer JSObject) *WebGLRenderer {
	w.p.Call("readRenderTargetPixels", renderTarget.JSObject(), x, y, width, height, buffer.JSObject())
	return w
}

// TODO: ForceContextLoss, SetClearAlpha, ClearColor, ClearDepth, ClearStencil, ResetGLState, GetCurrentRenderTarget
//
// SetClearColor example:
// // Creates a renderer with red background
//    var renderer = new THREE.WebGLRenderer();
//    renderer.setSize( 200, 100 );
//    renderer.setClearColor( 0xff0000 );
