package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Mesh is the base class for mesh objects.
//
// http://threejs.org/docs/index.html#Reference/Objects/Mesh
type Mesh struct{ p *js.Object }

// Mesh returns a Mesh object.
func (t *Three) Mesh() *Mesh {
	p := t.ctx.Get("Mesh")
	return &Mesh{p: p}
}

// New returns a new Mesh object.
//
//     geometry — an instance of geometry.
//     material — an instance of material (optional).
func (t *Mesh) New(geometry, material *js.Object) *Mesh {
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
