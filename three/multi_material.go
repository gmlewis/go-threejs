// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import "github.com/gopherjs/gopherjs/js"

// MultiMaterial represents a multimaterial.
type MultiMaterial struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (m *MultiMaterial) JSObject() *js.Object { return m.p }

// MultiMaterial returns a MultiMaterial JavaScript class.
func (t *Three) MultiMaterial() *MultiMaterial {
	p := t.ctx.Get("MultiMaterial")
	return MultiMaterialFromJSObject(p)
}

// MultiMaterialFromJSObject returns a wrapped MultiMaterial JavaScript class.
func MultiMaterialFromJSObject(p *js.Object) *MultiMaterial {
	return &MultiMaterial{p: p}
}

// NewMultiMaterial returns a new MultiMaterial object.
func (t *Three) NewMultiMaterial(materials []*Material) *MultiMaterial {
	var m []*js.Object
	for i := 0; i < len(materials); i++ {
		m = append(m, materials[i].JSObject())
	}
	p := t.ctx.Get("MultiMaterial").New(m)
	return MultiMaterialFromJSObject(p)
}

// ToJSON TODO description.
func (m *MultiMaterial) ToJSON(meta float64) *MultiMaterial {
	m.p.Call("toJSON", meta)
	return m
}

// Clone TODO description.
func (m *MultiMaterial) Clone() *MultiMaterial {
	m.p.Call("clone")
	return m
}
