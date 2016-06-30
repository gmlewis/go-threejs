// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// MeshLambertMaterial represents a material for non-shiny (Lambertian) surfaces, evaluated per vertex.
//
// http://threejs.org/docs/index.html#Reference/Materials/MeshLambertMaterial
type MeshLambertMaterial struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *MeshLambertMaterial) JSObject() *js.Object { return t.p }

// MeshLambertMaterial returns a MeshLambertMaterial object.
func (t *Three) MeshLambertMaterial() *MeshLambertMaterial {
	p := t.ctx.Get("MeshLambertMaterial")
	return &MeshLambertMaterial{p: p}
}

// New returns a new MeshLambertMaterial object.
//
// parameters is an object with one or more properties defining the material's appearance:
//     color — Line color in hexadecimal. Default is 0xffffff.
//     map — Sets the texture map. Default is null
//     lightMap — Set light map. Default is null.
//     aoMap — Set ao map. Default is null.
//     emissiveMap — Set emissive map. Default is null.
//     specularMap — Set specular map. Default is null.
//     alphaMap — Set alpha map. Default is null.
//     envMap — Set env map. Default is null.
//     fog — Define whether the material color is affected by global fog settings. Default is false.
//     wireframe — Render geometry as wireframe. Default is false (i.e. render as smooth shaded).
//     wireframeLinewidth — Controls wireframe thickness. Default is 1.
//     wireframeLinecap — Define appearance of line ends. Default is 'round'.
//     wireframeLinejoin — Define appearance of line joints. Default is 'round'.
//     vertexColors — Define how the vertices gets colored. Default is THREE.NoColors.
//     skinning — Define whether the material uses skinning. Default is false.
//     morphTargets — Define whether the material uses morphTargets. Default is false.
func (t *MeshLambertMaterial) New(parameters map[string]interface{}) *MeshLambertMaterial {
	p := t.p.New(parameters)
	return &MeshLambertMaterial{p: p}
}

// Copy TODO description.
func (m *MeshLambertMaterial) Copy(source *MeshLambertMaterial) *MeshLambertMaterial {
	m.p.Call("copy", source.p)
	return m
}
