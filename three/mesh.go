// Copyright 2016 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	return MeshFromJSObject(t.ctx.Get("Mesh"))
}

// MeshFromJSObject returns a wrapped Mesh JavaScript class.
func MeshFromJSObject(p *js.Object) *Mesh {
	return &Mesh{&Object3D{p: p}}
}

// NewMesh returns a new Mesh object.
//
//     geometry — an instance of geometry.
//     material — an instance of material (optional).
func (t *Three) NewMesh(geometry, material JSObject) *Mesh {
	return MeshFromJSObject(t.ctx.Get("Mesh").New(geometry.JSObject(), material.JSObject()))
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
