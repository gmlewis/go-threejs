package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// InstancedInterleavedBuffer represents an instancedinterleavedbuffer.
type InstancedInterleavedBuffer struct{ p *js.Object }

// InstancedInterleavedBuffer returns an instancedinterleavedbuffer object.
func (t *Three) InstancedInterleavedBuffer() *InstancedInterleavedBuffer {
	p := t.ctx.Get("InstancedInterleavedBuffer")
	return &InstancedInterleavedBuffer{p: p}
}

// NewInstancedInterleavedBuffer returns a new instancedinterleavedbuffer object.
func (t *InstancedInterleavedBuffer) New(array, stride, meshPerAttribute float64) *InstancedInterleavedBuffer {
	p := t.p.New(array, stride, meshPerAttribute)
	return &InstancedInterleavedBuffer{p: p}
}

// Copy TODO description.
func (i *InstancedInterleavedBuffer) Copy(source float64) *InstancedInterleavedBuffer {
	i.p.Call("copy", source)
	return i
}

