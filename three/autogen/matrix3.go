package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Matrix3 represents a matrix3.
type Matrix3 struct{ p *js.Object }

// Matrix3 returns a matrix3 object.
func (t *Three) Matrix3() *Matrix3 {
	p := t.ctx.Get("Matrix3")
	return &Matrix3{p: p}
}

// NewMatrix3 returns a new matrix3 object.
func (t *Matrix3) New() *Matrix3 {
	p := t.p.New()
	return &Matrix3{p: p}
}

// Set TODO description.
func (m *Matrix3) Set(n11, n12, n13, n21, n22, n23, n31, n32, n33 float64) *Matrix3 {
	m.p.Call("set", n11, n12, n13, n21, n22, n23, n31, n32, n33)
	return m
}

// Identity TODO description.
func (m *Matrix3) Identity() *Matrix3 {
	m.p.Call("identity")
	return m
}

// Clone TODO description.
func (m *Matrix3) Clone() *Matrix3 {
	m.p.Call("clone")
	return m
}

// Copy TODO description.
func (m *Matrix3) Copy(m float64) *Matrix3 {
	m.p.Call("copy", m)
	return m
}

// SetFromMatrix4 TODO description.
func (m *Matrix3) SetFromMatrix4(m float64) *Matrix3 {
	m.p.Call("setFromMatrix4", m)
	return m
}

// ApplyToVector3Array TODO description.
func (m *Matrix3) ApplyToVector3Array() *Matrix3 {
	m.p.Call("applyToVector3Array")
	return m
}

// ApplyToBuffer TODO description.
func (m *Matrix3) ApplyToBuffer() *Matrix3 {
	m.p.Call("applyToBuffer")
	return m
}

// MultiplyScalar TODO description.
func (m *Matrix3) MultiplyScalar(s float64) *Matrix3 {
	m.p.Call("multiplyScalar", s)
	return m
}

// Determinant TODO description.
func (m *Matrix3) Determinant() *Matrix3 {
	m.p.Call("determinant")
	return m
}

// GetInverse TODO description.
func (m *Matrix3) GetInverse(matrix, throwOnDegenerate float64) *Matrix3 {
	m.p.Call("getInverse", matrix, throwOnDegenerate)
	return m
}

// Transpose TODO description.
func (m *Matrix3) Transpose() *Matrix3 {
	m.p.Call("transpose")
	return m
}

// FlattenToArrayOffset TODO description.
func (m *Matrix3) FlattenToArrayOffset(array, offset float64) *Matrix3 {
	m.p.Call("flattenToArrayOffset", array, offset)
	return m
}

// GetNormalMatrix TODO description.
func (m *Matrix3) GetNormalMatrix(matrix4 float64) *Matrix3 {
	m.p.Call("getNormalMatrix", matrix4)
	return m
}

// TransposeIntoArray TODO description.
func (m *Matrix3) TransposeIntoArray(r float64) *Matrix3 {
	m.p.Call("transposeIntoArray", r)
	return m
}

// FromArray TODO description.
func (m *Matrix3) FromArray(array float64) *Matrix3 {
	m.p.Call("fromArray", array)
	return m
}

// ToArray TODO description.
func (m *Matrix3) ToArray() *Matrix3 {
	m.p.Call("toArray")
	return m
}
