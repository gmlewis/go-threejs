package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Vector4 represents a vector4.
type Vector4 struct{ p *js.Object }

// Vector4 returns a Vector4 object.
func (t *Three) Vector4() *Vector4 {
	p := t.ctx.Get("Vector4")
	return &Vector4{p: p}
}

// New returns a new Vector4 object.
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
func (v *Vector4) Copy(src *Vector4) *Vector4 {
	v.p.Call("copy", src.p)
	return v
}

// Add TODO description.
func (v *Vector4) Add(src *Vector4) *Vector4 {
	v.p.Call("add", src.p)
	return v
}

// AddScalar TODO description.
func (v *Vector4) AddScalar(s float64) *Vector4 {
	v.p.Call("addScalar", s)
	return v
}

// AddVectors TODO description.
func (v *Vector4) AddVectors(a, b *Vector4) *Vector4 {
	v.p.Call("addVectors", a.p, b.p)
	return v
}

// AddScaledVector TODO description.
func (v *Vector4) AddScaledVector(src *Vector4, s float64) *Vector4 {
	v.p.Call("addScaledVector", src.p, s)
	return v
}

// Sub TODO description.
func (v *Vector4) Sub(src *Vector4) *Vector4 {
	v.p.Call("sub", src.p)
	return v
}

// SubScalar TODO description.
func (v *Vector4) SubScalar(s float64) *Vector4 {
	v.p.Call("subScalar", s)
	return v
}

// SubVectors TODO description.
func (v *Vector4) SubVectors(a, b *Vector4) *Vector4 {
	v.p.Call("subVectors", a.p, b.p)
	return v
}

// MultiplyScalar TODO description.
func (v *Vector4) MultiplyScalar(scalar float64) *Vector4 {
	v.p.Call("multiplyScalar", scalar)
	return v
}

// ApplyMatrix4 TODO description.
func (v *Vector4) ApplyMatrix4(m *Matrix4) *Vector4 {
	v.p.Call("applyMatrix4", m.p)
	return v
}

// DivideScalar TODO description.
func (v *Vector4) DivideScalar(scalar float64) *Vector4 {
	v.p.Call("divideScalar", scalar)
	return v
}

// SetAxisAngleFromQuaternion TODO description.
func (v *Vector4) SetAxisAngleFromQuaternion(q *Quaternion) *Vector4 {
	v.p.Call("setAxisAngleFromQuaternion", q.p)
	return v
}

// SetAxisAngleFromRotationMatrix TODO description.
func (v *Vector4) SetAxisAngleFromRotationMatrix(m *Matrix4) *Vector4 {
	v.p.Call("setAxisAngleFromRotationMatrix", m.p)
	return v
}

// Min TODO description.
func (v *Vector4) Min(src *Vector4) *Vector4 {
	v.p.Call("min", src.p)
	return v
}

// Max TODO description.
func (v *Vector4) Max(src *Vector4) *Vector4 {
	v.p.Call("max", src.p)
	return v
}

// Clamp TODO description.
func (v *Vector4) Clamp(min, max *Vector4) *Vector4 {
	v.p.Call("clamp", min.p, max.p)
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
func (v *Vector4) Dot(src *Vector4) *Vector4 {
	v.p.Call("dot", src.p)
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
func (v *Vector4) Lerp(src *Vector4, alpha float64) *Vector4 {
	v.p.Call("lerp", src.p, alpha)
	return v
}

// LerpVectors TODO description.
func (v *Vector4) LerpVectors(v1, v2 *Vector4, alpha float64) *Vector4 {
	v.p.Call("lerpVectors", v1.p, v2.p, alpha)
	return v
}

// Equals TODO description.
func (v *Vector4) Equals(src *Vector4) bool {
	return v.p.Call("equals", src.p).Bool()
}

// FromArray TODO description.
func (v *Vector4) FromArray(array []float64, offset int) *Vector4 {
	v.p.Call("fromArray", array, offset)
	return v
}

/* TODO
// ToArray TODO description.
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
