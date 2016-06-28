package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// MeshBasicMaterial represents a meshbasicmaterial.
type MeshBasicMaterial struct{ p *js.Object }

// MeshBasicMaterial returns a MeshBasicMaterial object.
func (t *Three) MeshBasicMaterial() *MeshBasicMaterial {
	p := t.ctx.Get("MeshBasicMaterial")
	return &MeshBasicMaterial{p: p}
}

// New returns a new MeshBasicMaterial object.
func (t *MeshBasicMaterial) New(parameters float64) *MeshBasicMaterial {
	p := t.p.New(parameters)
	return &MeshBasicMaterial{p: p}
}

// Copy TODO description.
func (m *MeshBasicMaterial) Copy(source float64) *MeshBasicMaterial {
	m.p.Call("copy", source)
	return m
}

