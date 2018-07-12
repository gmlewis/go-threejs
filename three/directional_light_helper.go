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

// DirectionalLightHelper visualizes a DirectionalLight's effect on the scene.
//
// http://threejs.org/docs/index.html#Reference/Extras.Helpers/DirectionalLightHelper
type DirectionalLightHelper struct{ *Object3D }

// JSObject returns the underlying *js.Object.
func (d *DirectionalLightHelper) JSObject() *js.Object { return d.p }

// DirectionalLightHelper returns a DirectionalLightHelper JavaScript class.
func (t *Three) DirectionalLightHelper() *DirectionalLightHelper {
	p := t.ctx.Get("DirectionalLightHelper")
	return DirectionalLightHelperFromJSObject(p)
}

// DirectionalLightHelperFromJSObject returns a wrapped DirectionalLightHelper JavaScript class.
func DirectionalLightHelperFromJSObject(p *js.Object) *DirectionalLightHelper {
	return &DirectionalLightHelper{Object3DFromJSObject(p)}
}

// NewDirectionalLightHelper returns a new DirectionalLightHelper object.
func (t *Three) NewDirectionalLightHelper(light, size float64) *DirectionalLightHelper {
	p := t.ctx.Get("DirectionalLightHelper").New(light, size)
	return DirectionalLightHelperFromJSObject(p)
}
