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

// Vector4 represents a 4D vector.
//
// http://threejs.org/docs/index.html#Reference/Math/Vector4
type Vector4 struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (v *Vector4) JSObject() *js.Object { return v.p }

// Vector4 returns a Vector4 JavaScript class.
func (t *Three) Vector4() *Vector4 {
	p := t.ctx.Get("Vector4")
	return Vector4FromJSObject(p)
}

// Vector4FromJSObject returns a wrapped Vector4 JavaScript class.
func Vector4FromJSObject(p *js.Object) *Vector4 {
	return &Vector4{p: p}
}

// NewVector4 returns a new Vector4 object.
//
//     x -- Float
//     y -- Float
//     z -- Float
//     w -- Float
func (t *Three) NewVector4(x, y, z, w float64) *Vector4 {
	p := t.ctx.Get("Vector4").New(x, y, z, w)
	return Vector4FromJSObject(p)
}

// Set sets value of this vector.
func (v *Vector4) Set(x, y, z, w float64) *Vector4 {
	v.p.Call("set", x, y, z, w)
	return v
}

// SetScalar TODO description.
func (v *Vector4) SetScalar(scalar float64) *Vector4 {
	v.p.Call("setScalar", scalar)
	return v
}

// SetX sets the x component of the vector.
//
//     x -- Float
func (v *Vector4) SetX(x float64) *Vector4 {
	v.p.Call("setX", x)
	return v
}

// SetY sets the y component of the vector.
//
//     y -- Float
func (v *Vector4) SetY(y float64) *Vector4 {
	v.p.Call("setY", y)
	return v
}

// SetZ sets the z component of the vector.
//
//     z -- Float
func (v *Vector4) SetZ(z float64) *Vector4 {
	v.p.Call("setZ", z)
	return v
}

// SetW sets the w component of the vector.
//
//     w -- Float
func (v *Vector4) SetW(w float64) *Vector4 {
	v.p.Call("setW", w)
	return v
}

// SetComponent sets the value of the vector component	x, y, or z by an index.
//
//     index -- Integer 0 - 3
//         Index 0: x
//         Index 1: y
//         Index 2: z
//         Index 3: w
//     value -- Float
func (v *Vector4) SetComponent(index, value float64) *Vector4 {
	v.p.Call("setComponent", index, value)
	return v
}

// GetComponent returns the value of the vector component x, y, or z by an index.
//
//     index -- Integer 0, 1, 2, or 3
//         Index 0: x
//         Index 1: y
//         Index 2: z
//         Index 3: w
func (v *Vector4) GetComponent(index float64) *Vector4 {
	v.p.Call("getComponent", index)
	return v
}

// Clone clones this vector.
func (v *Vector4) Clone() *Vector4 {
	v.p.Call("clone")
	return v
}

// Copy copies value of v to this vector.
func (v *Vector4) Copy(src *Vector4) *Vector4 {
	v.p.Call("copy", src.p)
	return v
}

// Add adds v to this vector.
func (v *Vector4) Add(src *Vector4) *Vector4 {
	v.p.Call("add", src.p)
	return v
}

// AddScalar adds a scalar value to all of the vector's components.
//
//     s -- Float
func (v *Vector4) AddScalar(s float64) *Vector4 {
	v.p.Call("addScalar", s)
	return v
}

// AddVectors sets this vector to a + b.
func (v *Vector4) AddVectors(a, b *Vector4) *Vector4 {
	v.p.Call("addVectors", a.p, b.p)
	return v
}

// AddScaledVector adds the multiple of v and s to this vector
func (v *Vector4) AddScaledVector(src *Vector4, s float64) *Vector4 {
	v.p.Call("addScaledVector", src.p, s)
	return v
}

// Sub subtracts v from this vector.
func (v *Vector4) Sub(src *Vector4) *Vector4 {
	v.p.Call("sub", src.p)
	return v
}

// SubScalar TODO description.
func (v *Vector4) SubScalar(s float64) *Vector4 {
	v.p.Call("subScalar", s)
	return v
}

// SubVectors sets this vector to a - b.
func (v *Vector4) SubVectors(a, b *Vector4) *Vector4 {
	v.p.Call("subVectors", a.p, b.p)
	return v
}

// MultiplyScalar multiplies this vector by scalar s.
func (v *Vector4) MultiplyScalar(scalar float64) *Vector4 {
	v.p.Call("multiplyScalar", scalar)
	return v
}

// ApplyMatrix4 transforms the vector by the matrix.
//
//     m -- Matrix4
func (v *Vector4) ApplyMatrix4(m *Matrix4) *Vector4 {
	v.p.Call("applyMatrix4", m.p)
	return v
}

// DivideScalar Divides this vector by scalar s.
// Set vector to ( 0, 0, 0 ) if s == 0.
func (v *Vector4) DivideScalar(scalar float64) *Vector4 {
	v.p.Call("divideScalar", scalar)
	return v
}

// SetAxisAngleFromQuaternion sets this Vector4 to the computed axis-angle representation of the rotation
// defined by Quaternion q. The axis is stored in components (x, y, z) of the vector, and the rotation
// in radians is stored in component w.
//
//     q -- Quaternion
func (v *Vector4) SetAxisAngleFromQuaternion(q *Quaternion) *Vector4 {
	v.p.Call("setAxisAngleFromQuaternion", q.p)
	return v
}

// SetAxisAngleFromRotationMatrix Sets this Vector4 to the computed axis-angle representation of the rotation
// defined by Matrix4 m. Assumes the upper 3x3 of m is a pure rotation matrix (i.e, unscaled). The axis is
// stored in components (x, y, z) of the vector, and the rotation in radians is stored in component w.
//
//     m -- Matrix4
func (v *Vector4) SetAxisAngleFromRotationMatrix(m *Matrix4) *Vector4 {
	v.p.Call("setAxisAngleFromRotationMatrix", m.p)
	return v
}

// Min replaces the corresponing vector v value if this vector's x, y, z, or w value is greater than vector v's
// x, y, z, or w value.
//
//     v -- Vector4
func (v *Vector4) Min(src *Vector4) *Vector4 {
	v.p.Call("min", src.p)
	return v
}

// Max replaces the corresponing vector v value if this vector's x, y, z, or w value is less than vector v's
// x, y, z, or w value.
//
//     v -- Vector4
func (v *Vector4) Max(src *Vector4) *Vector4 {
	v.p.Call("max", src.p)
	return v
}

// Clamp replaces the corresponing vector value if this vector's x, y, z, or w value is greater than the max vector's
// x, y, z, or w value and replaces the corresponing vector value if this vector's x, y, z, or w value is less than
// the min vector's x, y, z, or w value.
//
//     min -- Vector4
//     max -- Vector4
func (v *Vector4) Clamp(min, max *Vector4) *Vector4 {
	v.p.Call("clamp", min.p, max.p)
	return v
}

// ClampScalar vector's x, y, z or w values are replaced by the max value if they are greater than the max value and
// if this vector's x, y, z or w values are less than the min value, they are replaced by the min value.
//
//     min -- Float the minimum value the components will be clamped to
//     max -- Float the maximum value the components will be clamped to
func (v *Vector4) ClampScalar() *Vector4 {
	v.p.Call("clampScalar")
	return v
}

// Floor components are rounded downwards (towards negative infinity) to an integer value.
func (v *Vector4) Floor() *Vector4 {
	v.p.Call("floor")
	return v
}

// Ceil components are rounded upwards (towards positive infinity) to an integer value.
func (v *Vector4) Ceil() *Vector4 {
	v.p.Call("ceil")
	return v
}

// Round components are rounded towards the nearest integer value.
func (v *Vector4) Round() *Vector4 {
	v.p.Call("round")
	return v
}

// RoundToZero components are rounded towards zero (up if negative, down if positive) to an integer value.
func (v *Vector4) RoundToZero() *Vector4 {
	v.p.Call("roundToZero")
	return v
}

// Negate inverts this vector.
func (v *Vector4) Negate() *Vector4 {
	v.p.Call("negate")
	return v
}

// Dot computes dot product of this vector and v.
func (v *Vector4) Dot(src *Vector4) *Vector4 {
	v.p.Call("dot", src.p)
	return v
}

// LengthSq computes squared length of this vector.
func (v *Vector4) LengthSq() *Vector4 {
	v.p.Call("lengthSq")
	return v
}

// Length computes length of this vector.
func (v *Vector4) Length() *Vector4 {
	v.p.Call("length")
	return v
}

// LengthManhattan computes Manhattan length of this vector.
//
// http://en.wikipedia.org/wiki/Taxicab_geometry
func (v *Vector4) LengthManhattan() *Vector4 {
	v.p.Call("lengthManhattan")
	return v
}

// Normalize normalizes this vector.
func (v *Vector4) Normalize() *Vector4 {
	v.p.Call("normalize")
	return v
}

// SetLength normalizes this vector and multiplies it by l.
func (v *Vector4) SetLength(length float64) *Vector4 {
	v.p.Call("setLength", length)
	return v
}

// Lerp linearly interpolate between this vector and v with alpha factor.
func (v *Vector4) Lerp(src *Vector4, alpha float64) *Vector4 {
	v.p.Call("lerp", src.p, alpha)
	return v
}

// LerpVectors sets this vector to be the vector linearly interpolated between v1 and v2 with alpha factor.
func (v *Vector4) LerpVectors(v1, v2 *Vector4, alpha float64) *Vector4 {
	v.p.Call("lerpVectors", v1.p, v2.p, alpha)
	return v
}

// Equals checks to see if this vector matches vector v.
//
//     v -- Vector4
func (v *Vector4) Equals(src *Vector4) bool {
	return v.p.Call("equals", src.p).Bool()
}

// FromArray sets the vector's components based on an array formatted like [x, y, z, w].
//
//     array -- The source array in the form [x, y, z, w].
//     offset -- An optional offset into the array.
func (v *Vector4) FromArray(array []float64, offset int) *Vector4 {
	v.p.Call("fromArray", array, offset)
	return v
}

/* TODO
// ToArray returns an array in the format [x, y, z, w].
//
//     array -- An optional array to store the vector.
//     offset -- An optional offset into the array.
func (v *Vector4) ToArray(array, offset float64) *Vector4 {
	v.p.Call("toArray", array, offset)
	return v
}
*/

// FromAttribute TODO description.
func (v *Vector4) FromAttribute(attribute, index, offset float64) *Vector4 {
	v.p.Call("fromAttribute", attribute, index, offset)
	return v
}

// TODO: Check Min, Max, and Clamp
