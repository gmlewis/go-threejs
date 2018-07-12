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

// MeshNormalMaterial represents a material that maps the normal vectors to RGB colors.
//
// http://threejs.org/docs/index.html#Reference/Materials/MeshNormalMaterial
type MeshNormalMaterial struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (m *MeshNormalMaterial) JSObject() *js.Object { return m.p }

// MeshNormalMaterial returns a MeshNormalMaterial JavaScript class.
func (t *Three) MeshNormalMaterial() *MeshNormalMaterial {
	p := t.ctx.Get("MeshNormalMaterial")
	return MeshNormalMaterialFromJSObject(p)
}

// MeshNormalMaterialOpts is a map with one or more properties defining
// the material's appearance:
//
//     wireframe -- Render geometry as wireframe. Default is false (i.e. render as smooth shaded).
//     wireframeLinewidth -- Controls wireframe thickness. Default is 1.
//     morphTargets -- Define whether the material uses morphTargets. Default is false.
type MeshNormalMaterialOpts map[string]interface{}

// MeshNormalMaterialFromJSObject returns a wrapped MeshNormalMaterial JavaScript class.
func MeshNormalMaterialFromJSObject(p *js.Object) *MeshNormalMaterial {
	return &MeshNormalMaterial{p: p}
}

// NewMeshNormalMaterial returns a new MeshNormalMaterial object.
func (t *Three) NewMeshNormalMaterial(parameters MeshNormalMaterialOpts) *MeshNormalMaterial {
	p := t.ctx.Get("MeshNormalMaterial").New(parameters)
	return MeshNormalMaterialFromJSObject(p)
}

// Copy TODO description.
func (m *MeshNormalMaterial) Copy(source *MeshNormalMaterial) *MeshNormalMaterial {
	m.p.Call("copy", source.p)
	return m
}
