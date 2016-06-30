// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLRenderer represents a webglrenderer.
type WebGLRenderer struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *WebGLRenderer) JSObject() *js.Object { return t.p }

// WebGLRenderer returns a WebGLRenderer object.
func (t *Three) WebGLRenderer() *WebGLRenderer {
	p := t.ctx.Get("WebGLRenderer")
	return &WebGLRenderer{p: p}
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

// New returns a new WebGLRenderer object.
func (t *WebGLRenderer) New(opts *WebGLRendererOpts) *WebGLRenderer {
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
	p := t.p.New(params)
	return &WebGLRenderer{p: p}
}

// DOMElement returns the DOM element associated with this renderer.
func (w *WebGLRenderer) DOMElement() *js.Object {
	return w.p.Get("domElement")
}

// GetContext TODO description.
func (w *WebGLRenderer) GetContext() *js.Object {
	return w.p.Call("getContext")
}

// GetContextAttributes TODO description.
func (w *WebGLRenderer) GetContextAttributes() *js.Object {
	return w.p.Call("getContextAttributes")
}

// ForceContextLoss TODO description.
func (w *WebGLRenderer) ForceContextLoss() *WebGLRenderer {
	w.p.Call("forceContextLoss")
	return w
}

// GetMaxAnisotropy TODO description.
func (w *WebGLRenderer) GetMaxAnisotropy() float64 {
	return w.p.Call("getMaxAnisotropy").Float()
}

// GetPrecision TODO description.
func (w *WebGLRenderer) GetPrecision() float64 {
	return w.p.Call("getPrecision").Float()
}

// GetPixelRatio TODO description.
func (w *WebGLRenderer) GetPixelRatio() float64 {
	return w.p.Call("getPixelRatio").Float()
}

// SetPixelRatio TODO description.
func (w *WebGLRenderer) SetPixelRatio(value float64) *WebGLRenderer {
	w.p.Call("setPixelRatio", value)
	return w
}

// GetSize TODO description.
func (w *WebGLRenderer) GetSize() (width, height float64) {
	result := w.p.Call("getSize")
	return result.Get("width").Float(), result.Get("height").Float()
}

// SetSize TODO description.
func (w *WebGLRenderer) SetSize(width, height float64, updateStyle bool) *WebGLRenderer {
	w.p.Call("setSize", width, height, updateStyle)
	return w
}

// SetViewport TODO description.
func (w *WebGLRenderer) SetViewport(x, y, width, height float64) *WebGLRenderer {
	w.p.Call("setViewport", x, y, width, height)
	return w
}

// SetScissor TODO description.
func (w *WebGLRenderer) SetScissor(x, y, width, height float64) *WebGLRenderer {
	w.p.Call("setScissor", x, y, width, height)
	return w
}

// SetScissorTest TODO description.
func (w *WebGLRenderer) SetScissorTest(value bool) *WebGLRenderer {
	w.p.Call("setScissorTest", value)
	return w
}

// GetClearColor TODO description.
func (w *WebGLRenderer) GetClearColor() *Color {
	p := w.p.Call("getClearColor")
	return &Color{p: p}
}

// SetClearColor TODO description.
func (w *WebGLRenderer) SetClearColor(color *Color, alpha float64) *WebGLRenderer {
	w.p.Call("setClearColor", color.p, alpha)
	return w
}

// GetClearAlpha TODO description.
func (w *WebGLRenderer) GetClearAlpha() float64 {
	return w.p.Call("getClearAlpha").Float()
}

// SetClearAlpha TODO description.
func (w *WebGLRenderer) SetClearAlpha(alpha float64) *WebGLRenderer {
	w.p.Call("setClearAlpha", alpha)
	return w
}

// Clear TODO description.
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

// ClearTarget TODO description.
func (w *WebGLRenderer) ClearTarget(renderTarget JSObject, color, depth, stencil bool) *WebGLRenderer {
	w.p.Call("clearStencil", renderTarget.JSObject(), color, depth, stencil)
	return w
}

// ResetGLState TODO description.
func (w *WebGLRenderer) ResetGLState() *WebGLRenderer {
	w.p.Call("resetGLState")
	return w
}

// RenderBufferImmediate TODO description.
func (w *WebGLRenderer) RenderBufferImmediate(object, program, material JSObject) *WebGLRenderer {
	w.p.Call("renderBufferImmediate", object.JSObject(), program.JSObject(), material.JSObject())
	return w
}

// RenderBufferDirect TODO description.
func (w *WebGLRenderer) RenderBufferDirect(camera, fog, geometry, material, object, group JSObject) *WebGLRenderer {
	w.p.Call("renderBufferDirect", camera.JSObject(), fog.JSObject(), geometry.JSObject(), material.JSObject(), object.JSObject(), group.JSObject())
	return w
}

// RenderOpts represents the optional arguments that can be passed to Render.
type RenderOpts struct {
	renderTarget JSObject
	forceClear   bool
}

// Render TODO description.
func (w *WebGLRenderer) Render(scene, camera JSObject, opts *RenderOpts) *WebGLRenderer {
	if opts != nil {
		w.p.Call("render", scene.JSObject(), camera.JSObject(), opts.renderTarget.JSObject(), opts.forceClear)
	} else {
		w.p.Call("render", scene.JSObject(), camera.JSObject())
	}
	return w
}

// SetFaceCulling TODO description.
func (w *WebGLRenderer) SetFaceCulling(cullFace, frontFaceDirection int) *WebGLRenderer {
	w.p.Call("setFaceCulling", cullFace, frontFaceDirection)
	return w
}

// SetTexture TODO description.
func (w *WebGLRenderer) SetTexture(texture, slot JSObject) *WebGLRenderer {
	w.p.Call("setTexture", texture.JSObject(), slot.JSObject())
	return w
}

// GetCurrentRenderTarget TODO description.
func (w *WebGLRenderer) GetCurrentRenderTarget() *js.Object {
	return w.p.Call("getCurrentRenderTarget")
}

// SetRenderTarget TODO description.
func (w *WebGLRenderer) SetRenderTarget(renderTarget JSObject) *WebGLRenderer {
	w.p.Call("setRenderTarget", renderTarget.JSObject())
	return w
}

// ReadRenderTargetPixels TODO description.
func (w *WebGLRenderer) ReadRenderTargetPixels(renderTarget JSObject, x, y, width, height int, buffer JSObject) *WebGLRenderer {
	w.p.Call("readRenderTargetPixels", renderTarget.JSObject(), x, y, width, height, buffer.JSObject())
	return w
}
