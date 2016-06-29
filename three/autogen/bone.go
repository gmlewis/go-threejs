package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Bone represents a bone which is part of a skeleton.
//
// http://threejs.org/docs/index.html#Reference/Objects/Bone
type Bone struct{ p *js.Object }

// Bone returns a Bone object.
func (t *Three) Bone() *Bone {
	p := t.ctx.Get("Bone")
	return &Bone{p: p}
}

// New returns a new Bone object.
func (t *Bone) New(skin float64) *Bone {
	p := t.p.New(skin)
	return &Bone{p: p}
}

// Copy TODO description.
func (b *Bone) Copy(source float64) *Bone {
	b.p.Call("copy", source)
	return b
}
