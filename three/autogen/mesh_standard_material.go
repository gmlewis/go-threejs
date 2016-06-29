package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// MeshStandardMaterial represents a meshstandardmaterial.
type MeshStandardMaterial struct{ p *js.Object }

// MeshStandardMaterial returns a MeshStandardMaterial object.
func (t *Three) MeshStandardMaterial() *MeshStandardMaterial {
	p := t.ctx.Get("MeshStandardMaterial")
	return &MeshStandardMaterial{p: p}
}

// New returns a new MeshStandardMaterial object.
func (t *MeshStandardMaterial) New(parameters float64) *MeshStandardMaterial {
	p := t.p.New(parameters)
	return &MeshStandardMaterial{p: p}
}

// Copy TODO description.
func (m *MeshStandardMaterial) Copy(source float64) *MeshStandardMaterial {
	m.p.Call("copy", source)
	return m
}
