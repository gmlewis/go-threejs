// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Vector2 represents a vector2.
type Vector2 struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (v *Vector2) JSObject() *js.Object { return v.p }

// Vector2 returns a Vector2 JavaScript class.
func (t *Three) Vector2() *Vector2 {
	p := t.ctx.Get("Vector2")
	return Vector2FromJSObject(p)
}

// Vector2FromJSObject returns a wrapped Vector2 JavaScript class.
func Vector2FromJSObject(p *js.Object) *Vector2 {
	return &Vector2{p: p}
}

// NewVector2 returns a new Vector2 object.
func (t *Three) NewVector2(x, y float64) *Vector2 {
	p := t.ctx.Get("Vector2").New(x, y)
	return Vector2FromJSObject(p)
}

// Width returns the width-component of the Vector2.
func (v *Vector2) Width() float64 {
	return v.p.Get("width").Float()
}

// Height returns the height-component of the Vector2.
func (v *Vector2) Height() float64 {
	return v.p.Get("height").Float()
}

// SetWidth sets the width-component of the Vector2.
func (v *Vector2) SetWidth(value float64) *Vector2 {
	v.p.Set("width", value)
	return v
}

// SetHeight sets the height-component of the Vector2.
func (v *Vector2) SetHeight(value float64) *Vector2 {
	v.p.Set("height", value)
	return v
}

// Set sets the vector values.
func (v *Vector2) Set(x, y float64) *Vector2 {
	v.p.Call("set", x, y)
	return v
}

// SetScalar sets all components of v to scalar.
func (v *Vector2) SetScalar(scalar float64) *Vector2 {
	v.p.Call("setScalar", scalar)
	return v
}

// SetX sets the X component of the vector.
func (v *Vector2) SetX(x float64) *Vector2 {
	v.p.Call("setX", x)
	return v
}

// SetY sets the Y component of the vector.
func (v *Vector2) SetY(y float64) *Vector2 {
	v.p.Call("setY", y)
	return v
}

// SetComponent sets the indexed component of the vector.
func (v *Vector2) SetComponent(index int, value float64) *Vector2 {
	v.p.Call("setComponent", index, value)
	return v
}

// GetComponent returns the indexed component of the vector.
func (v *Vector2) GetComponent(index int) float64 {
	return v.p.Call("getComponent", index).Float()
}

// Clone TODO description.
func (v *Vector2) Clone() *Vector2 {
	v.p.Call("clone")
	return v
}

// Copy copies the src vector to v.
func (v *Vector2) Copy(src *Vector2) *Vector2 {
	v.p.Call("copy", src.p)
	return v
}

// Add adds the src vector to v.
func (v *Vector2) Add(src *Vector2) *Vector2 {
	v.p.Call("add", src.p)
	return v
}

// AddScalar adds s to the components of v.
func (v *Vector2) AddScalar(s float64) *Vector2 {
	v.p.Call("addScalar", s)
	return v
}

// AddVectors adds vectors a and b and stores the result in v.
func (v *Vector2) AddVectors(a, b *Vector2) *Vector2 {
	v.p.Call("addVectors", a.p, b.p)
	return v
}

// AddScaledVector adds scaled vector src to v.
func (v *Vector2) AddScaledVector(src *Vector2, s float64) *Vector2 {
	v.p.Call("addScaledVector", src.p, s)
	return v
}

// Sub subtracts src from v and stores the result in v.
func (v *Vector2) Sub(src *Vector2) *Vector2 {
	v.p.Call("sub", src.p)
	return v
}

// SubScalar subtracts s from the components of v.
func (v *Vector2) SubScalar(s float64) *Vector2 {
	v.p.Call("subScalar", s)
	return v
}

// SubVectors subtracts b from a and stores the result in v.
func (v *Vector2) SubVectors(a, b *Vector2) *Vector2 {
	v.p.Call("subVectors", a.p, b.p)
	return v
}

// Multiply multiplies vector v by the src vector.
func (v *Vector2) Multiply(src *Vector2) *Vector2 {
	v.p.Call("multiply", src.p)
	return v
}

// MultiplyScalar multiplies the components of v by scalar.
func (v *Vector2) MultiplyScalar(scalar float64) *Vector2 {
	v.p.Call("multiplyScalar", scalar)
	return v
}

// Divide divides vector v by the src vector.
func (v *Vector2) Divide(src *Vector2) *Vector2 {
	v.p.Call("divide", src.p)
	return v
}

// DivideScalar divides the components of v by scalar.
func (v *Vector2) DivideScalar(scalar float64) *Vector2 {
	v.p.Call("divideScalar", scalar)
	return v
}

// Min sets the components of v to the min of v and src.
func (v *Vector2) Min(src *Vector2) *Vector2 {
	v.p.Call("min", src.p)
	return v
}

// Max sets the components of v to the max of v and src.
func (v *Vector2) Max(src *Vector2) *Vector2 {
	v.p.Call("max", src.p)
	return v
}

// Clamp clamps the components of v between min and max.
//
// Note that the min components must be less than the max components
// or this function will not operate correctly.
func (v *Vector2) Clamp(min, max *Vector2) *Vector2 {
	v.p.Call("clamp", min.p, max.p)
	return v
}

// ClampScalar TODO description.
func (v *Vector2) ClampScalar() *Vector2 {
	v.p.Call("clampScalar")
	return v
}

// ClampLength clamps length of v between min and max.
func (v *Vector2) ClampLength(min, max float64) *Vector2 {
	v.p.Call("clampLength", min, max)
	return v
}

// Floor calls floor on the components of v.
func (v *Vector2) Floor() *Vector2 {
	v.p.Call("floor")
	return v
}

// Ceil calls ceil on the components of v.
func (v *Vector2) Ceil() *Vector2 {
	v.p.Call("ceil")
	return v
}

// Round rounds the components of v.
func (v *Vector2) Round() *Vector2 {
	v.p.Call("round")
	return v
}

// RoundToZero rounds the components of v to zero.
func (v *Vector2) RoundToZero() *Vector2 {
	v.p.Call("roundToZero")
	return v
}

// Negate negates the components of v.
func (v *Vector2) Negate() *Vector2 {
	v.p.Call("negate")
	return v
}

// Dot returns the dot product of (v*src)
func (v *Vector2) Dot(src *Vector2) float64 {
	return v.p.Call("dot", src.p).Float()
}

// LengthSq returns the length of v squared.
func (v *Vector2) LengthSq() *Vector2 {
	v.p.Call("lengthSq")
	return v
}

// Length returns the length of v.
func (v *Vector2) Length() *Vector2 {
	v.p.Call("length")
	return v
}

// LengthManhattan returns the manhattan length of v.
func (v *Vector2) LengthManhattan() *Vector2 {
	v.p.Call("lengthManhattan")
	return v
}

// Normalize normalizes v to unit length.
func (v *Vector2) Normalize() *Vector2 {
	v.p.Call("normalize")
	return v
}

// Angle returns the angle in radians with respect to the positive X-axis.
func (v *Vector2) Angle() float64 {
	return v.p.Call("angle").Float()
}

// DistanceTo returns the distance from v to src.
func (v *Vector2) DistanceTo(src *Vector2) float64 {
	return v.p.Call("distanceTo", src.p).Float()
}

// DistanceToSquared returns the squared distance from v to src.
func (v *Vector2) DistanceToSquared(src *Vector2) float64 {
	return v.p.Call("distanceToSquared", src.p).Float()
}

// SetLength sets the length of v.
func (v *Vector2) SetLength(length float64) *Vector2 {
	v.p.Call("setLength", length)
	return v
}

// Lerp linearly interpolates between v and src by alpha.
//
// alpha = 0 returns v unmodified.
// alpha = 1 returns v set to src.
func (v *Vector2) Lerp(src *Vector2, alpha float64) *Vector2 {
	v.p.Call("lerp", src.p, alpha)
	return v
}

// LerpVectors linearly interpolates between v1 and v2 by alpha.
//
// alpha = 0 sets v to v1.
// alpha = 1 sets v to v2.
func (v *Vector2) LerpVectors(v1, v2 *Vector2, alpha float64) *Vector2 {
	v.p.Call("lerpVectors", v1.p, v2.p, alpha)
	return v
}

// Equals compares v to src and returns true if equal.
func (v *Vector2) Equals(src *Vector2) bool {
	return v.p.Call("equals", src.p).Bool()
}

// FromArray sets v to the offset component of array.
func (v *Vector2) FromArray(array []float64, offset float64) *Vector2 {
	v.p.Call("fromArray", array, offset)
	return v
}

/* TODO
// ToArray sets the offset components of array from v.
func (v *Vector2) ToArray(array []float64, offset float64) []float64 {
	return v.p.Call("toArray", array, offset)
}
*/

// FromAttribute TODO description.
func (v *Vector2) FromAttribute(attribute, index, offset float64) *Vector2 {
	v.p.Call("fromAttribute", attribute, index, offset)
	return v
}

// RotateAround rotates v around center by angle radians.
func (v *Vector2) RotateAround(center *Vector2, angle float64) *Vector2 {
	v.p.Call("rotateAround", center.p, angle)
	return v
}
