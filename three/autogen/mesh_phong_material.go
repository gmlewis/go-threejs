package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// MeshPhongMaterial represents a meshphongmaterial.
type MeshPhongMaterial struct{ p *js.Object }

// MeshPhongMaterial returns a meshphongmaterial object.
func (t *Three) MeshPhongMaterial() *MeshPhongMaterial {
	p := t.ctx.Get("MeshPhongMaterial")
	return &MeshPhongMaterial{p: p}
}

// NewMeshPhongMaterial returns a new meshphongmaterial object.
func (t *MeshPhongMaterial) New(parameters float64) *MeshPhongMaterial {
	p := t.p.New(parameters)
	return &MeshPhongMaterial{p: p}
}

// Copy TODO description.
func (m *MeshPhongMaterial) Copy(source float64) *MeshPhongMaterial {
	m.p.Call("copy", source)
	return m
}

