package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// MeshDepthMaterial represents a meshdepthmaterial.
type MeshDepthMaterial struct{ p *js.Object }

// MeshDepthMaterial returns a MeshDepthMaterial object.
func (t *Three) MeshDepthMaterial() *MeshDepthMaterial {
	p := t.ctx.Get("MeshDepthMaterial")
	return &MeshDepthMaterial{p: p}
}

// New returns a new MeshDepthMaterial object.
func (t *MeshDepthMaterial) New(parameters float64) *MeshDepthMaterial {
	p := t.p.New(parameters)
	return &MeshDepthMaterial{p: p}
}

// Copy TODO description.
func (m *MeshDepthMaterial) Copy(source float64) *MeshDepthMaterial {
	m.p.Call("copy", source)
	return m
}

