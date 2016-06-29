package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLRenderer represents a webglrenderer.
type WebGLRenderer struct{ p *js.Object }

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
