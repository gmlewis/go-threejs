package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// MeshLambertMaterial represents a meshlambertmaterial.
type MeshLambertMaterial struct{ p *js.Object }

// MeshLambertMaterial returns a MeshLambertMaterial object.
func (t *Three) MeshLambertMaterial() *MeshLambertMaterial {
	p := t.ctx.Get("MeshLambertMaterial")
	return &MeshLambertMaterial{p: p}
}

// New returns a new MeshLambertMaterial object.
func (t *MeshLambertMaterial) New(parameters float64) *MeshLambertMaterial {
	p := t.p.New(parameters)
	return &MeshLambertMaterial{p: p}
}

// Copy TODO description.
func (m *MeshLambertMaterial) Copy(source float64) *MeshLambertMaterial {
	m.p.Call("copy", source)
	return m
}

