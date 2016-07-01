// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

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
	p := t.ctx.Get("CubeCamera")
	return &CubeCamera{&Object3D{p: p}}
}

// NewCubeCamera returns a new CubeCamera object.
func (t *Three) NewCubeCamera(near, far, cubeResolution float64) *CubeCamera {
	p := t.ctx.Get("CubeCamera").New(near, far, cubeResolution)
	return &CubeCamera{&Object3D{p: p}}
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
