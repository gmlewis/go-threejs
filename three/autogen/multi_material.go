package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// MultiMaterial represents a multimaterial.
type MultiMaterial struct{ p *js.Object }

// MultiMaterial returns a multimaterial object.
func (t *Three) MultiMaterial() *MultiMaterial {
	p := t.ctx.Get("MultiMaterial")
	return &MultiMaterial{p: p}
}

// NewMultiMaterial returns a new multimaterial object.
func (t *MultiMaterial) New(materials float64) *MultiMaterial {
	p := t.p.New(materials)
	return &MultiMaterial{p: p}
}

// ToJSON TODO description.
func (m *MultiMaterial) ToJSON(meta float64) *MultiMaterial {
	m.p.Call("toJSON", meta)
	return m
}

// Clone TODO description.
func (m *MultiMaterial) Clone() *MultiMaterial {
	m.p.Call("clone")
	return m
}

