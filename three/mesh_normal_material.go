// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// MeshNormalMaterial represents a material that maps the normal vectors to RGB colors.
//
// http://threejs.org/docs/index.html#Reference/Materials/MeshNormalMaterial
type MeshNormalMaterial struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *MeshNormalMaterial) JSObject() *js.Object { return t.p }

// MeshNormalMaterial returns a MeshNormalMaterial object.
func (t *Three) MeshNormalMaterial() *MeshNormalMaterial {
	p := t.ctx.Get("MeshNormalMaterial")
	return &MeshNormalMaterial{p: p}
}

// New returns a new MeshNormalMaterial object.
//
// parameters is an object with one or more properties defining the material's appearance:
//     wireframe -- Render geometry as wireframe. Default is false (i.e. render as smooth shaded).
//     wireframeLinewidth -- Controls wireframe thickness. Default is 1.
//     morphTargets -- Define whether the material uses morphTargets. Default is false.
func (t *MeshNormalMaterial) New(parameters map[string]interface{}) *MeshNormalMaterial {
	p := t.p.New(parameters)
	return &MeshNormalMaterial{p: p}
}

// Copy TODO description.
func (m *MeshNormalMaterial) Copy(source *MeshNormalMaterial) *MeshNormalMaterial {
	m.p.Call("copy", source.p)
	return m
}
