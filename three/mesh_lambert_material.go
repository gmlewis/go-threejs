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

// MeshLambertMaterial represents a material for non-shiny (Lambertian) surfaces, evaluated per vertex.
//
// http://threejs.org/docs/index.html#Reference/Materials/MeshLambertMaterial
type MeshLambertMaterial struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (m *MeshLambertMaterial) JSObject() *js.Object { return m.p }

// MeshLambertMaterial returns a MeshLambertMaterial JavaScript class.
func (t *Three) MeshLambertMaterial() *MeshLambertMaterial {
	p := t.ctx.Get("MeshLambertMaterial")
	return MeshLambertMaterialFromJSObject(p)
}

// MeshLambertMaterialOpts is a map with one or more properties defining
// the material's appearance:
//
//     color — Line color in hexadecimal. Default is 0xffffff.
//     map — Sets the texture map. Default is null
//     lightMap — Set light map. Default is null.
//     aoMap — Set ao map. Default is null.
//     emissiveMap — Set emissive map. Default is null.
//     specularMap — Set specular map. Default is null.
//     alphaMap — Set alpha map. Default is null.
//     envMap — Set env map. Default is null.
//     fog — Define whether the material color is affected by global fog
//         settings. Default is false.
//     wireframe — Render geometry as wireframe. Default is false (i.e. render
//         as smooth shaded).
//     wireframeLinewidth — Controls wireframe thickness. Default is 1.
//     wireframeLinecap — Define appearance of line ends. Default is 'round'.
//     wireframeLinejoin — Define appearance of line joints. Default is 'round'.
//     vertexColors — Define how the vertices gets colored.
//         Default is THREE.NoColors.
//     skinning — Define whether the material uses skinning. Default is false.
//     morphTargets — Define whether the material uses morphTargets.
//         Default is false.
type MeshLambertMaterialOpts map[string]interface{}

// MeshLambertMaterialFromJSObject returns a wrapped MeshLambertMaterial JavaScript class.
func MeshLambertMaterialFromJSObject(p *js.Object) *MeshLambertMaterial {
	return &MeshLambertMaterial{p: p}
}

// NewMeshLambertMaterial returns a new MeshLambertMaterial object.
func (t *Three) NewMeshLambertMaterial(parameters MeshLambertMaterialOpts) *MeshLambertMaterial {
	p := t.ctx.Get("MeshLambertMaterial").New(parameters)
	return MeshLambertMaterialFromJSObject(p)
}

// Copy TODO description.
func (m *MeshLambertMaterial) Copy(source *MeshLambertMaterial) *MeshLambertMaterial {
	m.p.Call("copy", source.p)
	return m
}
