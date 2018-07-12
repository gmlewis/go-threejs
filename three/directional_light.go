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

// DirectionalLight is a light that affects objects using MeshLambertMaterial or
// MeshPhongMaterial.
//
// http://threejs.org/docs/index.html#Reference/Lights/DirectionalLight
type DirectionalLight struct{ *Light }

// JSObject returns the underlying *js.Object.
func (d *DirectionalLight) JSObject() *js.Object { return d.p }

// DirectionalLight returns a DirectionalLight JavaScript class.
func (t *Three) DirectionalLight() *DirectionalLight {
	return DirectionalLightFromJSObject(t.ctx.Get("DirectionalLight"))
}

// DirectionalLightFromJSObject returns a wrapped DirectionalLight JavaScript class.
func DirectionalLightFromJSObject(p *js.Object) *DirectionalLight {
	return &DirectionalLight{&Light{&Object3D{p: p}}}
}

// NewDirectionalLight returns a new DirectionalLight object.
func (t *Three) NewDirectionalLight(color, intensity float64) *DirectionalLight {
	return DirectionalLightFromJSObject(t.ctx.Get("DirectionalLight").New(color, intensity))
}

// Copy TODO description.
func (d *DirectionalLight) Copy(source *DirectionalLight) *DirectionalLight {
	d.p.Call("copy", source.p)
	return d
}
