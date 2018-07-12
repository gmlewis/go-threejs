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

// MeshBasicMaterial represents a material for drawing geometries in a simple shaded (flat or wireframe) way.
//
// http://threejs.org/docs/index.html#Reference/Materials/MeshBasicMaterial
type MeshBasicMaterial struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (m *MeshBasicMaterial) JSObject() *js.Object { return m.p }

// MeshBasicMaterial returns a MeshBasicMaterial JavaScript class.
func (t *Three) MeshBasicMaterial() *MeshBasicMaterial {
	p := t.ctx.Get("MeshBasicMaterial")
	return MeshBasicMaterialFromJSObject(p)
}

// MeshBasicMaterialOpts is a map with one or more properties defining the
// material's appearance:
//
//     color — geometry color in hexadecimal. Default is 0xffffff.
//     map — Set texture map. Default is null
//     aoMap — Set ambient occlusion map. Default is null
//     specularMap — Set specular map. Default is null.
//     alphaMap — Set alpha map. Default is null.
//     envMap — Set env map. Default is null.
//     fog — Define whether the material color is affected by global fog
//         settings. Default is true.
//     shading — Define shading type. Default is THREE.SmoothShading.
//     wireframe — render geometry as wireframe. Default is false.
//     wireframeLinewidth — Line thickness. Default is 1.
//     wireframeLinecap — Define appearance of line ends. Default is 'round'.
//     wireframeLinejoin — Define appearance of line joints. Default is 'round'.
//     vertexColors — Define how the vertices gets colored.
//         Default is THREE.NoColors.
//     skinning — Define whether the material uses skinning. Default is false.
//     morphTargets — Define whether the material uses morphTargets.
//         Default is false.
type MeshBasicMaterialOpts map[string]interface{}

// MeshBasicMaterialFromJSObject returns a wrapped MeshBasicMaterial JavaScript class.
func MeshBasicMaterialFromJSObject(p *js.Object) *MeshBasicMaterial {
	return &MeshBasicMaterial{p: p}
}

// NewMeshBasicMaterial returns a new MeshBasicMaterial object.
func (t *Three) NewMeshBasicMaterial(parameters MeshBasicMaterialOpts) *MeshBasicMaterial {
	p := t.ctx.Get("MeshBasicMaterial").New(parameters)
	return MeshBasicMaterialFromJSObject(p)
}

// Copy TODO description.
func (m *MeshBasicMaterial) Copy(source *MeshBasicMaterial) *MeshBasicMaterial {
	m.p.Call("copy", source.p)
	return m
}
