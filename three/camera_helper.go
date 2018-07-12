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

// CameraHelper is an Object3D which helps visualizing what
// a camera contains in its frustum.
// It visualizes the frustum with a Line geometry.
//
// http://threejs.org/docs/index.html#Reference/Extras.Helpers/CameraHelper
type CameraHelper struct{ *Line }

// JSObject returns the underlying *js.Object.
func (c *CameraHelper) JSObject() *js.Object { return c.p }

// CameraHelper returns a CameraHelper JavaScript class.
func (t *Three) CameraHelper() *CameraHelper {
	p := t.ctx.Get("CameraHelper")
	return CameraHelperFromJSObject(p)
}

// CameraHelperFromJSObject returns a wrapped CameraHelper JavaScript class.
func CameraHelperFromJSObject(p *js.Object) *CameraHelper {
	return &CameraHelper{LineFromJSObject(p)}
}

// NewCameraHelper returns a new CameraHelper object.
func (t *Three) NewCameraHelper(camera float64) *CameraHelper {
	p := t.ctx.Get("CameraHelper").New(camera)
	return CameraHelperFromJSObject(p)
}
