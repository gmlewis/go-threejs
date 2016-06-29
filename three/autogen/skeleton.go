package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Skeleton uses an array of bones to create a skeleton that can be used by a SkinnedMesh.
//
// http://threejs.org/docs/index.html#Reference/Objects/Skeleton
type Skeleton struct{ p *js.Object }

// Skeleton returns a Skeleton object.
func (t *Three) Skeleton() *Skeleton {
	p := t.ctx.Get("Skeleton")
	return &Skeleton{p: p}
}

// New returns a new Skeleton object.
//
//     bones — The array of bones
//     boneInverses — (optional) An array of Matrix4s
//     useVertexTexture — (optional) Whether or not to use a vertex texture in the shader.
func (t *Skeleton) New(bones, boneInverses []*js.Object, useVertexTexture bool) *Skeleton {
	p := t.p.New(bones, boneInverses, useVertexTexture)
	return &Skeleton{p: p}
}
