package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// MeshNormalMaterial represents a meshnormalmaterial.
type MeshNormalMaterial struct{ p *js.Object }

// MeshNormalMaterial returns a meshnormalmaterial object.
func (t *Three) MeshNormalMaterial() *MeshNormalMaterial {
	p := t.ctx.Get("MeshNormalMaterial")
	return &MeshNormalMaterial{p: p}
}

// NewMeshNormalMaterial returns a new meshnormalmaterial object.
func (t *MeshNormalMaterial) New(parameters float64) *MeshNormalMaterial {
	p := t.p.New(parameters)
	return &MeshNormalMaterial{p: p}
}

// Copy TODO description.
func (m *MeshNormalMaterial) Copy(source float64) *MeshNormalMaterial {
	m.p.Call("copy", source)
	return m
}

