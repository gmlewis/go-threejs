package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Vector2 represents a vector2.
type Vector2 struct{ p *js.Object }

// Vector2 returns a Vector2 object.
func (t *Three) Vector2() *Vector2 {
	p := t.ctx.Get("Vector2")
	return &Vector2{p: p}
}

// New returns a new Vector2 object.
func (t *Vector2) New(x, y float64) *Vector2 {
	p := t.p.New(x, y)
	return &Vector2{p: p}
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

// Set TODO description.
func (v *Vector2) Set(x, y float64) *Vector2 {
	v.p.Call("set", x, y)
	return v
}

// SetScalar TODO description.
func (v *Vector2) SetScalar(scalar float64) *Vector2 {
	v.p.Call("setScalar", scalar)
	return v
}

// SetX TODO description.
func (v *Vector2) SetX(x float64) *Vector2 {
	v.p.Call("setX", x)
	return v
}

// SetY TODO description.
func (v *Vector2) SetY(y float64) *Vector2 {
	v.p.Call("setY", y)
	return v
}

// SetComponent TODO description.
func (v *Vector2) SetComponent(index, value float64) *Vector2 {
	v.p.Call("setComponent", index, value)
	return v
}

// GetComponent TODO description.
func (v *Vector2) GetComponent(index float64) *Vector2 {
	v.p.Call("getComponent", index)
	return v
}

// Clone TODO description.
func (v *Vector2) Clone() *Vector2 {
	v.p.Call("clone")
	return v
}

// Copy TODO description.
func (v *Vector2) Copy(v float64) *Vector2 {
	v.p.Call("copy", v)
	return v
}

// Add TODO description.
func (v *Vector2) Add(v, w float64) *Vector2 {
	v.p.Call("add", v, w)
	return v
}

// AddScalar TODO description.
func (v *Vector2) AddScalar(s float64) *Vector2 {
	v.p.Call("addScalar", s)
	return v
}

// AddVectors TODO description.
func (v *Vector2) AddVectors(a, b float64) *Vector2 {
	v.p.Call("addVectors", a, b)
	return v
}

// AddScaledVector TODO description.
func (v *Vector2) AddScaledVector(v, s float64) *Vector2 {
	v.p.Call("addScaledVector", v, s)
	return v
}

// Sub TODO description.
func (v *Vector2) Sub(v, w float64) *Vector2 {
	v.p.Call("sub", v, w)
	return v
}

// SubScalar TODO description.
func (v *Vector2) SubScalar(s float64) *Vector2 {
	v.p.Call("subScalar", s)
	return v
}

// SubVectors TODO description.
func (v *Vector2) SubVectors(a, b float64) *Vector2 {
	v.p.Call("subVectors", a, b)
	return v
}

// Multiply TODO description.
func (v *Vector2) Multiply(v float64) *Vector2 {
	v.p.Call("multiply", v)
	return v
}

// MultiplyScalar TODO description.
func (v *Vector2) MultiplyScalar(scalar float64) *Vector2 {
	v.p.Call("multiplyScalar", scalar)
	return v
}

// Divide TODO description.
func (v *Vector2) Divide(v float64) *Vector2 {
	v.p.Call("divide", v)
	return v
}

// DivideScalar TODO description.
func (v *Vector2) DivideScalar(scalar float64) *Vector2 {
	v.p.Call("divideScalar", scalar)
	return v
}

// Min TODO description.
func (v *Vector2) Min(v float64) *Vector2 {
	v.p.Call("min", v)
	return v
}

// Max TODO description.
func (v *Vector2) Max(v float64) *Vector2 {
	v.p.Call("max", v)
	return v
}

// Clamp TODO description.
func (v *Vector2) Clamp(min, max float64) *Vector2 {
	v.p.Call("clamp", min, max)
	return v
}

// ClampScalar TODO description.
func (v *Vector2) ClampScalar() *Vector2 {
	v.p.Call("clampScalar")
	return v
}

// ClampLength TODO description.
func (v *Vector2) ClampLength(min, max float64) *Vector2 {
	v.p.Call("clampLength", min, max)
	return v
}

// Floor TODO description.
func (v *Vector2) Floor() *Vector2 {
	v.p.Call("floor")
	return v
}

// Ceil TODO description.
func (v *Vector2) Ceil() *Vector2 {
	v.p.Call("ceil")
	return v
}

// Round TODO description.
func (v *Vector2) Round() *Vector2 {
	v.p.Call("round")
	return v
}

// RoundToZero TODO description.
func (v *Vector2) RoundToZero() *Vector2 {
	v.p.Call("roundToZero")
	return v
}

// Negate TODO description.
func (v *Vector2) Negate() *Vector2 {
	v.p.Call("negate")
	return v
}

// Dot TODO description.
func (v *Vector2) Dot(v float64) *Vector2 {
	v.p.Call("dot", v)
	return v
}

// LengthSq TODO description.
func (v *Vector2) LengthSq() *Vector2 {
	v.p.Call("lengthSq")
	return v
}

// Length TODO description.
func (v *Vector2) Length() *Vector2 {
	v.p.Call("length")
	return v
}

// LengthManhattan TODO description.
func (v *Vector2) LengthManhattan() *Vector2 {
	v.p.Call("lengthManhattan")
	return v
}

// Normalize TODO description.
func (v *Vector2) Normalize() *Vector2 {
	v.p.Call("normalize")
	return v
}

// Angle TODO description.
func (v *Vector2) Angle() *Vector2 {
	v.p.Call("angle")
	return v
}

// DistanceTo TODO description.
func (v *Vector2) DistanceTo(v float64) *Vector2 {
	v.p.Call("distanceTo", v)
	return v
}

// DistanceToSquared TODO description.
func (v *Vector2) DistanceToSquared(v float64) *Vector2 {
	v.p.Call("distanceToSquared", v)
	return v
}

// SetLength TODO description.
func (v *Vector2) SetLength(length float64) *Vector2 {
	v.p.Call("setLength", length)
	return v
}

// Lerp TODO description.
func (v *Vector2) Lerp(v, alpha float64) *Vector2 {
	v.p.Call("lerp", v, alpha)
	return v
}

// LerpVectors TODO description.
func (v *Vector2) LerpVectors(v1, v2, alpha float64) *Vector2 {
	v.p.Call("lerpVectors", v1, v2, alpha)
	return v
}

// Equals TODO description.
func (v *Vector2) Equals(v float64) *Vector2 {
	v.p.Call("equals", v)
	return v
}

// FromArray TODO description.
func (v *Vector2) FromArray(array, offset float64) *Vector2 {
	v.p.Call("fromArray", array, offset)
	return v
}

// ToArray TODO description.
func (v *Vector2) ToArray(array, offset float64) *Vector2 {
	v.p.Call("toArray", array, offset)
	return v
}

// FromAttribute TODO description.
func (v *Vector2) FromAttribute(attribute, index, offset float64) *Vector2 {
	v.p.Call("fromAttribute", attribute, index, offset)
	return v
}

// RotateAround TODO description.
func (v *Vector2) RotateAround(center, angle float64) *Vector2 {
	v.p.Call("rotateAround", center, angle)
	return v
}

