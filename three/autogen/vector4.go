package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Vector4 represents a vector4.
type Vector4 struct{ p *js.Object }

// Vector4 returns a vector4 object.
func (t *Three) Vector4() *Vector4 {
	p := t.ctx.Get("Vector4")
	return &Vector4{p: p}
}

// NewVector4 returns a new vector4 object.
func (t *Vector4) New(x, y, z, w float64) *Vector4 {
	p := t.p.New(x, y, z, w)
	return &Vector4{p: p}
}

// Set TODO description.
func (v *Vector4) Set(x, y, z, w float64) *Vector4 {
	v.p.Call("set", x, y, z, w)
	return v
}

// SetScalar TODO description.
func (v *Vector4) SetScalar(scalar float64) *Vector4 {
	v.p.Call("setScalar", scalar)
	return v
}

// SetX TODO description.
func (v *Vector4) SetX(x float64) *Vector4 {
	v.p.Call("setX", x)
	return v
}

// SetY TODO description.
func (v *Vector4) SetY(y float64) *Vector4 {
	v.p.Call("setY", y)
	return v
}

// SetZ TODO description.
func (v *Vector4) SetZ(z float64) *Vector4 {
	v.p.Call("setZ", z)
	return v
}

// SetW TODO description.
func (v *Vector4) SetW(w float64) *Vector4 {
	v.p.Call("setW", w)
	return v
}

// SetComponent TODO description.
func (v *Vector4) SetComponent(index, value float64) *Vector4 {
	v.p.Call("setComponent", index, value)
	return v
}

// GetComponent TODO description.
func (v *Vector4) GetComponent(index float64) *Vector4 {
	v.p.Call("getComponent", index)
	return v
}

// Clone TODO description.
func (v *Vector4) Clone() *Vector4 {
	v.p.Call("clone")
	return v
}

// Copy TODO description.
func (v *Vector4) Copy(v float64) *Vector4 {
	v.p.Call("copy", v)
	return v
}

// Add TODO description.
func (v *Vector4) Add(v, w float64) *Vector4 {
	v.p.Call("add", v, w)
	return v
}

// AddScalar TODO description.
func (v *Vector4) AddScalar(s float64) *Vector4 {
	v.p.Call("addScalar", s)
	return v
}

// AddVectors TODO description.
func (v *Vector4) AddVectors(a, b float64) *Vector4 {
	v.p.Call("addVectors", a, b)
	return v
}

// AddScaledVector TODO description.
func (v *Vector4) AddScaledVector(v, s float64) *Vector4 {
	v.p.Call("addScaledVector", v, s)
	return v
}

// Sub TODO description.
func (v *Vector4) Sub(v, w float64) *Vector4 {
	v.p.Call("sub", v, w)
	return v
}

// SubScalar TODO description.
func (v *Vector4) SubScalar(s float64) *Vector4 {
	v.p.Call("subScalar", s)
	return v
}

// SubVectors TODO description.
func (v *Vector4) SubVectors(a, b float64) *Vector4 {
	v.p.Call("subVectors", a, b)
	return v
}

// MultiplyScalar TODO description.
func (v *Vector4) MultiplyScalar(scalar float64) *Vector4 {
	v.p.Call("multiplyScalar", scalar)
	return v
}

// ApplyMatrix4 TODO description.
func (v *Vector4) ApplyMatrix4(m float64) *Vector4 {
	v.p.Call("applyMatrix4", m)
	return v
}

// DivideScalar TODO description.
func (v *Vector4) DivideScalar(scalar float64) *Vector4 {
	v.p.Call("divideScalar", scalar)
	return v
}

// SetAxisAngleFromQuaternion TODO description.
func (v *Vector4) SetAxisAngleFromQuaternion(q float64) *Vector4 {
	v.p.Call("setAxisAngleFromQuaternion", q)
	return v
}

// SetAxisAngleFromRotationMatrix TODO description.
func (v *Vector4) SetAxisAngleFromRotationMatrix(m float64) *Vector4 {
	v.p.Call("setAxisAngleFromRotationMatrix", m)
	return v
}

// Min TODO description.
func (v *Vector4) Min(v float64) *Vector4 {
	v.p.Call("min", v)
	return v
}

// Max TODO description.
func (v *Vector4) Max(v float64) *Vector4 {
	v.p.Call("max", v)
	return v
}

// Clamp TODO description.
func (v *Vector4) Clamp(min, max float64) *Vector4 {
	v.p.Call("clamp", min, max)
	return v
}

// ClampScalar TODO description.
func (v *Vector4) ClampScalar() *Vector4 {
	v.p.Call("clampScalar")
	return v
}

// Floor TODO description.
func (v *Vector4) Floor() *Vector4 {
	v.p.Call("floor")
	return v
}

// Ceil TODO description.
func (v *Vector4) Ceil() *Vector4 {
	v.p.Call("ceil")
	return v
}

// Round TODO description.
func (v *Vector4) Round() *Vector4 {
	v.p.Call("round")
	return v
}

// RoundToZero TODO description.
func (v *Vector4) RoundToZero() *Vector4 {
	v.p.Call("roundToZero")
	return v
}

// Negate TODO description.
func (v *Vector4) Negate() *Vector4 {
	v.p.Call("negate")
	return v
}

// Dot TODO description.
func (v *Vector4) Dot(v float64) *Vector4 {
	v.p.Call("dot", v)
	return v
}

// LengthSq TODO description.
func (v *Vector4) LengthSq() *Vector4 {
	v.p.Call("lengthSq")
	return v
}

// Length TODO description.
func (v *Vector4) Length() *Vector4 {
	v.p.Call("length")
	return v
}

// LengthManhattan TODO description.
func (v *Vector4) LengthManhattan() *Vector4 {
	v.p.Call("lengthManhattan")
	return v
}

// Normalize TODO description.
func (v *Vector4) Normalize() *Vector4 {
	v.p.Call("normalize")
	return v
}

// SetLength TODO description.
func (v *Vector4) SetLength(length float64) *Vector4 {
	v.p.Call("setLength", length)
	return v
}

// Lerp TODO description.
func (v *Vector4) Lerp(v, alpha float64) *Vector4 {
	v.p.Call("lerp", v, alpha)
	return v
}

// LerpVectors TODO description.
func (v *Vector4) LerpVectors(v1, v2, alpha float64) *Vector4 {
	v.p.Call("lerpVectors", v1, v2, alpha)
	return v
}

// Equals TODO description.
func (v *Vector4) Equals(v float64) *Vector4 {
	v.p.Call("equals", v)
	return v
}

// FromArray TODO description.
func (v *Vector4) FromArray(array, offset float64) *Vector4 {
	v.p.Call("fromArray", array, offset)
	return v
}

// ToArray TODO description.
func (v *Vector4) ToArray(array, offset float64) *Vector4 {
	v.p.Call("toArray", array, offset)
	return v
}

// FromAttribute TODO description.
func (v *Vector4) FromAttribute(attribute, index, offset float64) *Vector4 {
	v.p.Call("fromAttribute", attribute, index, offset)
	return v
}

