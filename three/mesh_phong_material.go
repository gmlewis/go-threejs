// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// MeshPhongMaterial represents a material for shiny surfaces, evaluated per pixel.
//
// http://threejs.org/docs/index.html#Reference/Materials/MeshPhongMaterial
type MeshPhongMaterial struct{ *Material }

// JSObject returns the underlying *js.Object.
func (m *MeshPhongMaterial) JSObject() *js.Object { return m.p }

// MeshPhongMaterial returns a MeshPhongMaterial JavaScript class.
func (t *Three) MeshPhongMaterial() *MeshPhongMaterial {
	p := t.ctx.Get("MeshPhongMaterial")
	return meshPhongMaterial(p)
}

// meshPhongMaterial returns a wrapped MeshPhongMaterial JavaScript class.
func meshPhongMaterial(p *js.Object) *MeshPhongMaterial {
	return &MeshPhongMaterial{material(p)}
}

// MeshPhongMaterialOpts is a map with one or more properties defining
// the material's appearance:
//
//     color — geometry color in hexadecimal. Default is 0xffffff.
//     map — Set texture map. Default is null
//     lightMap — Set light map. Default is null.
//     aoMap — Set ao map. Default is null.
//     emissiveMap — Set emissive map. Default is null.
//     specularMap — Set specular map. Default is null.
//     alphaMap — Set alpha map. Default is null.
//     displacementMap — Set displacement map. Default is null.
//     displacementScale — Set displacement scale. Default is 1.
//     displacementBias — Set displacement offset. Default is 0.
//     envMap — Set env map. Default is null.
//     fog — Define whether the material color is affected by global fog settings. Default is true.
//     shading — Define shading type. Default is THREE.SmoothShading.
//     wireframe — render geometry as wireframe. Default is false.
//     wireframeLinewidth — Line thickness. Default is 1.
//     wireframeLinecap — Define appearance of line ends. Default is 'round'.
//     wireframeLinejoin — Define appearance of line joints. Default is 'round'.
//     vertexColors — Define how the vertices gets colored. Default is THREE.NoColors.
//     skinning — Define whether the material uses skinning. Default is false.
//     morphTargets — Define whether the material uses morphTargets. Default is false.
type MeshPhongMaterialOpts map[string]interface{}

// NewMeshPhongMaterial returns a new MeshPhongMaterial object.
func (t *Three) NewMeshPhongMaterial(parameters MeshPhongMaterialOpts) *MeshPhongMaterial {
	p := t.ctx.Get("MeshPhongMaterial").New(parameters)
	return meshPhongMaterial(p)
}

// Copy TODO description.
func (m *MeshPhongMaterial) Copy(source *MeshPhongMaterial) *MeshPhongMaterial {
	m.p.Call("copy", source.p)
	return m
}
