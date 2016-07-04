// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Matrix4 represents a matrix4.
type Matrix4 struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (m *Matrix4) JSObject() *js.Object { return m.p }

// Matrix4 returns a Matrix4 JavaScript class.
func (t *Three) Matrix4() *Matrix4 {
	p := t.ctx.Get("Matrix4")
	return matrix4(p)
}

// matrix4 returns a wrapped Matrix4 JavaScript class.
func matrix4(p *js.Object) *Matrix4 {
	return &Matrix4{p: p}
}

// NewMatrix4 returns a new Matrix4 object.
func (t *Three) NewMatrix4() *Matrix4 {
	p := t.ctx.Get("Matrix4").New()
	return matrix4(p)
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
func (m *Matrix4) Copy(src *Matrix4) *Matrix4 {
	m.p.Call("copy", src.p)
	return m
}

// CopyPosition TODO description.
func (m *Matrix4) CopyPosition(src *Matrix4) *Matrix4 {
	m.p.Call("copyPosition", src.p)
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
func (m *Matrix4) Multiply(src *Matrix4) *Matrix4 {
	m.p.Call("multiply", src.p)
	return m
}

// MultiplyMatrices TODO description.
func (m *Matrix4) MultiplyMatrices(a, b *Matrix4) *Matrix4 {
	m.p.Call("multiplyMatrices", a.p, b.p)
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
func (m *Matrix4) GetInverse(src *Matrix4, throwOnDegenerate bool) *Matrix4 {
	m.p.Call("getInverse", src.p, throwOnDegenerate)
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
func (m *Matrix4) MakeRotationAxis(axis *Vector3, angle float64) *Matrix4 {
	m.p.Call("makeRotationAxis", axis.p, angle)
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
