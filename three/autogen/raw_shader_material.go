package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// RawShaderMaterial represents a rawshadermaterial.
type RawShaderMaterial struct{ p *js.Object }

// RawShaderMaterial returns a rawshadermaterial object.
func (t *Three) RawShaderMaterial() *RawShaderMaterial {
	p := t.ctx.Get("RawShaderMaterial")
	return &RawShaderMaterial{p: p}
}

// NewRawShaderMaterial returns a new rawshadermaterial object.
func (t *RawShaderMaterial) New(parameters float64) *RawShaderMaterial {
	p := t.p.New(parameters)
	return &RawShaderMaterial{p: p}
}

