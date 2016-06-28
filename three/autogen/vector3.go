package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Vector3 represents a vector3.
type Vector3 struct{ p *js.Object }

// Vector3 returns a Vector3 object.
func (t *Three) Vector3() *Vector3 {
	p := t.ctx.Get("Vector3")
	return &Vector3{p: p}
}

// New returns a new Vector3 object.
func (t *Vector3) New(x, y, z float64) *Vector3 {
	p := t.p.New(x, y, z)
	return &Vector3{p: p}
}

// Set TODO description.
func (v *Vector3) Set(x, y, z float64) *Vector3 {
	v.p.Call("set", x, y, z)
	return v
}

// SetScalar TODO description.
func (v *Vector3) SetScalar(scalar float64) *Vector3 {
	v.p.Call("setScalar", scalar)
	return v
}

// SetX TODO description.
func (v *Vector3) SetX(x float64) *Vector3 {
	v.p.Call("setX", x)
	return v
}

// SetY TODO description.
func (v *Vector3) SetY(y float64) *Vector3 {
	v.p.Call("setY", y)
	return v
}

// SetZ TODO description.
func (v *Vector3) SetZ(z float64) *Vector3 {
	v.p.Call("setZ", z)
	return v
}

// SetComponent TODO description.
func (v *Vector3) SetComponent(index, value float64) *Vector3 {
	v.p.Call("setComponent", index, value)
	return v
}

// GetComponent TODO description.
func (v *Vector3) GetComponent(index float64) *Vector3 {
	v.p.Call("getComponent", index)
	return v
}

// Clone TODO description.
func (v *Vector3) Clone() *Vector3 {
	v.p.Call("clone")
	return v
}

// Copy TODO description.
func (v *Vector3) Copy(v float64) *Vector3 {
	v.p.Call("copy", v)
	return v
}

// Add TODO description.
func (v *Vector3) Add(v, w float64) *Vector3 {
	v.p.Call("add", v, w)
	return v
}

// AddScalar TODO description.
func (v *Vector3) AddScalar(s float64) *Vector3 {
	v.p.Call("addScalar", s)
	return v
}

// AddVectors TODO description.
func (v *Vector3) AddVectors(a, b float64) *Vector3 {
	v.p.Call("addVectors", a, b)
	return v
}

// AddScaledVector TODO description.
func (v *Vector3) AddScaledVector(v, s float64) *Vector3 {
	v.p.Call("addScaledVector", v, s)
	return v
}

// Sub TODO description.
func (v *Vector3) Sub(v, w float64) *Vector3 {
	v.p.Call("sub", v, w)
	return v
}

// SubScalar TODO description.
func (v *Vector3) SubScalar(s float64) *Vector3 {
	v.p.Call("subScalar", s)
	return v
}

// SubVectors TODO description.
func (v *Vector3) SubVectors(a, b float64) *Vector3 {
	v.p.Call("subVectors", a, b)
	return v
}

// Multiply TODO description.
func (v *Vector3) Multiply(v, w float64) *Vector3 {
	v.p.Call("multiply", v, w)
	return v
}

// MultiplyScalar TODO description.
func (v *Vector3) MultiplyScalar(scalar float64) *Vector3 {
	v.p.Call("multiplyScalar", scalar)
	return v
}

// MultiplyVectors TODO description.
func (v *Vector3) MultiplyVectors(a, b float64) *Vector3 {
	v.p.Call("multiplyVectors", a, b)
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

// ApplyMatrix3 TODO description.
func (v *Vector3) ApplyMatrix3(m float64) *Vector3 {
	v.p.Call("applyMatrix3", m)
	return v
}

// ApplyMatrix4 TODO description.
func (v *Vector3) ApplyMatrix4(m float64) *Vector3 {
	v.p.Call("applyMatrix4", m)
	return v
}

// ApplyProjection TODO description.
func (v *Vector3) ApplyProjection(m float64) *Vector3 {
	v.p.Call("applyProjection", m)
	return v
}

// ApplyQuaternion TODO description.
func (v *Vector3) ApplyQuaternion(q float64) *Vector3 {
	v.p.Call("applyQuaternion", q)
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

// TransformDirection TODO description.
func (v *Vector3) TransformDirection(m float64) *Vector3 {
	v.p.Call("transformDirection", m)
	return v
}

// Divide TODO description.
func (v *Vector3) Divide(v float64) *Vector3 {
	v.p.Call("divide", v)
	return v
}

// DivideScalar TODO description.
func (v *Vector3) DivideScalar(scalar float64) *Vector3 {
	v.p.Call("divideScalar", scalar)
	return v
}

// Min TODO description.
func (v *Vector3) Min(v float64) *Vector3 {
	v.p.Call("min", v)
	return v
}

// Max TODO description.
func (v *Vector3) Max(v float64) *Vector3 {
	v.p.Call("max", v)
	return v
}

// Clamp TODO description.
func (v *Vector3) Clamp(min, max float64) *Vector3 {
	v.p.Call("clamp", min, max)
	return v
}

// ClampScalar TODO description.
func (v *Vector3) ClampScalar() *Vector3 {
	v.p.Call("clampScalar")
	return v
}

// ClampLength TODO description.
func (v *Vector3) ClampLength(min, max float64) *Vector3 {
	v.p.Call("clampLength", min, max)
	return v
}

// Floor TODO description.
func (v *Vector3) Floor() *Vector3 {
	v.p.Call("floor")
	return v
}

// Ceil TODO description.
func (v *Vector3) Ceil() *Vector3 {
	v.p.Call("ceil")
	return v
}

// Round TODO description.
func (v *Vector3) Round() *Vector3 {
	v.p.Call("round")
	return v
}

// RoundToZero TODO description.
func (v *Vector3) RoundToZero() *Vector3 {
	v.p.Call("roundToZero")
	return v
}

// Negate TODO description.
func (v *Vector3) Negate() *Vector3 {
	v.p.Call("negate")
	return v
}

// Dot TODO description.
func (v *Vector3) Dot(v float64) *Vector3 {
	v.p.Call("dot", v)
	return v
}

// LengthSq TODO description.
func (v *Vector3) LengthSq() *Vector3 {
	v.p.Call("lengthSq")
	return v
}

// Length TODO description.
func (v *Vector3) Length() *Vector3 {
	v.p.Call("length")
	return v
}

// LengthManhattan TODO description.
func (v *Vector3) LengthManhattan() *Vector3 {
	v.p.Call("lengthManhattan")
	return v
}

// Normalize TODO description.
func (v *Vector3) Normalize() *Vector3 {
	v.p.Call("normalize")
	return v
}

// SetLength TODO description.
func (v *Vector3) SetLength(length float64) *Vector3 {
	v.p.Call("setLength", length)
	return v
}

// Lerp TODO description.
func (v *Vector3) Lerp(v, alpha float64) *Vector3 {
	v.p.Call("lerp", v, alpha)
	return v
}

// LerpVectors TODO description.
func (v *Vector3) LerpVectors(v1, v2, alpha float64) *Vector3 {
	v.p.Call("lerpVectors", v1, v2, alpha)
	return v
}

// Cross TODO description.
func (v *Vector3) Cross(v, w float64) *Vector3 {
	v.p.Call("cross", v, w)
	return v
}

// CrossVectors TODO description.
func (v *Vector3) CrossVectors(a, b float64) *Vector3 {
	v.p.Call("crossVectors", a, b)
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
func (v *Vector3) AngleTo(v float64) *Vector3 {
	v.p.Call("angleTo", v)
	return v
}

// DistanceTo TODO description.
func (v *Vector3) DistanceTo(v float64) *Vector3 {
	v.p.Call("distanceTo", v)
	return v
}

// DistanceToSquared TODO description.
func (v *Vector3) DistanceToSquared(v float64) *Vector3 {
	v.p.Call("distanceToSquared", v)
	return v
}

// SetFromSpherical TODO description.
func (v *Vector3) SetFromSpherical(s float64) *Vector3 {
	v.p.Call("setFromSpherical", s)
	return v
}

// SetFromMatrixPosition TODO description.
func (v *Vector3) SetFromMatrixPosition(m float64) *Vector3 {
	v.p.Call("setFromMatrixPosition", m)
	return v
}

// SetFromMatrixScale TODO description.
func (v *Vector3) SetFromMatrixScale(m float64) *Vector3 {
	v.p.Call("setFromMatrixScale", m)
	return v
}

// SetFromMatrixColumn TODO description.
func (v *Vector3) SetFromMatrixColumn(m, index float64) *Vector3 {
	v.p.Call("setFromMatrixColumn", m, index)
	return v
}

// Equals TODO description.
func (v *Vector3) Equals(v float64) *Vector3 {
	v.p.Call("equals", v)
	return v
}

// FromArray TODO description.
func (v *Vector3) FromArray(array, offset float64) *Vector3 {
	v.p.Call("fromArray", array, offset)
	return v
}

// ToArray TODO description.
func (v *Vector3) ToArray(array, offset float64) *Vector3 {
	v.p.Call("toArray", array, offset)
	return v
}

// FromAttribute TODO description.
func (v *Vector3) FromAttribute(attribute, index, offset float64) *Vector3 {
	v.p.Call("fromAttribute", attribute, index, offset)
	return v
}

