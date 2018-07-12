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

// CubeCamera is 6 cameras that render to a WebGLRenderTargetCube.
//
// http://threejs.org/docs/index.html#Reference/Cameras/CubeCamera
type CubeCamera struct{ *Object3D }

// JSObject returns the underlying *js.Object.
func (c *CubeCamera) JSObject() *js.Object { return c.p }

// CubeCamera returns a CubeCamera JavaScript class.
func (t *Three) CubeCamera() *CubeCamera {
	return CubeCameraFromJSObject(t.ctx.Get("CubeCamera"))
}

// CubeCameraFromJSObject returns a wrapped CubeCamera JavaScript class.
func CubeCameraFromJSObject(p *js.Object) *CubeCamera {
	return &CubeCamera{&Object3D{p: p}}
}

// NewCubeCamera returns a new CubeCamera object.
func (t *Three) NewCubeCamera(near, far, cubeResolution float64) *CubeCamera {
	return CubeCameraFromJSObject(t.ctx.Get("CubeCamera").New(near, far, cubeResolution))
}

// Type returns the property of the same name.
func (c *CubeCamera) Type() string {
	return c.p.Get("type").String()
}

// RenderTarget returns the property of the same name.
func (c *CubeCamera) RenderTarget() *WebGLRenderTargetCube {
	return &WebGLRenderTargetCube{p: c.p.Get("renderTarget")}
}

// UpdateCubeMap TODO description.
func (c *CubeCamera) UpdateCubeMap(renderer, scene JSObject) *CubeCamera {
	c.p.Call("updateCubeMap", renderer.JSObject(), scene.JSObject())
	return c
}
