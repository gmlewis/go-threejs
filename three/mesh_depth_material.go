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

// MeshDepthMaterial represents a material for drawing geometry by depth.
// Depth is based off of the camera near and far plane. White is nearest, black is farthest.
//
// http://threejs.org/docs/index.html#Reference/Materials/MeshDepthMaterial
type MeshDepthMaterial struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (m *MeshDepthMaterial) JSObject() *js.Object { return m.p }

// MeshDepthMaterial returns a MeshDepthMaterial JavaScript class.
func (t *Three) MeshDepthMaterial() *MeshDepthMaterial {
	p := t.ctx.Get("MeshDepthMaterial")
	return MeshDepthMaterialFromJSObject(p)
}

// MeshDepthMaterialOpts is a map with one or more properties defining
// the material's appearance:
//
//     morphTargets -- Define whether the material uses morphTargets. Default is false.
//     wireframe -- Render geometry as wireframe. Default is false (i.e. render as smooth shaded).
//     wireframeLinewidth -- Controls wireframe thickness. Default is 1.
type MeshDepthMaterialOpts map[string]interface{}

// MeshDepthMaterialFromJSObject returns a wrapped MeshDepthMaterial JavaScript class.
func MeshDepthMaterialFromJSObject(p *js.Object) *MeshDepthMaterial {
	return &MeshDepthMaterial{p: p}
}

// NewMeshDepthMaterial returns a new MeshDepthMaterial object.
func (t *Three) NewMeshDepthMaterial(parameters MeshDepthMaterialOpts) *MeshDepthMaterial {
	p := t.ctx.Get("MeshDepthMaterial").New(parameters)
	return MeshDepthMaterialFromJSObject(p)
}

// Copy TODO description.
func (m *MeshDepthMaterial) Copy(source *MeshDepthMaterial) *MeshDepthMaterial {
	m.p.Call("copy", source.p)
	return m
}
