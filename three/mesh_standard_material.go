// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

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
	return &MeshStandardMaterial{p: p}
}

// NewMeshStandardMaterial returns a new MeshStandardMaterial object.
func (t *Three) NewMeshStandardMaterial(parameters float64) *MeshStandardMaterial {
	p := t.ctx.Get("MeshStandardMaterial").New(parameters)
	return &MeshStandardMaterial{p: p}
}

// Copy TODO description.
func (m *MeshStandardMaterial) Copy(source *MeshStandardMaterial) *MeshStandardMaterial {
	m.p.Call("copy", source.p)
	return m
}
