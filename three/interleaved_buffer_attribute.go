// Copyright 2016 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// InterleavedBufferAttribute represents an interleavedbufferattribute.
type InterleavedBufferAttribute struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (i *InterleavedBufferAttribute) JSObject() *js.Object { return i.p }

// InterleavedBufferAttribute returns an InterleavedBufferAttribute JavaScript class.
func (t *Three) InterleavedBufferAttribute() *InterleavedBufferAttribute {
	p := t.ctx.Get("InterleavedBufferAttribute")
	return InterleavedBufferAttributeFromJSObject(p)
}

// InterleavedBufferAttributeFromJSObject returns a wrapped InterleavedBufferAttribute JavaScript class.
func InterleavedBufferAttributeFromJSObject(p *js.Object) *InterleavedBufferAttribute {
	return &InterleavedBufferAttribute{p: p}
}

// NewInterleavedBufferAttribute returns a new InterleavedBufferAttribute object.
func (t *Three) NewInterleavedBufferAttribute(interleavedBuffer, itemSize, offset float64) *InterleavedBufferAttribute {
	p := t.ctx.Get("InterleavedBufferAttribute").New(interleavedBuffer, itemSize, offset)
	return InterleavedBufferAttributeFromJSObject(p)
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
