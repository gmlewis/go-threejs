// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Matrix3 represents a matrix3.
type Matrix3 struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (m *Matrix3) JSObject() *js.Object { return m.p }

// Matrix3 returns a Matrix3 JavaScript class.
func (t *Three) Matrix3() *Matrix3 {
	p := t.ctx.Get("Matrix3")
	return &Matrix3{p: p}
}

// NewMatrix3 returns a new Matrix3 object.
func (t *Three) NewMatrix3() *Matrix3 {
	p := t.ctx.Get("Matrix3").New()
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
func (m *Matrix3) Copy(src *Matrix3) *Matrix3 {
	m.p.Call("copy", src.p)
	return m
}

// SetFromMatrix4 TODO description.
func (m *Matrix3) SetFromMatrix4(src *Matrix4) *Matrix3 {
	m.p.Call("setFromMatrix4", src.p)
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
