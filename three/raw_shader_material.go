package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// RawShaderMaterial represents a rawshadermaterial.
type RawShaderMaterial struct{ p *js.Object }

// RawShaderMaterial returns a RawShaderMaterial object.
func (t *Three) RawShaderMaterial() *RawShaderMaterial {
	p := t.ctx.Get("RawShaderMaterial")
	return &RawShaderMaterial{p: p}
}

// New returns a new RawShaderMaterial object.
func (t *RawShaderMaterial) New(parameters float64) *RawShaderMaterial {
	p := t.p.New(parameters)
	return &RawShaderMaterial{p: p}
}

