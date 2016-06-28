package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Skeleton represents a skeleton.
type Skeleton struct{ p *js.Object }

// Skeleton returns a skeleton object.
func (t *Three) Skeleton() *Skeleton {
	p := t.ctx.Get("Skeleton")
	return &Skeleton{p: p}
}

// NewSkeleton returns a new skeleton object.
func (t *Skeleton) New(bones, boneInverses, useVertexTexture float64) *Skeleton {
	p := t.p.New(bones, boneInverses, useVertexTexture)
	return &Skeleton{p: p}
}

