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
