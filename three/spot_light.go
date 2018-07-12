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
	return SpotLightFromJSObject(t.ctx.Get("SpotLight"))
}

// SpotLightFromJSObject returns a wrapped SpotLight JavaScript class.
func SpotLightFromJSObject(p *js.Object) *SpotLight {
	return &SpotLight{&Light{&Object3D{p: p}}}
}

// NewSpotLight returns a new SpotLight object.
func (t *Three) NewSpotLight(color, intensity, distance, angle, penumbra, decay float64) *SpotLight {
	return SpotLightFromJSObject(t.ctx.Get("SpotLight").New(color, intensity, distance, angle, penumbra, decay))
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
