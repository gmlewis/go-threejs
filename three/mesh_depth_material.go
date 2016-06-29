package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// MeshDepthMaterial represents a material for drawing geometry by depth.
// Depth is based off of the camera near and far plane. White is nearest, black is farthest.
//
// http://threejs.org/docs/index.html#Reference/Materials/MeshDepthMaterial
type MeshDepthMaterial struct{ p *js.Object }

// MeshDepthMaterial returns a MeshDepthMaterial object.
func (t *Three) MeshDepthMaterial() *MeshDepthMaterial {
	p := t.ctx.Get("MeshDepthMaterial")
	return &MeshDepthMaterial{p: p}
}

// New returns a new MeshDepthMaterial object.
//
// parameters is an object with one or more properties defining the material's appearance:
//     morphTargets -- Define whether the material uses morphTargets. Default is false.
//     wireframe -- Render geometry as wireframe. Default is false (i.e. render as smooth shaded).
//     wireframeLinewidth -- Controls wireframe thickness. Default is 1.
func (t *MeshDepthMaterial) New(parameters map[string]interface{}) *MeshDepthMaterial {
	p := t.p.New(parameters)
	return &MeshDepthMaterial{p: p}
}

// Copy TODO description.
func (m *MeshDepthMaterial) Copy(source *MeshDepthMaterial) *MeshDepthMaterial {
	m.p.Call("copy", source.p)
	return m
}
