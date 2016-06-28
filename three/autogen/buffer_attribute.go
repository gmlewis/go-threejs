package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// BufferAttribute represents a bufferattribute.
type BufferAttribute struct{ p *js.Object }

// BufferAttribute returns a bufferattribute object.
func (t *Three) BufferAttribute() *BufferAttribute {
	p := t.ctx.Get("BufferAttribute")
	return &BufferAttribute{p: p}
}

// NewBufferAttribute returns a new bufferattribute object.
func (t *BufferAttribute) New(array, itemSize float64) *BufferAttribute {
	p := t.p.New(array, itemSize)
	return &BufferAttribute{p: p}
}

// Int8Attribute represents an int8attribute.
type Int8Attribute struct{ p *js.Object }

// Int8Attribute returns an int8attribute object.
func (t *Three) Int8Attribute() *Int8Attribute {
	p := t.ctx.Get("Int8Attribute")
	return &Int8Attribute{p: p}
}

// NewInt8Attribute returns a new int8attribute object.
func (t *Int8Attribute) New(array, itemSize float64) *Int8Attribute {
	p := t.p.New(array, itemSize)
	return &Int8Attribute{p: p}
}

// Uint8Attribute represents an uint8attribute.
type Uint8Attribute struct{ p *js.Object }

// Uint8Attribute returns an uint8attribute object.
func (t *Three) Uint8Attribute() *Uint8Attribute {
	p := t.ctx.Get("Uint8Attribute")
	return &Uint8Attribute{p: p}
}

// NewUint8Attribute returns a new uint8attribute object.
func (t *Uint8Attribute) New(array, itemSize float64) *Uint8Attribute {
	p := t.p.New(array, itemSize)
	return &Uint8Attribute{p: p}
}

// Uint8ClampedAttribute represents an uint8clampedattribute.
type Uint8ClampedAttribute struct{ p *js.Object }

// Uint8ClampedAttribute returns an uint8clampedattribute object.
func (t *Three) Uint8ClampedAttribute() *Uint8ClampedAttribute {
	p := t.ctx.Get("Uint8ClampedAttribute")
	return &Uint8ClampedAttribute{p: p}
}

// NewUint8ClampedAttribute returns a new uint8clampedattribute object.
func (t *Uint8ClampedAttribute) New(array, itemSize float64) *Uint8ClampedAttribute {
	p := t.p.New(array, itemSize)
	return &Uint8ClampedAttribute{p: p}
}

// Int16Attribute represents an int16attribute.
type Int16Attribute struct{ p *js.Object }

// Int16Attribute returns an int16attribute object.
func (t *Three) Int16Attribute() *Int16Attribute {
	p := t.ctx.Get("Int16Attribute")
	return &Int16Attribute{p: p}
}

// NewInt16Attribute returns a new int16attribute object.
func (t *Int16Attribute) New(array, itemSize float64) *Int16Attribute {
	p := t.p.New(array, itemSize)
	return &Int16Attribute{p: p}
}

// Uint16Attribute represents an uint16attribute.
type Uint16Attribute struct{ p *js.Object }

// Uint16Attribute returns an uint16attribute object.
func (t *Three) Uint16Attribute() *Uint16Attribute {
	p := t.ctx.Get("Uint16Attribute")
	return &Uint16Attribute{p: p}
}

// NewUint16Attribute returns a new uint16attribute object.
func (t *Uint16Attribute) New(array, itemSize float64) *Uint16Attribute {
	p := t.p.New(array, itemSize)
	return &Uint16Attribute{p: p}
}

// Int32Attribute represents an int32attribute.
type Int32Attribute struct{ p *js.Object }

// Int32Attribute returns an int32attribute object.
func (t *Three) Int32Attribute() *Int32Attribute {
	p := t.ctx.Get("Int32Attribute")
	return &Int32Attribute{p: p}
}

// NewInt32Attribute returns a new int32attribute object.
func (t *Int32Attribute) New(array, itemSize float64) *Int32Attribute {
	p := t.p.New(array, itemSize)
	return &Int32Attribute{p: p}
}

// Uint32Attribute represents an uint32attribute.
type Uint32Attribute struct{ p *js.Object }

// Uint32Attribute returns an uint32attribute object.
func (t *Three) Uint32Attribute() *Uint32Attribute {
	p := t.ctx.Get("Uint32Attribute")
	return &Uint32Attribute{p: p}
}

// NewUint32Attribute returns a new uint32attribute object.
func (t *Uint32Attribute) New(array, itemSize float64) *Uint32Attribute {
	p := t.p.New(array, itemSize)
	return &Uint32Attribute{p: p}
}

// Float32Attribute represents a float32attribute.
type Float32Attribute struct{ p *js.Object }

// Float32Attribute returns a float32attribute object.
func (t *Three) Float32Attribute() *Float32Attribute {
	p := t.ctx.Get("Float32Attribute")
	return &Float32Attribute{p: p}
}

// NewFloat32Attribute returns a new float32attribute object.
func (t *Float32Attribute) New(array, itemSize float64) *Float32Attribute {
	p := t.p.New(array, itemSize)
	return &Float32Attribute{p: p}
}

// Float64Attribute represents a float64attribute.
type Float64Attribute struct{ p *js.Object }

// Float64Attribute returns a float64attribute object.
func (t *Three) Float64Attribute() *Float64Attribute {
	p := t.ctx.Get("Float64Attribute")
	return &Float64Attribute{p: p}
}

// NewFloat64Attribute returns a new float64attribute object.
func (t *Float64Attribute) New(array, itemSize float64) *Float64Attribute {
	p := t.p.New(array, itemSize)
	return &Float64Attribute{p: p}
}

// DynamicBufferAttribute represents a dynamicbufferattribute.
type DynamicBufferAttribute struct{ p *js.Object }

// DynamicBufferAttribute returns a dynamicbufferattribute object.
func (t *Three) DynamicBufferAttribute() *DynamicBufferAttribute {
	p := t.ctx.Get("DynamicBufferAttribute")
	return &DynamicBufferAttribute{p: p}
}

// NewDynamicBufferAttribute returns a new dynamicbufferattribute object.
func (t *DynamicBufferAttribute) New(array, itemSize float64) *DynamicBufferAttribute {
	p := t.p.New(array, itemSize)
	return &DynamicBufferAttribute{p: p}
}

// Count returns the count-component of the DynamicBufferAttribute.
func (d *DynamicBufferAttribute) Count() float64 {
	return d.p.Get("count").Float()
}

// SetNeedsUpdate sets the needsUpdate-component of the DynamicBufferAttribute.
func (d *DynamicBufferAttribute) SetNeedsUpdate(value float64) *DynamicBufferAttribute {
	d.p.Set("needsUpdate", value)
	return d
}

// SetDynamic TODO description.
func (d *DynamicBufferAttribute) SetDynamic(value float64) *DynamicBufferAttribute {
	d.p.Call("setDynamic", value)
	return d
}

// Copy TODO description.
func (d *DynamicBufferAttribute) Copy(source float64) *DynamicBufferAttribute {
	d.p.Call("copy", source)
	return d
}

// CopyAt TODO description.
func (d *DynamicBufferAttribute) CopyAt(index1, attribute, index2 float64) *DynamicBufferAttribute {
	d.p.Call("copyAt", index1, attribute, index2)
	return d
}

// CopyArray TODO description.
func (d *DynamicBufferAttribute) CopyArray(array float64) *DynamicBufferAttribute {
	d.p.Call("copyArray", array)
	return d
}

// CopyColorsArray TODO description.
func (d *DynamicBufferAttribute) CopyColorsArray(colors float64) *DynamicBufferAttribute {
	d.p.Call("copyColorsArray", colors)
	return d
}

// CopyIndicesArray TODO description.
func (d *DynamicBufferAttribute) CopyIndicesArray(indices float64) *DynamicBufferAttribute {
	d.p.Call("copyIndicesArray", indices)
	return d
}

// CopyVector2sArray TODO description.
func (d *DynamicBufferAttribute) CopyVector2sArray(vectors float64) *DynamicBufferAttribute {
	d.p.Call("copyVector2sArray", vectors)
	return d
}

// CopyVector3sArray TODO description.
func (d *DynamicBufferAttribute) CopyVector3sArray(vectors float64) *DynamicBufferAttribute {
	d.p.Call("copyVector3sArray", vectors)
	return d
}

// CopyVector4sArray TODO description.
func (d *DynamicBufferAttribute) CopyVector4sArray(vectors float64) *DynamicBufferAttribute {
	d.p.Call("copyVector4sArray", vectors)
	return d
}

// Set TODO description.
func (d *DynamicBufferAttribute) Set(value, offset float64) *DynamicBufferAttribute {
	d.p.Call("set", value, offset)
	return d
}

// GetX TODO description.
func (d *DynamicBufferAttribute) GetX(index float64) *DynamicBufferAttribute {
	d.p.Call("getX", index)
	return d
}

// SetX TODO description.
func (d *DynamicBufferAttribute) SetX(index, x float64) *DynamicBufferAttribute {
	d.p.Call("setX", index, x)
	return d
}

// GetY TODO description.
func (d *DynamicBufferAttribute) GetY(index float64) *DynamicBufferAttribute {
	d.p.Call("getY", index)
	return d
}

// SetY TODO description.
func (d *DynamicBufferAttribute) SetY(index, y float64) *DynamicBufferAttribute {
	d.p.Call("setY", index, y)
	return d
}

// GetZ TODO description.
func (d *DynamicBufferAttribute) GetZ(index float64) *DynamicBufferAttribute {
	d.p.Call("getZ", index)
	return d
}

// SetZ TODO description.
func (d *DynamicBufferAttribute) SetZ(index, z float64) *DynamicBufferAttribute {
	d.p.Call("setZ", index, z)
	return d
}

// GetW TODO description.
func (d *DynamicBufferAttribute) GetW(index float64) *DynamicBufferAttribute {
	d.p.Call("getW", index)
	return d
}

// SetW TODO description.
func (d *DynamicBufferAttribute) SetW(index, w float64) *DynamicBufferAttribute {
	d.p.Call("setW", index, w)
	return d
}

// SetXY TODO description.
func (d *DynamicBufferAttribute) SetXY(index, x, y float64) *DynamicBufferAttribute {
	d.p.Call("setXY", index, x, y)
	return d
}

// SetXYZ TODO description.
func (d *DynamicBufferAttribute) SetXYZ(index, x, y, z float64) *DynamicBufferAttribute {
	d.p.Call("setXYZ", index, x, y, z)
	return d
}

// SetXYZW TODO description.
func (d *DynamicBufferAttribute) SetXYZW(index, x, y, z, w float64) *DynamicBufferAttribute {
	d.p.Call("setXYZW", index, x, y, z, w)
	return d
}

// Clone TODO description.
func (d *DynamicBufferAttribute) Clone() *DynamicBufferAttribute {
	d.p.Call("clone")
	return d
}

