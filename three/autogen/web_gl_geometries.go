package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WebGLGeometries represents a webglgeometries.
type WebGLGeometries struct{ p *js.Object }

// WebGLGeometries returns a webglgeometries object.
func (t *Three) WebGLGeometries() *WebGLGeometries {
	p := t.ctx.Get("WebGLGeometries")
	return &WebGLGeometries{p: p}
}

// NewWebGLGeometries returns a new webglgeometries object.
func (t *WebGLGeometries) New(gl, properties, info float64) *WebGLGeometries {
	p := t.p.New(gl, properties, info)
	return &WebGLGeometries{p: p}
}

