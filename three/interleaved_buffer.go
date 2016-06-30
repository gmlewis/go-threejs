// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// InterleavedBuffer represents an interleavedbuffer.
type InterleavedBuffer struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *InterleavedBuffer) JSObject() *js.Object { return t.p }

// InterleavedBuffer returns an InterleavedBuffer object.
func (t *Three) InterleavedBuffer() *InterleavedBuffer {
	p := t.ctx.Get("InterleavedBuffer")
	return &InterleavedBuffer{p: p}
}

// New returns a new InterleavedBuffer object.
func (t *InterleavedBuffer) New(array, stride float64) *InterleavedBuffer {
	p := t.p.New(array, stride)
	return &InterleavedBuffer{p: p}
}

// Length returns the length-component of the InterleavedBuffer.
func (i *InterleavedBuffer) Length() float64 {
	return i.p.Get("length").Float()
}

// Count returns the count-component of the InterleavedBuffer.
func (i *InterleavedBuffer) Count() float64 {
	return i.p.Get("count").Float()
}

// SetNeedsUpdate sets the needsUpdate-component of the InterleavedBuffer.
func (i *InterleavedBuffer) SetNeedsUpdate(value float64) *InterleavedBuffer {
	i.p.Set("needsUpdate", value)
	return i
}

// SetDynamic TODO description.
func (i *InterleavedBuffer) SetDynamic(value float64) *InterleavedBuffer {
	i.p.Call("setDynamic", value)
	return i
}

// Copy TODO description.
func (i *InterleavedBuffer) Copy(source float64) *InterleavedBuffer {
	i.p.Call("copy", source)
	return i
}

// CopyAt TODO description.
func (i *InterleavedBuffer) CopyAt(index1, attribute, index2 float64) *InterleavedBuffer {
	i.p.Call("copyAt", index1, attribute, index2)
	return i
}

// Set TODO description.
func (i *InterleavedBuffer) Set(value, offset float64) *InterleavedBuffer {
	i.p.Call("set", value, offset)
	return i
}

// Clone TODO description.
func (i *InterleavedBuffer) Clone() *InterleavedBuffer {
	i.p.Call("clone")
	return i
}