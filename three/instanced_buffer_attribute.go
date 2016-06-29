package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// InstancedBufferAttribute represents an instancedbufferattribute.
type InstancedBufferAttribute struct{ p *js.Object }

// InstancedBufferAttribute returns an InstancedBufferAttribute object.
func (t *Three) InstancedBufferAttribute() *InstancedBufferAttribute {
	p := t.ctx.Get("InstancedBufferAttribute")
	return &InstancedBufferAttribute{p: p}
}

// New returns a new InstancedBufferAttribute object.
func (t *InstancedBufferAttribute) New(array, itemSize, meshPerAttribute float64) *InstancedBufferAttribute {
	p := t.p.New(array, itemSize, meshPerAttribute)
	return &InstancedBufferAttribute{p: p}
}

// Copy TODO description.
func (i *InstancedBufferAttribute) Copy(source float64) *InstancedBufferAttribute {
	i.p.Call("copy", source)
	return i
}

