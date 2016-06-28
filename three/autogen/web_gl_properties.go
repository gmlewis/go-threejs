package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLProperties represents a webglproperties.
type WebGLProperties struct{ p *js.Object }

// WebGLProperties returns a WebGLProperties object.
func (t *Three) WebGLProperties() *WebGLProperties {
	p := t.ctx.Get("WebGLProperties")
	return &WebGLProperties{p: p}
}

// New returns a new WebGLProperties object.
func (t *WebGLProperties) New() *WebGLProperties {
	p := t.p.New()
	return &WebGLProperties{p: p}
}

