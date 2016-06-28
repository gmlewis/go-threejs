package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLShadowMap represents a webglshadowmap.
type WebGLShadowMap struct{ p *js.Object }

// WebGLShadowMap returns a WebGLShadowMap object.
func (t *Three) WebGLShadowMap() *WebGLShadowMap {
	p := t.ctx.Get("WebGLShadowMap")
	return &WebGLShadowMap{p: p}
}

// New returns a new WebGLShadowMap object.
func (t *WebGLShadowMap) New(_renderer, _lights, _objects float64) *WebGLShadowMap {
	p := t.p.New(_renderer, _lights, _objects)
	return &WebGLShadowMap{p: p}
}

