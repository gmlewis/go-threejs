package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Mesh represents a mesh.
type Mesh struct{ p *js.Object }

// Mesh returns a mesh object.
func (t *Three) Mesh() *Mesh {
	p := t.ctx.Get("Mesh")
	return &Mesh{p: p}
}

// NewMesh returns a new mesh object.
func (t *Mesh) New(geometry, material float64) *Mesh {
	p := t.p.New(geometry, material)
	return &Mesh{p: p}
}

// SetDrawMode TODO description.
func (m *Mesh) SetDrawMode(value float64) *Mesh {
	m.p.Call("setDrawMode", value)
	return m
}

// GetMorphTargetIndexByName TODO description.
func (m *Mesh) GetMorphTargetIndexByName(name float64) *Mesh {
	m.p.Call("getMorphTargetIndexByName", name)
	return m
}

