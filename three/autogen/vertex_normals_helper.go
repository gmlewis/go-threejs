package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// VertexNormalsHelper represents a vertexnormalshelper.
type VertexNormalsHelper struct{ p *js.Object }

// VertexNormalsHelper returns a VertexNormalsHelper object.
func (t *Three) VertexNormalsHelper() *VertexNormalsHelper {
	p := t.ctx.Get("VertexNormalsHelper")
	return &VertexNormalsHelper{p: p}
}

// New returns a new VertexNormalsHelper object.
func (t *VertexNormalsHelper) New(object, size, hex, linewidth float64) *VertexNormalsHelper {
	p := t.p.New(object, size, hex, linewidth)
	return &VertexNormalsHelper{p: p}
}

