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
func (t *InstancedInterleavedBuffer) JSObject() *js.Object { return t.p }

// InstancedInterleavedBuffer returns an InstancedInterleavedBuffer object.
func (t *Three) InstancedInterleavedBuffer() *InstancedInterleavedBuffer {
	p := t.ctx.Get("InstancedInterleavedBuffer")
	return &InstancedInterleavedBuffer{p: p}
}

// New returns a new InstancedInterleavedBuffer object.
func (t *InstancedInterleavedBuffer) New(array, stride, meshPerAttribute float64) *InstancedInterleavedBuffer {
	p := t.p.New(array, stride, meshPerAttribute)
	return &InstancedInterleavedBuffer{p: p}
}

// Copy TODO description.
func (i *InstancedInterleavedBuffer) Copy(source float64) *InstancedInterleavedBuffer {
	i.p.Call("copy", source)
	return i
}
