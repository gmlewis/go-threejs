// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// BufferAttribute represents a bufferattribute.
type BufferAttribute struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (b *BufferAttribute) JSObject() *js.Object { return b.p }

// BufferAttribute returns a BufferAttribute object.
func (t *Three) BufferAttribute() *BufferAttribute {
	p := t.ctx.Get("BufferAttribute")
	return &BufferAttribute{p: p}
}

// New returns a new BufferAttribute object.
func (b *BufferAttribute) New(array, itemSize float64) *BufferAttribute {
	p := b.p.New(array, itemSize)
	return &BufferAttribute{p: p}
}

// Int8Attribute represents an int8attribute.
type Int8Attribute struct{ p *js.Object }

// Int8Attribute returns an Int8Attribute object.
func (t *Three) Int8Attribute() *Int8Attribute {
	p := t.ctx.Get("Int8Attribute")
	return &Int8Attribute{p: p}
}

// New returns a new Int8Attribute object.
func (i *Int8Attribute) New(array, itemSize float64) *Int8Attribute {
	p := i.p.New(array, itemSize)
	return &Int8Attribute{p: p}
}

// Uint8Attribute represents an uint8attribute.
type Uint8Attribute struct{ p *js.Object }

// Uint8Attribute returns an Uint8Attribute object.
func (t *Three) Uint8Attribute() *Uint8Attribute {
	p := t.ctx.Get("Uint8Attribute")
	return &Uint8Attribute{p: p}
}

// New returns a new Uint8Attribute object.
func (u *Uint8Attribute) New(array, itemSize float64) *Uint8Attribute {
	p := u.p.New(array, itemSize)
	return &Uint8Attribute{p: p}
}

// Uint8ClampedAttribute represents an uint8clampedattribute.
type Uint8ClampedAttribute struct{ p *js.Object }

// Uint8ClampedAttribute returns an Uint8ClampedAttribute object.
func (t *Three) Uint8ClampedAttribute() *Uint8ClampedAttribute {
	p := t.ctx.Get("Uint8ClampedAttribute")
	return &Uint8ClampedAttribute{p: p}
}

// New returns a new Uint8ClampedAttribute object.
func (u *Uint8ClampedAttribute) New(array, itemSize float64) *Uint8ClampedAttribute {
	p := u.p.New(array, itemSize)
	return &Uint8ClampedAttribute{p: p}
}

// Int16Attribute represents an int16attribute.
type Int16Attribute struct{ p *js.Object }

// Int16Attribute returns an Int16Attribute object.
func (t *Three) Int16Attribute() *Int16Attribute {
	p := t.ctx.Get("Int16Attribute")
	return &Int16Attribute{p: p}
}

// New returns a new Int16Attribute object.
func (i *Int16Attribute) New(array, itemSize float64) *Int16Attribute {
	p := i.p.New(array, itemSize)
	return &Int16Attribute{p: p}
}

// Uint16Attribute represents an uint16attribute.
type Uint16Attribute struct{ p *js.Object }

// Uint16Attribute returns an Uint16Attribute object.
func (t *Three) Uint16Attribute() *Uint16Attribute {
	p := t.ctx.Get("Uint16Attribute")
	return &Uint16Attribute{p: p}
}

// New returns a new Uint16Attribute object.
func (u *Uint16Attribute) New(array, itemSize float64) *Uint16Attribute {
	p := u.p.New(array, itemSize)
	return &Uint16Attribute{p: p}
}

// Int32Attribute represents an int32attribute.
type Int32Attribute struct{ p *js.Object }

// Int32Attribute returns an Int32Attribute object.
func (t *Three) Int32Attribute() *Int32Attribute {
	p := t.ctx.Get("Int32Attribute")
	return &Int32Attribute{p: p}
}

// New returns a new Int32Attribute object.
func (i *Int32Attribute) New(array, itemSize float64) *Int32Attribute {
	p := i.p.New(array, itemSize)
	return &Int32Attribute{p: p}
}

// Uint32Attribute represents an uint32attribute.
type Uint32Attribute struct{ p *js.Object }

// Uint32Attribute returns an Uint32Attribute object.
func (t *Three) Uint32Attribute() *Uint32Attribute {
	p := t.ctx.Get("Uint32Attribute")
	return &Uint32Attribute{p: p}
}

// New returns a new Uint32Attribute object.
func (u *Uint32Attribute) New(array, itemSize float64) *Uint32Attribute {
	p := u.p.New(array, itemSize)
	return &Uint32Attribute{p: p}
}

// Float32Attribute represents a float32attribute.
type Float32Attribute struct{ p *js.Object }

// Float32Attribute returns a Float32Attribute object.
func (t *Three) Float32Attribute() *Float32Attribute {
	p := t.ctx.Get("Float32Attribute")
	return &Float32Attribute{p: p}
}

// New returns a new Float32Attribute object.
func (f *Float32Attribute) New(array, itemSize float64) *Float32Attribute {
	p := f.p.New(array, itemSize)
	return &Float32Attribute{p: p}
}

// Float64Attribute represents a float64attribute.
type Float64Attribute struct{ p *js.Object }

// Float64Attribute returns a Float64Attribute object.
func (t *Three) Float64Attribute() *Float64Attribute {
	p := t.ctx.Get("Float64Attribute")
	return &Float64Attribute{p: p}
}

// New returns a new Float64Attribute object.
func (f *Float64Attribute) New(array, itemSize float64) *Float64Attribute {
	p := f.p.New(array, itemSize)
	return &Float64Attribute{p: p}
}

// Count returns the count-component of the BufferAttribute.
func (d *BufferAttribute) Count() float64 {
	return d.p.Get("count").Float()
}

// SetNeedsUpdate sets the needsUpdate-component of the BufferAttribute.
func (d *BufferAttribute) SetNeedsUpdate(value float64) *BufferAttribute {
	d.p.Set("needsUpdate", value)
	return d
}

// SetDynamic TODO description.
func (d *BufferAttribute) SetDynamic(value float64) *BufferAttribute {
	d.p.Call("setDynamic", value)
	return d
}

// Copy TODO description.
func (d *BufferAttribute) Copy(source float64) *BufferAttribute {
	d.p.Call("copy", source)
	return d
}

// CopyAt TODO description.
func (d *BufferAttribute) CopyAt(index1, attribute, index2 float64) *BufferAttribute {
	d.p.Call("copyAt", index1, attribute, index2)
	return d
}

// CopyArray TODO description.
func (d *BufferAttribute) CopyArray(array float64) *BufferAttribute {
	d.p.Call("copyArray", array)
	return d
}

// CopyColorsArray TODO description.
func (d *BufferAttribute) CopyColorsArray(colors float64) *BufferAttribute {
	d.p.Call("copyColorsArray", colors)
	return d
}

// CopyIndicesArray TODO description.
func (d *BufferAttribute) CopyIndicesArray(indices float64) *BufferAttribute {
	d.p.Call("copyIndicesArray", indices)
	return d
}

// CopyVector2sArray TODO description.
func (d *BufferAttribute) CopyVector2sArray(vectors float64) *BufferAttribute {
	d.p.Call("copyVector2sArray", vectors)
	return d
}

// CopyVector3sArray TODO description.
func (d *BufferAttribute) CopyVector3sArray(vectors float64) *BufferAttribute {
	d.p.Call("copyVector3sArray", vectors)
	return d
}

// CopyVector4sArray TODO description.
func (d *BufferAttribute) CopyVector4sArray(vectors float64) *BufferAttribute {
	d.p.Call("copyVector4sArray", vectors)
	return d
}

// Set TODO description.
func (d *BufferAttribute) Set(value, offset float64) *BufferAttribute {
	d.p.Call("set", value, offset)
	return d
}

// GetX TODO description.
func (d *BufferAttribute) GetX(index float64) *BufferAttribute {
	d.p.Call("getX", index)
	return d
}

// SetX TODO description.
func (d *BufferAttribute) SetX(index, x float64) *BufferAttribute {
	d.p.Call("setX", index, x)
	return d
}

// GetY TODO description.
func (d *BufferAttribute) GetY(index float64) *BufferAttribute {
	d.p.Call("getY", index)
	return d
}

// SetY TODO description.
func (d *BufferAttribute) SetY(index, y float64) *BufferAttribute {
	d.p.Call("setY", index, y)
	return d
}

// GetZ TODO description.
func (d *BufferAttribute) GetZ(index float64) *BufferAttribute {
	d.p.Call("getZ", index)
	return d
}

// SetZ TODO description.
func (d *BufferAttribute) SetZ(index, z float64) *BufferAttribute {
	d.p.Call("setZ", index, z)
	return d
}

// GetW TODO description.
func (d *BufferAttribute) GetW(index float64) *BufferAttribute {
	d.p.Call("getW", index)
	return d
}

// SetW TODO description.
func (d *BufferAttribute) SetW(index, w float64) *BufferAttribute {
	d.p.Call("setW", index, w)
	return d
}

// SetXY TODO description.
func (d *BufferAttribute) SetXY(index, x, y float64) *BufferAttribute {
	d.p.Call("setXY", index, x, y)
	return d
}

// SetXYZ TODO description.
func (d *BufferAttribute) SetXYZ(index, x, y, z float64) *BufferAttribute {
	d.p.Call("setXYZ", index, x, y, z)
	return d
}

// SetXYZW TODO description.
func (d *BufferAttribute) SetXYZW(index, x, y, z, w float64) *BufferAttribute {
	d.p.Call("setXYZW", index, x, y, z, w)
	return d
}

// Clone TODO description.
func (d *BufferAttribute) Clone() *BufferAttribute {
	d.p.Call("clone")
	return d
}
