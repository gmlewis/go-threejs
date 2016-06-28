package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// SkinnedMesh represents a skinnedmesh.
type SkinnedMesh struct{ p *js.Object }

// SkinnedMesh returns a SkinnedMesh object.
func (t *Three) SkinnedMesh() *SkinnedMesh {
	p := t.ctx.Get("SkinnedMesh")
	return &SkinnedMesh{p: p}
}

// New returns a new SkinnedMesh object.
func (t *SkinnedMesh) New(geometry, material, useVertexTexture float64) *SkinnedMesh {
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

