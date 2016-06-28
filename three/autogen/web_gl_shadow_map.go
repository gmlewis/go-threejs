package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLShadowMap represents a webglshadowmap.
type WebGLShadowMap struct{ p *js.Object }

// WebGLShadowMap returns a webglshadowmap object.
func (t *Three) WebGLShadowMap() *WebGLShadowMap {
	p := t.ctx.Get("WebGLShadowMap")
	return &WebGLShadowMap{p: p}
}

// NewWebGLShadowMap returns a new webglshadowmap object.
func (t *WebGLShadowMap) New(_renderer, _lights, _objects float64) *WebGLShadowMap {
	p := t.p.New(_renderer, _lights, _objects)
	return &WebGLShadowMap{p: p}
}

