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

// PointLight is a light that affects objects using MeshLambertMaterial or
// MeshPhongMaterial.
//
// http://threejs.org/docs/index.html#Reference/Lights/PointLight
type PointLight struct{ *Light }

// JSObject returns the underlying *js.Object.
func (p *PointLight) JSObject() *js.Object { return p.p }

// PointLight returns a PointLight JavaScript class.
func (t *Three) PointLight() *PointLight {
	return PointLightFromJSObject(t.ctx.Get("PointLight"))
}

// PointLightFromJSObject returns a wrapped PointLight JavaScript class.
func PointLightFromJSObject(p *js.Object) *PointLight {
	return &PointLight{&Light{&Object3D{p: p}}}
}

// NewPointLight returns a new PointLight object.
func (t *Three) NewPointLight(color, intensity, distance, decay float64) *PointLight {
	return PointLightFromJSObject(t.ctx.Get("PointLight").New(color, intensity, distance, decay))
}

// Get TODO description.
func (p *PointLight) Get() *PointLight {
	p.p.Call("get")
	return p
}

// Set TODO description.
func (p *PointLight) Set(power float64) *PointLight {
	p.p.Call("set", power)
	return p
}

// Copy TODO description.
func (p *PointLight) Copy(source *PointLight) *PointLight {
	p.p.Call("copy", source.p)
	return p
}
