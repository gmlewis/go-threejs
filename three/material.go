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

// Material represents a generic material.
type Material struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (m *Material) JSObject() *js.Object { return m.p }

// Material returns a Material JavaScript class.
func (t *Three) Material() *Material {
	p := t.ctx.Get("Material")
	return MaterialFromJSObject(p)
}

// MaterialFromJSObject returns a wrapped Material JavaScript class.
func MaterialFromJSObject(p *js.Object) *Material {
	return &Material{p: p}
}

// NewMaterial returns a new Material object.
func (t *Three) NewMaterial() *Material {
	p := t.ctx.Get("Material").New()
	return MaterialFromJSObject(p)
}

// NeedsUpdate returns the needsUpdate-component of the Material.
func (m *Material) NeedsUpdate() float64 {
	return m.p.Get("needsUpdate").Float()
}

// SetNeedsUpdate sets the needsUpdate-component of the Material.
func (m *Material) SetNeedsUpdate(value float64) *Material {
	m.p.Set("needsUpdate", value)
	return m
}

// SetValues TODO description.
func (m *Material) SetValues(values float64) *Material {
	m.p.Call("setValues", values)
	return m
}

// ToJSON TODO description.
func (m *Material) ToJSON(meta float64) *Material {
	m.p.Call("toJSON", meta)
	return m
}

// Clone TODO description.
func (m *Material) Clone() *Material {
	m.p.Call("clone")
	return m
}

// Copy TODO description.
func (m *Material) Copy(source *Material) *Material {
	m.p.Call("copy", source.p)
	return m
}

// Update TODO description.
func (m *Material) Update() *Material {
	m.p.Call("update")
	return m
}

// Dispose TODO description.
func (m *Material) Dispose() *Material {
	m.p.Call("dispose")
	return m
}

// SetShading set the shading-component of the Material.
func (m *Material) SetShading(value int) *Material {
	m.p.Set("shading", value)
	return m
}
