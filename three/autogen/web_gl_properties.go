package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLProperties represents a webglproperties.
type WebGLProperties struct{ p *js.Object }

// WebGLProperties returns a webglproperties object.
func (t *Three) WebGLProperties() *WebGLProperties {
	p := t.ctx.Get("WebGLProperties")
	return &WebGLProperties{p: p}
}

// NewWebGLProperties returns a new webglproperties object.
func (t *WebGLProperties) New() *WebGLProperties {
	p := t.p.New()
	return &WebGLProperties{p: p}
}

