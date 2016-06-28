package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Matrix4 represents a matrix4.
type Matrix4 struct{ p *js.Object }

// Matrix4 returns a Matrix4 object.
func (t *Three) Matrix4() *Matrix4 {
	p := t.ctx.Get("Matrix4")
	return &Matrix4{p: p}
}

// New returns a new Matrix4 object.
func (t *Matrix4) New() *Matrix4 {
	p := t.p.New()
	return &Matrix4{p: p}
}

// Set TODO description.
func (m *Matrix4) Set(n11, n12, n13, n14, n21, n22, n23, n24, n31, n32, n33, n34, n41, n42, n43, n44 float64) *Matrix4 {
	m.p.Call("set", n11, n12, n13, n14, n21, n22, n23, n24, n31, n32, n33, n34, n41, n42, n43, n44)
	return m
}

// Identity TODO description.
func (m *Matrix4) Identity() *Matrix4 {
	m.p.Call("identity")
	return m
}

// Clone TODO description.
func (m *Matrix4) Clone() *Matrix4 {
	m.p.Call("clone")
	return m
}

// Copy TODO description.
func (m *Matrix4) Copy(m float64) *Matrix4 {
	m.p.Call("copy", m)
	return m
}

// CopyPosition TODO description.
func (m *Matrix4) CopyPosition(m float64) *Matrix4 {
	m.p.Call("copyPosition", m)
	return m
}

// ExtractBasis TODO description.
func (m *Matrix4) ExtractBasis(xAxis, yAxis, zAxis float64) *Matrix4 {
	m.p.Call("extractBasis", xAxis, yAxis, zAxis)
	return m
}

// MakeBasis TODO description.
func (m *Matrix4) MakeBasis(xAxis, yAxis, zAxis float64) *Matrix4 {
	m.p.Call("makeBasis", xAxis, yAxis, zAxis)
	return m
}

// ExtractRotation TODO description.
func (m *Matrix4) ExtractRotation() *Matrix4 {
	m.p.Call("extractRotation")
	return m
}

// MakeRotationFromEuler TODO description.
func (m *Matrix4) MakeRotationFromEuler(euler float64) *Matrix4 {
	m.p.Call("makeRotationFromEuler", euler)
	return m
}

// MakeRotationFromQuaternion TODO description.
func (m *Matrix4) MakeRotationFromQuaternion(q float64) *Matrix4 {
	m.p.Call("makeRotationFromQuaternion", q)
	return m
}

// LookAt TODO description.
func (m *Matrix4) LookAt() *Matrix4 {
	m.p.Call("lookAt")
	return m
}

// Multiply TODO description.
func (m *Matrix4) Multiply(m, n float64) *Matrix4 {
	m.p.Call("multiply", m, n)
	return m
}

// MultiplyMatrices TODO description.
func (m *Matrix4) MultiplyMatrices(a, b float64) *Matrix4 {
	m.p.Call("multiplyMatrices", a, b)
	return m
}

// MultiplyToArray TODO description.
func (m *Matrix4) MultiplyToArray(a, b, r float64) *Matrix4 {
	m.p.Call("multiplyToArray", a, b, r)
	return m
}

// MultiplyScalar TODO description.
func (m *Matrix4) MultiplyScalar(s float64) *Matrix4 {
	m.p.Call("multiplyScalar", s)
	return m
}

// ApplyToVector3Array TODO description.
func (m *Matrix4) ApplyToVector3Array() *Matrix4 {
	m.p.Call("applyToVector3Array")
	return m
}

// ApplyToBuffer TODO description.
func (m *Matrix4) ApplyToBuffer() *Matrix4 {
	m.p.Call("applyToBuffer")
	return m
}

// Determinant TODO description.
func (m *Matrix4) Determinant() *Matrix4 {
	m.p.Call("determinant")
	return m
}

// Transpose TODO description.
func (m *Matrix4) Transpose() *Matrix4 {
	m.p.Call("transpose")
	return m
}

// FlattenToArrayOffset TODO description.
func (m *Matrix4) FlattenToArrayOffset(array, offset float64) *Matrix4 {
	m.p.Call("flattenToArrayOffset", array, offset)
	return m
}

// GetPosition TODO description.
func (m *Matrix4) GetPosition() *Matrix4 {
	m.p.Call("getPosition")
	return m
}

// SetPosition TODO description.
func (m *Matrix4) SetPosition(v float64) *Matrix4 {
	m.p.Call("setPosition", v)
	return m
}

// GetInverse TODO description.
func (m *Matrix4) GetInverse(m, throwOnDegenerate float64) *Matrix4 {
	m.p.Call("getInverse", m, throwOnDegenerate)
	return m
}

// Scale TODO description.
func (m *Matrix4) Scale(v float64) *Matrix4 {
	m.p.Call("scale", v)
	return m
}

// GetMaxScaleOnAxis TODO description.
func (m *Matrix4) GetMaxScaleOnAxis() *Matrix4 {
	m.p.Call("getMaxScaleOnAxis")
	return m
}

// MakeTranslation TODO description.
func (m *Matrix4) MakeTranslation(x, y, z float64) *Matrix4 {
	m.p.Call("makeTranslation", x, y, z)
	return m
}

// MakeRotationX TODO description.
func (m *Matrix4) MakeRotationX(theta float64) *Matrix4 {
	m.p.Call("makeRotationX", theta)
	return m
}

// MakeRotationY TODO description.
func (m *Matrix4) MakeRotationY(theta float64) *Matrix4 {
	m.p.Call("makeRotationY", theta)
	return m
}

// MakeRotationZ TODO description.
func (m *Matrix4) MakeRotationZ(theta float64) *Matrix4 {
	m.p.Call("makeRotationZ", theta)
	return m
}

// MakeRotationAxis TODO description.
func (m *Matrix4) MakeRotationAxis(axis, angle float64) *Matrix4 {
	m.p.Call("makeRotationAxis", axis, angle)
	return m
}

// MakeScale TODO description.
func (m *Matrix4) MakeScale(x, y, z float64) *Matrix4 {
	m.p.Call("makeScale", x, y, z)
	return m
}

// Compose TODO description.
func (m *Matrix4) Compose(position, quaternion, scale float64) *Matrix4 {
	m.p.Call("compose", position, quaternion, scale)
	return m
}

// Decompose TODO description.
func (m *Matrix4) Decompose() *Matrix4 {
	m.p.Call("decompose")
	return m
}

// MakeFrustum TODO description.
func (m *Matrix4) MakeFrustum(left, right, bottom, top, near, far float64) *Matrix4 {
	m.p.Call("makeFrustum", left, right, bottom, top, near, far)
	return m
}

// MakePerspective TODO description.
func (m *Matrix4) MakePerspective(fov, aspect, near, far float64) *Matrix4 {
	m.p.Call("makePerspective", fov, aspect, near, far)
	return m
}

// MakeOrthographic TODO description.
func (m *Matrix4) MakeOrthographic(left, right, top, bottom, near, far float64) *Matrix4 {
	m.p.Call("makeOrthographic", left, right, top, bottom, near, far)
	return m
}

// Equals TODO description.
func (m *Matrix4) Equals(matrix float64) *Matrix4 {
	m.p.Call("equals", matrix)
	return m
}

// FromArray TODO description.
func (m *Matrix4) FromArray(array float64) *Matrix4 {
	m.p.Call("fromArray", array)
	return m
}

// ToArray TODO description.
func (m *Matrix4) ToArray() *Matrix4 {
	m.p.Call("toArray")
	return m
}

