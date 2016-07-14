// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// InstancedInterleavedBuffer represents an instancedinterleavedbuffer.
type InstancedInterleavedBuffer struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (i *InstancedInterleavedBuffer) JSObject() *js.Object { return i.p }

// InstancedInterleavedBuffer returns an InstancedInterleavedBuffer JavaScript class.
func (t *Three) InstancedInterleavedBuffer() *InstancedInterleavedBuffer {
	p := t.ctx.Get("InstancedInterleavedBuffer")
	return InstancedInterleavedBufferFromJSObject(p)
}

// InstancedInterleavedBufferFromJSObject returns a wrapped InstancedInterleavedBuffer JavaScript class.
func InstancedInterleavedBufferFromJSObject(p *js.Object) *InstancedInterleavedBuffer {
	return &InstancedInterleavedBuffer{p: p}
}

// NewInstancedInterleavedBuffer returns a new InstancedInterleavedBuffer object.
func (t *Three) NewInstancedInterleavedBuffer(array, stride, meshPerAttribute float64) *InstancedInterleavedBuffer {
	p := t.ctx.Get("InstancedInterleavedBuffer").New(array, stride, meshPerAttribute)
	return InstancedInterleavedBufferFromJSObject(p)
}

// Copy TODO description.
func (i *InstancedInterleavedBuffer) Copy(source *InstancedInterleavedBuffer) *InstancedInterleavedBuffer {
	i.p.Call("copy", source.p)
	return i
}
