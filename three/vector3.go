// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Vector3 represents a three-dimensional vector.
//
// http://threejs.org/docs/index.html#Reference/Math/Vector3
type Vector3 struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (v *Vector3) JSObject() *js.Object { return v.p }

// Vector3 returns a Vector3 JavaScript class.
func (t *Three) Vector3() *Vector3 {
	p := t.ctx.Get("Vector3")
	return vector3(p)
}

// vector3 returns a wrapped Vector3 JavaScript class.
func vector3(p *js.Object) *Vector3 {
	return &Vector3{p: p}
}

// NewVector3 returns a new Vector3 object.
func (t *Three) NewVector3(x, y, z float64) *Vector3 {
	p := t.ctx.Get("Vector3").New(x, y, z)
	return vector3(p)
}

// Set sets the vector values.
func (v *Vector3) Set(x, y, z float64) *Vector3 {
	v.p.Call("set", x, y, z)
	return v
}

// SetScalar sets all components of v to scalar.
func (v *Vector3) SetScalar(scalar float64) *Vector3 {
	v.p.Call("setScalar", scalar)
	return v
}

// SetX sets the X component of the vector.
func (v *Vector3) SetX(x float64) *Vector3 {
	v.p.Call("setX", x)
	return v
}

// SetY sets the Y component of the vector.
func (v *Vector3) SetY(y float64) *Vector3 {
	v.p.Call("setY", y)
	return v
}

// SetZ sets the Z component of the vector.
func (v *Vector3) SetZ(z float64) *Vector3 {
	v.p.Call("setZ", z)
	return v
}

// SetComponent sets the indexed component of the vector.
func (v *Vector3) SetComponent(index int, value float64) *Vector3 {
	v.p.Call("setComponent", index, value)
	return v
}

// GetComponent returns the indexed component of the vector.
func (v *Vector3) GetComponent(index int) *Vector3 {
	v.p.Call("getComponent", index)
	return v
}

// Clone TODO description.
func (v *Vector3) Clone() *Vector3 {
	v.p.Call("clone")
	return v
}

// Copy copies the src vector to v.
func (v *Vector3) Copy(src *Vector3) *Vector3 {
	v.p.Call("copy", src.p)
	return v
}

// Add adds the src vector to v.
func (v *Vector3) Add(src *Vector3) *Vector3 {
	v.p.Call("add", src.p)
	return v
}

// AddScalar adds scalar to the components of v.
func (v *Vector3) AddScalar(s float64) *Vector3 {
	v.p.Call("addScalar", s)
	return v
}

// AddVectors adds the two vectors and stores the result in v.
func (v *Vector3) AddVectors(a, b *Vector3) *Vector3 {
	v.p.Call("addVectors", a.p, b.p)
	return v
}

// AddScaledVector adds scaled vector src to v.
func (v *Vector3) AddScaledVector(src *Vector3, s float64) *Vector3 {
	v.p.Call("addScaledVector", src.p, s)
	return v
}

// Sub subtracts src from v and stores the result in v.
func (v *Vector3) Sub(src *Vector3) *Vector3 {
	v.p.Call("sub", src.p)
	return v
}

// SubScalar subtracts scalar from the components of v.
func (v *Vector3) SubScalar(s float64) *Vector3 {
	v.p.Call("subScalar", s)
	return v
}

// SubVectors subtracts the two vectors and stores the result in v.
func (v *Vector3) SubVectors(a, b *Vector3) *Vector3 {
	v.p.Call("subVectors", a.p, b.p)
	return v
}

// Multiply multiplies vector v by the src vector.
func (v *Vector3) Multiply(src *Vector3) *Vector3 {
	v.p.Call("multiply", src.p)
	return v
}

// MultiplyScalar multiplies the components of v by scalar.
func (v *Vector3) MultiplyScalar(scalar float64) *Vector3 {
	v.p.Call("multiplyScalar", scalar)
	return v
}

// MultiplyVectors multiplies a by b and stores the result in v.
func (v *Vector3) MultiplyVectors(a, b *Vector3) *Vector3 {
	v.p.Call("multiplyVectors", a.p, b.p)
	return v
}

// ApplyEuler TODO description.
func (v *Vector3) ApplyEuler() *Vector3 {
	v.p.Call("applyEuler")
	return v
}

// ApplyAxisAngle TODO description.
func (v *Vector3) ApplyAxisAngle() *Vector3 {
	v.p.Call("applyAxisAngle")
	return v
}

// ApplyMatrix3 multiplies m by v and stores the result in v.
func (v *Vector3) ApplyMatrix3(m *Matrix3) *Vector3 {
	v.p.Call("applyMatrix3", m.p)
	return v
}

// ApplyMatrix4 multiplies m by v and stores the result in v.
func (v *Vector3) ApplyMatrix4(m *Matrix4) *Vector3 {
	v.p.Call("applyMatrix4", m.p)
	return v
}

// ApplyProjection applies the projection matrix m to v
// and stores the result in v.
func (v *Vector3) ApplyProjection(m *Matrix4) *Vector3 {
	v.p.Call("applyProjection", m.p)
	return v
}

// ApplyQuaternion applies the quaternion q to v
// and stores the result in v.
func (v *Vector3) ApplyQuaternion(q *Quaternion) *Vector3 {
	v.p.Call("applyQuaternion", q.p)
	return v
}

// Project TODO description.
func (v *Vector3) Project() *Vector3 {
	v.p.Call("project")
	return v
}

// Unproject TODO description.
func (v *Vector3) Unproject() *Vector3 {
	v.p.Call("unproject")
	return v
}

// TransformDirection transforms v by affine matrix m
// and stores the results in v.
func (v *Vector3) TransformDirection(m *Matrix4) *Vector3 {
	v.p.Call("transformDirection", m.p)
	return v
}

// Divide divides vector v by the src vector.
func (v *Vector3) Divide(src *Vector3) *Vector3 {
	v.p.Call("divide", src.p)
	return v
}

// DivideScalar divides the components of v by scalar.
func (v *Vector3) DivideScalar(scalar float64) *Vector3 {
	v.p.Call("divideScalar", scalar)
	return v
}

// Min sets the components of v to the min of v and src.
func (v *Vector3) Min(src *Vector3) *Vector3 {
	v.p.Call("min", src.p)
	return v
}

// Max sets the components of v to the max of v and src.
func (v *Vector3) Max(src *Vector3) *Vector3 {
	v.p.Call("max", src.p)
	return v
}

// Clamp clamps the components of v between min and max.
//
// Note that the min components must be less than the max components
// or this function will not operate correctly.
func (v *Vector3) Clamp(min, max *Vector3) *Vector3 {
	v.p.Call("clamp", min.p, max.p)
	return v
}

// ClampScalar TODO description.
func (v *Vector3) ClampScalar() *Vector3 {
	v.p.Call("clampScalar")
	return v
}

// ClampLength clamps length of v between min and max.
func (v *Vector3) ClampLength(min, max float64) *Vector3 {
	v.p.Call("clampLength", min, max)
	return v
}

// Floor calls floor on the components of v.
func (v *Vector3) Floor() *Vector3 {
	v.p.Call("floor")
	return v
}

// Ceil calls ceil on the components of v.
func (v *Vector3) Ceil() *Vector3 {
	v.p.Call("ceil")
	return v
}

// Round rounds the components of v.
func (v *Vector3) Round() *Vector3 {
	v.p.Call("round")
	return v
}

// RoundToZero rounds the components of v to zero.
func (v *Vector3) RoundToZero() *Vector3 {
	v.p.Call("roundToZero")
	return v
}

// Negate negates the components of v.
func (v *Vector3) Negate() *Vector3 {
	v.p.Call("negate")
	return v
}

// Dot returns the dot product of (v*src)
func (v *Vector3) Dot(src *Vector3) float64 {
	return v.p.Call("dot", src.p).Float()
}

// LengthSq returns the length of v squared.
func (v *Vector3) LengthSq() float64 {
	return v.p.Call("lengthSq").Float()
}

// Length returns the length of v.
func (v *Vector3) Length() float64 {
	return v.p.Call("length").Float()
}

// LengthManhattan returns the manhattan length of v.
func (v *Vector3) LengthManhattan() float64 {
	return v.p.Call("lengthManhattan").Float()
}

// Normalize normalizes v to unit length.
func (v *Vector3) Normalize() *Vector3 {
	v.p.Call("normalize")
	return v
}

// SetLength sets the length of v.
func (v *Vector3) SetLength(length float64) *Vector3 {
	v.p.Call("setLength", length)
	return v
}

// Lerp linearly interpolates between v and src by alpha.
//
// alpha = 0 returns v unmodified.
// alpha = 1 returns v set to src.
func (v *Vector3) Lerp(src *Vector3, alpha float64) *Vector3 {
	v.p.Call("lerp", src.p, alpha)
	return v
}

// LerpVectors linearly interpolates between v1 and v2 by alpha.
//
// alpha = 0 sets v to v1.
// alpha = 1 sets v to v2.
func (v *Vector3) LerpVectors(v1, v2 *Vector3, alpha float64) *Vector3 {
	v.p.Call("lerpVectors", v1.p, v2.p, alpha)
	return v
}

// Cross TODO description.
func (v *Vector3) Cross(src *Vector3) *Vector3 {
	v.p.Call("cross", src.p)
	return v
}

// CrossVectors TODO description.
func (v *Vector3) CrossVectors(a, b *Vector3) *Vector3 {
	v.p.Call("crossVectors", a.p, b.p)
	return v
}

// ProjectOnVector TODO description.
func (v *Vector3) ProjectOnVector() *Vector3 {
	v.p.Call("projectOnVector")
	return v
}

// ProjectOnPlane TODO description.
func (v *Vector3) ProjectOnPlane() *Vector3 {
	v.p.Call("projectOnPlane")
	return v
}

// Reflect TODO description.
func (v *Vector3) Reflect() *Vector3 {
	v.p.Call("reflect")
	return v
}

// AngleTo TODO description.
func (v *Vector3) AngleTo(src *Vector3) float64 {
	return v.p.Call("angleTo", src.p).Float()
}

// DistanceTo TODO description.
func (v *Vector3) DistanceTo(src *Vector3) float64 {
	return v.p.Call("distanceTo", src.p).Float()
}

// DistanceToSquared TODO description.
func (v *Vector3) DistanceToSquared(src *Vector3) float64 {
	return v.p.Call("distanceToSquared", src.p).Float()
}

// SetFromSpherical TODO description.
func (v *Vector3) SetFromSpherical(s float64) *Vector3 {
	v.p.Call("setFromSpherical", s)
	return v
}

// SetFromMatrixPosition TODO description.
func (v *Vector3) SetFromMatrixPosition(m *Matrix3) *Vector3 {
	v.p.Call("setFromMatrixPosition", m.p)
	return v
}

// SetFromMatrixScale TODO description.
func (v *Vector3) SetFromMatrixScale(m *Matrix3) *Vector3 {
	v.p.Call("setFromMatrixScale", m.p)
	return v
}

// SetFromMatrixColumn TODO description.
func (v *Vector3) SetFromMatrixColumn(m *Matrix3, index int) *Vector3 {
	v.p.Call("setFromMatrixColumn", m.p, index)
	return v
}

// Equals compares v to src and returns true if equal.
func (v *Vector3) Equals(src *Vector3) bool {
	return v.p.Call("equals", v).Bool()
}

// FromArray sets v to the offset component of array.
func (v *Vector3) FromArray(array []float64, offset int) *Vector3 {
	v.p.Call("fromArray", array, offset)
	return v
}

// ToArray sets the offset components of array from v.
func (v *Vector3) ToArray(array []float64, offset int) *Vector3 {
	v.p.Call("toArray", array, offset)
	return v
}

// FromAttribute TODO description.
func (v *Vector3) FromAttribute(attribute, index, offset float64) *Vector3 {
	v.p.Call("fromAttribute", attribute, index, offset)
	return v
}
