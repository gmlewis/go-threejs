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
type Mesh struct{ *Object3D }

// JSObject returns the underlying *js.Object.
func (m *Mesh) JSObject() *js.Object { return m.p }

// Mesh returns a Mesh JavaScript class.
func (t *Three) Mesh() *Mesh {
	p := t.ctx.Get("Mesh")
	return &Mesh{&Object3D{p: p}}
}

// NewMesh returns a new Mesh object.
//
//     geometry — an instance of geometry.
//     material — an instance of material (optional).
func (t *Three) NewMesh(geometry, material JSObject) *Mesh {
	p := t.ctx.Get("Mesh").New(geometry.JSObject(), material.JSObject())
	return &Mesh{&Object3D{p: p}}
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
