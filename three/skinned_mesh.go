// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// SkinnedMesh represents a mesh that has a Skeleton with bones that can then be used
// to animate the vertices of the geometry.
//
// http://threejs.org/docs/index.html#Reference/Objects/SkinnedMesh
type SkinnedMesh struct{ *Mesh }

// JSObject returns the underlying *js.Object.
func (s *SkinnedMesh) JSObject() *js.Object { return s.p }

// SkinnedMesh returns a SkinnedMesh JavaScript class.
func (t *Three) SkinnedMesh() *SkinnedMesh {
	return SkinnedMeshFromJSObject(t.ctx.Get("SkinnedMesh"))
}

// SkinnedMeshFromJSObject returns a wrapped SkinnedMesh JavaScript class.
func SkinnedMeshFromJSObject(p *js.Object) *SkinnedMesh {
	return &SkinnedMesh{&Mesh{&Object3D{p: p}}}
}

// NewSkinnedMesh returns a new SkinnedMesh object.
//
//     geometry — An instance of Geometry. Geometry.skinIndices and Geometry.skinWeights should be set.
//     material — An instance of Material (optional).
//     useVertexTexture -- Defines whether a vertex texture can be used (optional).
func (t *Three) NewSkinnedMesh(geometry, material *js.Object, useVertexTexture bool) *SkinnedMesh {
	return SkinnedMeshFromJSObject(t.ctx.Get("SkinnedMesh").New(geometry, material, useVertexTexture))
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
