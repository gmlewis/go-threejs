// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

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
	return &MeshBasicMaterial{p: p}
}

type MeshBasicMaterialOpts map[string]interface{}

// NewMeshBasicMaterial returns a new MeshBasicMaterial object.
//
// parameters is an object with one or more properties defining the material's appearance:
//     color — geometry color in hexadecimal. Default is 0xffffff.
//     map — Set texture map. Default is null
//     aoMap — Set ambient occlusion map. Default is null
//     specularMap — Set specular map. Default is null.
//     alphaMap — Set alpha map. Default is null.
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
func (t *Three) NewMeshBasicMaterial(parameters MeshBasicMaterialOpts) *MeshBasicMaterial {
	p := t.ctx.Get("MeshBasicMaterial").New(parameters)
	return &MeshBasicMaterial{p: p}
}

// Copy TODO description.
func (m *MeshBasicMaterial) Copy(source *MeshBasicMaterial) *MeshBasicMaterial {
	m.p.Call("copy", source.p)
	return m
}
