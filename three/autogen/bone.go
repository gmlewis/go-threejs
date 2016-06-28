package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Bone represents a bone.
type Bone struct{ p *js.Object }

// Bone returns a bone object.
func (t *Three) Bone() *Bone {
	p := t.ctx.Get("Bone")
	return &Bone{p: p}
}

// NewBone returns a new bone object.
func (t *Bone) New(skin float64) *Bone {
	p := t.p.New(skin)
	return &Bone{p: p}
}

// Copy TODO description.
func (b *Bone) Copy(source float64) *Bone {
	b.p.Call("copy", source)
	return b
}

