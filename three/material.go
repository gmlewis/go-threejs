// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Material represents a generic material.
type Material struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (m *Material) JSObject() *js.Object { return m.p }

// Material returns a Material object.
func (t *Three) Material() *Material {
	p := t.ctx.Get("Material")
	return &Material{p: p}
}

// New returns a new Material object.
func (m *Material) New() *Material {
	p := m.p.New()
	return &Material{p: p}
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
func (m *Material) Copy(source float64) *Material {
	m.p.Call("copy", source)
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
