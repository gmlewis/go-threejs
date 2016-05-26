package three

import "github.com/gopherjs/gopherjs/js"

// shaders/...

// webgl/...

// WebGLRendererOpts provides optional parameters to WebGLRenderer.
type WebGLRendererOpts struct {
	// canvas — A Canvas where the renderer draws its output.
	// context — The RenderingContext context to use.
	Precision              *string // Shader precision. Can be "highp", "mediump" or "lowp". Defaults to "highp" if supported by the device.
	Alpha                  bool    // default is false.
	PremultipliedAlpha     *bool   // default is true.
	Antialias              bool    // default is false.
	Stencil                *bool   // default is true.
	PreserveDrawingBuffer  bool    // default is false.
	Depth                  *bool   // default is true.
	LogarithmicDepthBuffer bool    // default is false.
}

// WebGLRenderer.js
// http://threejs.org/docs/index.html#Reference/Renderers/WebGLRenderer
func (t *Three) WebGLRenderer(opts *WebGLRendererOpts) *js.Object {
	params := map[string]interface{}{}
	if opts != nil {
		if opts.Precision != nil {
			params["precision"] = *opts.Precision
		}
		params["alpha"] = opts.Alpha
		if opts.PremultipliedAlpha != nil {
			params["premultipliedAlpha"] = *opts.PremultipliedAlpha
		}
		params["antialias"] = opts.Antialias
		if opts.Stencil != nil {
			params["stencil"] = *opts.Stencil
		}
		params["preserveDrawingBuffer"] = opts.PreserveDrawingBuffer
		if opts.Depth != nil {
			params["depth"] = *opts.Depth
		}
		params["logarithmicDepthBuffer"] = opts.LogarithmicDepthBuffer
	}
	return t.ctx.Get("WebGLRenderer").New(params)
}

// WebGLRenderTargetCube.js

// WebGLRenderTarget.js
