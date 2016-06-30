// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Mesh is the base class for mesh objects.
//
// http://threejs.org/docs/index.html#Reference/Objects/Mesh
type Mesh struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *Mesh) JSObject() *js.Object { return t.p }

// Mesh returns a Mesh object.
func (t *Three) Mesh() *Mesh {
	p := t.ctx.Get("Mesh")
	return &Mesh{p: p}
}

// New returns a new Mesh object.
//
//     geometry — an instance of geometry.
//     material — an instance of material (optional).
func (t *Mesh) New(geometry, material JSObject) *Mesh {
	p := t.p.New(geometry.JSObject(), material.JSObject())
	return &Mesh{p: p}
}

// SetDrawMode TODO description.
func (m *Mesh) SetDrawMode(value bool) *Mesh {
	m.p.Call("setDrawMode", value)
	return m
}

// GetMorphTargetIndexByName TODO description.
func (m *Mesh) GetMorphTargetIndexByName(name string) *Mesh {
	m.p.Call("getMorphTargetIndexByName", name)
	return m
}

// Object3D Properties

// ID is the Object3D property of the same name.
func (m *Mesh) ID() int {
	return m.p.Get("id").Int()
}

// Position is the Object3D property of the same name.
func (m *Mesh) Position() *Vector3 {
	return &Vector3{p: m.p.Get("position")}
}

// Rotation is the Object3D property of the same name.
func (m *Mesh) Rotation() *Euler {
	return &Euler{p: m.p.Get("rotation")}
}

// Quaternion is the Object3D property of the same name.
func (m *Mesh) Quaternion() *Quaternion {
	return &Quaternion{p: m.p.Get("quaternion")}
}

// Scale is the Object3D property of the same name.
func (m *Mesh) Scale() *Vector3 {
	return &Vector3{p: m.p.Get("scale")}
}

// ModelViewMatrix is the Object3D property of the same name.
func (m *Mesh) ModelViewMatrix() *Matrix4 {
	return &Matrix4{p: m.p.Get("modelViewMatrix")}
}

// NormalMatrix is the Object3D property of the same name.
func (m *Mesh) NormalMatrix() *Matrix3 {
	return &Matrix3{p: m.p.Get("normalMatrix")}
}
