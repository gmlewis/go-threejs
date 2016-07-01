// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// SpotLight is a point light that can cast a shadow in one direction within a
// falloff cone.  It affects objects using MeshLambertMaterial or
// MeshPhongMaterial.
//
// http://threejs.org/docs/index.html#Reference/Lights/SpotLight
type SpotLight struct{ *Light }

// JSObject returns the underlying *js.Object.
func (s *SpotLight) JSObject() *js.Object { return s.p }

// SpotLight returns a SpotLight JavaScript class.
func (t *Three) SpotLight() *SpotLight {
	return spotLight(t.ctx.Get("SpotLight"))
}

// spotLight returns a wrapped SpotLight JavaScript class.
func spotLight(p *js.Object) *SpotLight {
	return &SpotLight{&Light{&Object3D{p: p}}}
}

// NewSpotLight returns a new SpotLight object.
func (t *Three) NewSpotLight(color, intensity, distance, angle, penumbra, decay float64) *SpotLight {
	return spotLight(t.ctx.Get("SpotLight").New(color, intensity, distance, angle, penumbra, decay))
}

// Get TODO description.
func (s *SpotLight) Get() *SpotLight {
	s.p.Call("get")
	return s
}

// Set TODO description.
func (s *SpotLight) Set(power float64) *SpotLight {
	s.p.Call("set", power)
	return s
}

// Copy TODO description.
func (s *SpotLight) Copy(source *SpotLight) *SpotLight {
	s.p.Call("copy", source.p)
	return s
}

// Shadow returns the property of the same name.
func (o *Object3D) Shadow() *LightShadow {
	return &LightShadow{p: o.p.Get("shadow")}
}
