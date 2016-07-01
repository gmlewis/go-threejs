// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// BufferAttribute stores data for an attribute associated with a BufferGeometry.
// See that page for details and a usage example. This class is used to store
// builtin attributes such as vertex position, normals, color, etc., but can also
// be used in your code to store custom attributes in a BufferGeometry.
//
// http://threejs.org/docs/index.html#Reference/Core/BufferAttribute
type BufferAttribute struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (b *BufferAttribute) JSObject() *js.Object { return b.p }

// BufferAttribute returns a BufferAttribute JavaScript class.
func (t *Three) BufferAttribute() *BufferAttribute {
	p := t.ctx.Get("BufferAttribute")
	return &BufferAttribute{p: p}
}

// NewBufferAttribute returns a new BufferAttribute object.
func (t *Three) NewBufferAttribute(array []float64, itemSize int) *BufferAttribute {
	p := t.ctx.Get("BufferAttribute").New(array, itemSize)
	return &BufferAttribute{p: p}
}

// Int8Attribute represents an int8attribute.
type Int8Attribute struct{ p *js.Object }

// Int8Attribute returns an Int8Attribute object.
func (t *Three) Int8Attribute() *Int8Attribute {
	p := t.ctx.Get("Int8Attribute")
	return &Int8Attribute{p: p}
}

// NewInt8Attribute returns a new Int8Attribute object.
func (t *Three) NewInt8Attribute(array []int8, itemSize int) *Int8Attribute {
	p := t.ctx.Get("Int8Attribute").New(array, itemSize)
	return &Int8Attribute{p: p}
}

// Uint8Attribute represents an uint8attribute.
type Uint8Attribute struct{ p *js.Object }

// Uint8Attribute returns an Uint8Attribute object.
func (t *Three) Uint8Attribute() *Uint8Attribute {
	p := t.ctx.Get("Uint8Attribute")
	return &Uint8Attribute{p: p}
}

// NewUint8Attribute returns a new Uint8Attribute object.
func (t *Three) NewUint8Attribute(array []uint8, itemSize int) *Uint8Attribute {
	p := t.ctx.Get("Uint8Attribute").New(array, itemSize)
	return &Uint8Attribute{p: p}
}

// Uint8ClampedAttribute represents an uint8clampedattribute.
type Uint8ClampedAttribute struct{ p *js.Object }

// Uint8ClampedAttribute returns an Uint8ClampedAttribute object.
func (t *Three) Uint8ClampedAttribute() *Uint8ClampedAttribute {
	p := t.ctx.Get("Uint8ClampedAttribute")
	return &Uint8ClampedAttribute{p: p}
}

// NewUint8ClampedAttribute returns a new Uint8ClampedAttribute object.
func (t *Three) NewUint8ClampedAttribute(array []uint8, itemSize int) *Uint8ClampedAttribute {
	p := t.ctx.Get("Uint8ClampedAttribute").New(array, itemSize)
	return &Uint8ClampedAttribute{p: p}
}

// Int16Attribute represents an int16attribute.
type Int16Attribute struct{ p *js.Object }

// Int16Attribute returns an Int16Attribute object.
func (t *Three) Int16Attribute() *Int16Attribute {
	p := t.ctx.Get("Int16Attribute")
	return &Int16Attribute{p: p}
}

// NewInt16Attribute returns a new Int16Attribute object.
func (t *Three) NewInt16Attribute(array []int16, itemSize int) *Int16Attribute {
	p := t.ctx.Get("Int16Attribute").New(array, itemSize)
	return &Int16Attribute{p: p}
}

// Uint16Attribute represents an uint16attribute.
type Uint16Attribute struct{ p *js.Object }

// Uint16Attribute returns an Uint16Attribute object.
func (t *Three) Uint16Attribute() *Uint16Attribute {
	p := t.ctx.Get("Uint16Attribute")
	return &Uint16Attribute{p: p}
}

// NewUint16Attribute returns a new Uint16Attribute object.
func (t *Three) NewUint16Attribute(array []uint16, itemSize int) *Uint16Attribute {
	p := t.ctx.Get("Uint16Attribute").New(array, itemSize)
	return &Uint16Attribute{p: p}
}

// Int32Attribute represents an int32attribute.
type Int32Attribute struct{ p *js.Object }

// Int32Attribute returns an Int32Attribute object.
func (t *Three) Int32Attribute() *Int32Attribute {
	p := t.ctx.Get("Int32Attribute")
	return &Int32Attribute{p: p}
}

// NewInt32Attribute returns a new Int32Attribute object.
func (t *Three) NewInt32Attribute(array []int32, itemSize int) *Int32Attribute {
	p := t.ctx.Get("Int32Attribute").New(array, itemSize)
	return &Int32Attribute{p: p}
}

// Uint32Attribute represents an uint32attribute.
type Uint32Attribute struct{ p *js.Object }

// Uint32Attribute returns an Uint32Attribute object.
func (t *Three) Uint32Attribute() *Uint32Attribute {
	p := t.ctx.Get("Uint32Attribute")
	return &Uint32Attribute{p: p}
}

// NewUint32Attribute returns a new Uint32Attribute object.
func (t *Three) NewUint32Attribute(array []uint32, itemSize int) *Uint32Attribute {
	p := t.ctx.Get("Uint32Attribute").New(array, itemSize)
	return &Uint32Attribute{p: p}
}

// Float32Attribute represents a float32attribute.
type Float32Attribute struct{ p *js.Object }

// Float32Attribute returns a Float32Attribute object.
func (t *Three) Float32Attribute() *Float32Attribute {
	p := t.ctx.Get("Float32Attribute")
	return &Float32Attribute{p: p}
}

// NewFloat32Attribute returns a new Float32Attribute object.
func (t *Three) NewFloat32Attribute(array []float32, itemSize int) *Float32Attribute {
	p := t.ctx.Get("Float32Attribute").New(array, itemSize)
	return &Float32Attribute{p: p}
}

// Float64Attribute represents a float64attribute.
type Float64Attribute struct{ p *js.Object }

// Float64Attribute returns a Float64Attribute object.
func (t *Three) Float64Attribute() *Float64Attribute {
	p := t.ctx.Get("Float64Attribute")
	return &Float64Attribute{p: p}
}

// NewFloat64Attribute returns a new Float64Attribute object.
func (t *Three) NewFloat64Attribute(array []float64, itemSize int) *Float64Attribute {
	p := t.ctx.Get("Float64Attribute").New(array, itemSize)
	return &Float64Attribute{p: p}
}

// Count returns the count-component of the BufferAttribute.
func (b *BufferAttribute) Count() float64 {
	return b.p.Get("count").Float()
}

// SetNeedsUpdate sets the needsUpdate-component of the BufferAttribute.
func (b *BufferAttribute) SetNeedsUpdate(value float64) *BufferAttribute {
	b.p.Set("needsUpdate", value)
	return b
}

// SetDynamic TODO description.
func (b *BufferAttribute) SetDynamic(value float64) *BufferAttribute {
	b.p.Call("setDynamic", value)
	return b
}

// Copy TODO description.
func (b *BufferAttribute) Copy(source *BufferAttribute) *BufferAttribute {
	b.p.Call("copy", source.p)
	return b
}

// CopyAt TODO description.
func (b *BufferAttribute) CopyAt(index1, attribute, index2 float64) *BufferAttribute {
	b.p.Call("copyAt", index1, attribute, index2)
	return b
}

// CopyArray TODO description.
func (b *BufferAttribute) CopyArray(array []float64) *BufferAttribute {
	b.p.Call("copyArray", array)
	return b
}

// CopyColorsArray TODO description.
func (b *BufferAttribute) CopyColorsArray(colors []float64) *BufferAttribute {
	b.p.Call("copyColorsArray", colors)
	return b
}

// CopyIndicesArray TODO description.
func (b *BufferAttribute) CopyIndicesArray(indices []int) *BufferAttribute {
	b.p.Call("copyIndicesArray", indices)
	return b
}

// CopyVector2sArray TODO description.
func (b *BufferAttribute) CopyVector2sArray(vectors []*Vector2) *BufferAttribute {
	b.p.Call("copyVector2sArray", vectors)
	return b
}

// CopyVector3sArray TODO description.
func (b *BufferAttribute) CopyVector3sArray(vectors []*Vector3) *BufferAttribute {
	b.p.Call("copyVector3sArray", vectors)
	return b
}

// CopyVector4sArray TODO description.
func (b *BufferAttribute) CopyVector4sArray(vectors []*Vector4) *BufferAttribute {
	b.p.Call("copyVector4sArray", vectors)
	return b
}

// Set TODO description.
func (b *BufferAttribute) Set(value float64, offset int) *BufferAttribute {
	b.p.Call("set", value, offset)
	return b
}

// GetX TODO description.
func (b *BufferAttribute) GetX(index int) *BufferAttribute {
	b.p.Call("getX", index)
	return b
}

// SetX TODO description.
func (b *BufferAttribute) SetX(index int, x float64) *BufferAttribute {
	b.p.Call("setX", index, x)
	return b
}

// GetY TODO description.
func (b *BufferAttribute) GetY(index int) *BufferAttribute {
	b.p.Call("getY", index)
	return b
}

// SetY TODO description.
func (b *BufferAttribute) SetY(index int, y float64) *BufferAttribute {
	b.p.Call("setY", index, y)
	return b
}

// GetZ TODO description.
func (b *BufferAttribute) GetZ(index int) *BufferAttribute {
	b.p.Call("getZ", index)
	return b
}

// SetZ TODO description.
func (b *BufferAttribute) SetZ(index int, z float64) *BufferAttribute {
	b.p.Call("setZ", index, z)
	return b
}

// GetW TODO description.
func (b *BufferAttribute) GetW(index int) *BufferAttribute {
	b.p.Call("getW", index)
	return b
}

// SetW TODO description.
func (b *BufferAttribute) SetW(index int, w float64) *BufferAttribute {
	b.p.Call("setW", index, w)
	return b
}

// SetXY TODO description.
func (b *BufferAttribute) SetXY(index int, x, y float64) *BufferAttribute {
	b.p.Call("setXY", index, x, y)
	return b
}

// SetXYZ TODO description.
func (b *BufferAttribute) SetXYZ(index int, x, y, z float64) *BufferAttribute {
	b.p.Call("setXYZ", index, x, y, z)
	return b
}

// SetXYZW TODO description.
func (b *BufferAttribute) SetXYZW(index int, x, y, z, w float64) *BufferAttribute {
	b.p.Call("setXYZW", index, x, y, z, w)
	return b
}

// Clone TODO description.
func (b *BufferAttribute) Clone() *BufferAttribute {
	b.p.Call("clone")
	return b
}

// UUID returns the property of the same name.
func (b *BufferAttribute) UUID() int {
	return b.p.Get("uuid").Int()
}

// Array returns the property of the same name.
func (b *BufferAttribute) Array() *js.Object {
	return b.p.Get("array")
}

// ItemSize returns the property of the same name.
func (b *BufferAttribute) ItemSize() int {
	return b.p.Get("itemSize").Int()
}

// Dynamic returns the property of the same name.
func (b *BufferAttribute) Dynamic() bool {
	return b.p.Get("dynamic").Bool()
}

// UpdateRange returns the property of the same name.
func (b *BufferAttribute) UpdateRange() (offset, count int) {
	t := b.p.Get("updateRange")
	return t.Get("offset").Int(), t.Get("count").Int()
}

// Version returns the property of the same name.
func (b *BufferAttribute) Version() int {
	return b.p.Get("version").Int()
}
