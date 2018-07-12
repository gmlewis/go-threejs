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

// MeshStandardMaterial represents a meshstandardmaterial.
type MeshStandardMaterial struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (m *MeshStandardMaterial) JSObject() *js.Object { return m.p }

// MeshStandardMaterial returns a MeshStandardMaterial JavaScript class.
func (t *Three) MeshStandardMaterial() *MeshStandardMaterial {
	p := t.ctx.Get("MeshStandardMaterial")
	return MeshStandardMaterialFromJSObject(p)
}

// MeshStandardMaterialOpts is a map with one or more properties defining
// the material's appearance.
type MeshStandardMaterialOpts map[string]interface{}

// MeshStandardMaterialFromJSObject returns a wrapped MeshStandardMaterial JavaScript class.
func MeshStandardMaterialFromJSObject(p *js.Object) *MeshStandardMaterial {
	return &MeshStandardMaterial{p: p}
}

// NewMeshStandardMaterial returns a new MeshStandardMaterial object.
func (t *Three) NewMeshStandardMaterial(parameters MeshStandardMaterialOpts) *MeshStandardMaterial {
	p := t.ctx.Get("MeshStandardMaterial").New(parameters)
	return MeshStandardMaterialFromJSObject(p)
}

// Copy TODO description.
func (m *MeshStandardMaterial) Copy(source *MeshStandardMaterial) *MeshStandardMaterial {
	m.p.Call("copy", source.p)
	return m
}
