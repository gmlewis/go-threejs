// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// InterleavedBufferAttribute represents an interleavedbufferattribute.
type InterleavedBufferAttribute struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (i *InterleavedBufferAttribute) JSObject() *js.Object { return i.p }

// InterleavedBufferAttribute returns an InterleavedBufferAttribute object.
func (t *Three) InterleavedBufferAttribute() *InterleavedBufferAttribute {
	p := t.ctx.Get("InterleavedBufferAttribute")
	return &InterleavedBufferAttribute{p: p}
}

// New returns a new InterleavedBufferAttribute object.
func (i *InterleavedBufferAttribute) New(interleavedBuffer, itemSize, offset float64) *InterleavedBufferAttribute {
	p := i.p.New(interleavedBuffer, itemSize, offset)
	return &InterleavedBufferAttribute{p: p}
}

// Length returns the length-component of the InterleavedBufferAttribute.
func (i *InterleavedBufferAttribute) Length() float64 {
	return i.p.Get("length").Float()
}

// Count returns the count-component of the InterleavedBufferAttribute.
func (i *InterleavedBufferAttribute) Count() float64 {
	return i.p.Get("count").Float()
}

// SetX TODO description.
func (i *InterleavedBufferAttribute) SetX(index, x float64) *InterleavedBufferAttribute {
	i.p.Call("setX", index, x)
	return i
}

// SetY TODO description.
func (i *InterleavedBufferAttribute) SetY(index, y float64) *InterleavedBufferAttribute {
	i.p.Call("setY", index, y)
	return i
}

// SetZ TODO description.
func (i *InterleavedBufferAttribute) SetZ(index, z float64) *InterleavedBufferAttribute {
	i.p.Call("setZ", index, z)
	return i
}

// SetW TODO description.
func (i *InterleavedBufferAttribute) SetW(index, w float64) *InterleavedBufferAttribute {
	i.p.Call("setW", index, w)
	return i
}

// GetX TODO description.
func (i *InterleavedBufferAttribute) GetX(index float64) *InterleavedBufferAttribute {
	i.p.Call("getX", index)
	return i
}

// GetY TODO description.
func (i *InterleavedBufferAttribute) GetY(index float64) *InterleavedBufferAttribute {
	i.p.Call("getY", index)
	return i
}

// GetZ TODO description.
func (i *InterleavedBufferAttribute) GetZ(index float64) *InterleavedBufferAttribute {
	i.p.Call("getZ", index)
	return i
}

// GetW TODO description.
func (i *InterleavedBufferAttribute) GetW(index float64) *InterleavedBufferAttribute {
	i.p.Call("getW", index)
	return i
}

// SetXY TODO description.
func (i *InterleavedBufferAttribute) SetXY(index, x, y float64) *InterleavedBufferAttribute {
	i.p.Call("setXY", index, x, y)
	return i
}

// SetXYZ TODO description.
func (i *InterleavedBufferAttribute) SetXYZ(index, x, y, z float64) *InterleavedBufferAttribute {
	i.p.Call("setXYZ", index, x, y, z)
	return i
}

// SetXYZW TODO description.
func (i *InterleavedBufferAttribute) SetXYZW(index, x, y, z, w float64) *InterleavedBufferAttribute {
	i.p.Call("setXYZW", index, x, y, z, w)
	return i
}
