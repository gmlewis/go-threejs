package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Quaternion represents a quaternion.
type Quaternion struct{ p *js.Object }

// Quaternion returns a quaternion object.
func (t *Three) Quaternion() *Quaternion {
	p := t.ctx.Get("Quaternion")
	return &Quaternion{p: p}
}

// NewQuaternion returns a new quaternion object.
func (t *Quaternion) New(x, y, z, w float64) *Quaternion {
	p := t.p.New(x, y, z, w)
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

// Set TODO description.
func (q *Quaternion) Set(x, y, z, w float64) *Quaternion {
	q.p.Call("set", x, y, z, w)
	return q
}

// Clone TODO description.
func (q *Quaternion) Clone() *Quaternion {
	q.p.Call("clone")
	return q
}

// Copy TODO description.
func (q *Quaternion) Copy(quaternion float64) *Quaternion {
	q.p.Call("copy", quaternion)
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
func (q *Quaternion) SetFromRotationMatrix(m float64) *Quaternion {
	q.p.Call("setFromRotationMatrix", m)
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
func (q *Quaternion) Dot(v float64) *Quaternion {
	q.p.Call("dot", v)
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
func (q *Quaternion) Multiply(q, p float64) *Quaternion {
	q.p.Call("multiply", q, p)
	return q
}

// MultiplyQuaternions TODO description.
func (q *Quaternion) MultiplyQuaternions(a, b float64) *Quaternion {
	q.p.Call("multiplyQuaternions", a, b)
	return q
}

// Slerp TODO description.
func (q *Quaternion) Slerp(qb, t float64) *Quaternion {
	q.p.Call("slerp", qb, t)
	return q
}

// Equals TODO description.
func (q *Quaternion) Equals(quaternion float64) *Quaternion {
	q.p.Call("equals", quaternion)
	return q
}

// FromArray TODO description.
func (q *Quaternion) FromArray(array, offset float64) *Quaternion {
	q.p.Call("fromArray", array, offset)
	return q
}

// ToArray TODO description.
func (q *Quaternion) ToArray(array, offset float64) *Quaternion {
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
func (q *Quaternion) SlerpFlat(dst, dstOffset, src0, srcOffset0, src1, srcOffset1, t float64) *Quaternion {
	q.p.Call("slerpFlat", dst, dstOffset, src0, srcOffset0, src1, srcOffset1, t)
	return q
}
