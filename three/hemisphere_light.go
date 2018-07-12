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

// HemisphereLight is a light that is positioned directly above the scene.
//
// http://threejs.org/docs/index.html#Reference/Lights/HemisphereLight
type HemisphereLight struct{ *Light }

// JSObject returns the underlying *js.Object.
func (h *HemisphereLight) JSObject() *js.Object { return h.p }

// HemisphereLight returns a HemisphereLight JavaScript class.
func (t *Three) HemisphereLight() *HemisphereLight {
	return HemisphereLightFromJSObject(t.ctx.Get("HemisphereLight"))
}

// HemisphereLightFromJSObject returns a wrapped HemisphereLight JavaScript class.
func HemisphereLightFromJSObject(p *js.Object) *HemisphereLight {
	return &HemisphereLight{&Light{&Object3D{p: p}}}
}

// NewHemisphereLight returns a new HemisphereLight object.
func (t *Three) NewHemisphereLight(skyColor, groundColor, intensity float64) *HemisphereLight {
	return HemisphereLightFromJSObject(t.ctx.Get("HemisphereLight").New(skyColor, groundColor, intensity))
}

// Copy TODO description.
func (h *HemisphereLight) Copy(source *HemisphereLight) *HemisphereLight {
	h.p.Call("copy", source.p)
	return h
}
