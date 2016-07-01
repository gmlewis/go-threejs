// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// DirectionalLight is a light that affects objects using MeshLambertMaterial or
// MeshPhongMaterial.
//
// http://threejs.org/docs/index.html#Reference/Lights/DirectionalLight
type DirectionalLight struct{ *Light }

// JSObject returns the underlying *js.Object.
func (d *DirectionalLight) JSObject() *js.Object { return d.p }

// DirectionalLight returns a DirectionalLight JavaScript class.
func (t *Three) DirectionalLight() *DirectionalLight {
	return directionalLight(t.ctx.Get("DirectionalLight"))
}

// directionalLight returns a wrapped DirectionalLight JavaScript class.
func directionalLight(p *js.Object) *DirectionalLight {
	return &DirectionalLight{&Light{&Object3D{p: p}}}
}

// NewDirectionalLight returns a new DirectionalLight object.
func (t *Three) NewDirectionalLight(color, intensity float64) *DirectionalLight {
	return directionalLight(t.ctx.Get("DirectionalLight").New(color, intensity))
}

// Copy TODO description.
func (d *DirectionalLight) Copy(source *DirectionalLight) *DirectionalLight {
	d.p.Call("copy", source.p)
	return d
}
