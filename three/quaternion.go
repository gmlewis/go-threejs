// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Quaternion represents a quaternion.
type Quaternion struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (q *Quaternion) JSObject() *js.Object { return q.p }

// Quaternion returns a Quaternion object.
func (t *Three) Quaternion() *Quaternion {
	p := t.ctx.Get("Quaternion")
	return &Quaternion{p: p}
}

// New returns a new Quaternion object.
func (q *Quaternion) New(x, y, z, w float64) *Quaternion {
	p := q.p.New(x, y, z, w)
	return &Quaternion{p: p}
}

// X returns the x-component of the Quaternion.
func (q *Quaternion) X() float64 {
	return q.p.Get("x").Float()
}

// Y returns the y-component of the Quaternion.
func (q *Quaternion) Y() float64 {
	return q.p.Get("y").Float()
}

// Z returns the z-component of the Quaternion.
func (q *Quaternion) Z() float64 {
	return q.p.Get("z").Float()
}

// W returns the w-component of the Quaternion.
func (q *Quaternion) W() float64 {
	return q.p.Get("w").Float()
}

// SetX sets the x-component of the Quaternion.
func (q *Quaternion) SetX(value float64) *Quaternion {
	q.p.Set("x", value)
	return q
}

// SetY sets the y-component of the Quaternion.
func (q *Quaternion) SetY(value float64) *Quaternion {
	q.p.Set("y", value)
	return q
}

// SetZ sets the z-component of the Quaternion.
func (q *Quaternion) SetZ(value float64) *Quaternion {
	q.p.Set("z", value)
	return q
}

// SetW sets the w-component of the Quaternion.
func (q *Quaternion) SetW(value float64) *Quaternion {
	q.p.Set("w", value)
	return q
}

// Set sets the components of the quaternion.
func (q *Quaternion) Set(x, y, z, w float64) *Quaternion {
	q.p.Call("set", x, y, z, w)
	return q
}

// Clone returns a clone of this quaternion.
func (q *Quaternion) Clone() *Quaternion {
	q.p.Call("clone")
	return q
}

// Copy copies the values from quaternion into q.
func (q *Quaternion) Copy(quaternion *Quaternion) *Quaternion {
	q.p.Call("copy", quaternion.p)
	return q
}

// SetFromEuler TODO description.
func (q *Quaternion) SetFromEuler(euler, update float64) *Quaternion {
	q.p.Call("setFromEuler", euler, update)
	return q
}

// SetFromAxisAngle TODO description.
func (q *Quaternion) SetFromAxisAngle(axis, angle float64) *Quaternion {
	q.p.Call("setFromAxisAngle", axis, angle)
	return q
}

// SetFromRotationMatrix TODO description.
func (q *Quaternion) SetFromRotationMatrix(m *Matrix3) *Quaternion {
	q.p.Call("setFromRotationMatrix", m.p)
	return q
}

// SetFromUnitVectors TODO description.
func (q *Quaternion) SetFromUnitVectors() *Quaternion {
	q.p.Call("setFromUnitVectors")
	return q
}

// Inverse TODO description.
func (q *Quaternion) Inverse() *Quaternion {
	q.p.Call("inverse")
	return q
}

// Conjugate TODO description.
func (q *Quaternion) Conjugate() *Quaternion {
	q.p.Call("conjugate")
	return q
}

// Dot TODO description.
func (q *Quaternion) Dot(src *Quaternion) *Quaternion {
	q.p.Call("dot", src.p)
	return q
}

// LengthSq TODO description.
func (q *Quaternion) LengthSq() *Quaternion {
	q.p.Call("lengthSq")
	return q
}

// Length TODO description.
func (q *Quaternion) Length() *Quaternion {
	q.p.Call("length")
	return q
}

// Normalize TODO description.
func (q *Quaternion) Normalize() *Quaternion {
	q.p.Call("normalize")
	return q
}

// Multiply TODO description.
func (q *Quaternion) Multiply(src *Quaternion) *Quaternion {
	q.p.Call("multiply", src.p)
	return q
}

// MultiplyQuaternions TODO description.
func (q *Quaternion) MultiplyQuaternions(a, b *Quaternion) *Quaternion {
	q.p.Call("multiplyQuaternions", a.p, b.p)
	return q
}

// Slerp TODO description.
func (q *Quaternion) Slerp(qb *Quaternion, t float64) *Quaternion {
	q.p.Call("slerp", qb.p, t)
	return q
}

// Equals TODO description.
func (q *Quaternion) Equals(quaternion *Quaternion) bool {
	return q.p.Call("equals", quaternion.p).Bool()
}

// FromArray TODO description.
func (q *Quaternion) FromArray(array []float64, offset int) *Quaternion {
	q.p.Call("fromArray", array, offset)
	return q
}

// ToArray TODO description.
func (q *Quaternion) ToArray(array []float64, offset int) *Quaternion {
	q.p.Call("toArray", array, offset)
	return q
}

// OnChange TODO description.
func (q *Quaternion) OnChange(callback float64) *Quaternion {
	q.p.Call("onChange", callback)
	return q
}

// OnChangeCallback TODO description.
func (q *Quaternion) OnChangeCallback() *Quaternion {
	q.p.Call("onChangeCallback")
	return q
}

// SlerpFlat TODO description.
func (q *Quaternion) SlerpFlat(dsq *Quaternion, dstOffset int, src0 *Quaternion, srcOffset0 int, src1 *Quaternion, srcOffset1 int, t float64) *Quaternion {
	q.p.Call("slerpFlat", dsq.p, dstOffset, src0.p, srcOffset0, src1.p, srcOffset1, t)
	return q
}
