package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// SkinnedMesh represents a mesh that has a Skeleton with bones that can then be used
// to animate the vertices of the geometry.
//
// http://threejs.org/docs/index.html#Reference/Objects/SkinnedMesh
type SkinnedMesh struct{ p *js.Object }

// SkinnedMesh returns a SkinnedMesh object.
func (t *Three) SkinnedMesh() *SkinnedMesh {
	p := t.ctx.Get("SkinnedMesh")
	return &SkinnedMesh{p: p}
}

// New returns a new SkinnedMesh object.
//
//     geometry — An instance of Geometry. Geometry.skinIndices and Geometry.skinWeights should be set.
//     material — An instance of Material (optional).
//     useVertexTexture -- Defines whether a vertex texture can be used (optional).
func (t *SkinnedMesh) New(geometry, material *js.Object, useVertexTexture bool) *SkinnedMesh {
	p := t.p.New(geometry, material, useVertexTexture)
	return &SkinnedMesh{p: p}
}

// Bind TODO description.
func (s *SkinnedMesh) Bind(skeleton, bindMatrix float64) *SkinnedMesh {
	s.p.Call("bind", skeleton, bindMatrix)
	return s
}

// UpdateMatrixWorld TODO description.
func (s *SkinnedMesh) UpdateMatrixWorld(force float64) *SkinnedMesh {
	s.p.Call("updateMatrixWorld", force)
	return s
}
