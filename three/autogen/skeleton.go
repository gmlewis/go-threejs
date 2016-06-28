package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Skeleton represents a skeleton.
type Skeleton struct{ p *js.Object }

// Skeleton returns a Skeleton object.
func (t *Three) Skeleton() *Skeleton {
	p := t.ctx.Get("Skeleton")
	return &Skeleton{p: p}
}

// New returns a new Skeleton object.
func (t *Skeleton) New(bones, boneInverses, useVertexTexture float64) *Skeleton {
	p := t.p.New(bones, boneInverses, useVertexTexture)
	return &Skeleton{p: p}
}

