// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Camera is the JavaScript base class for cameras.
// It should always be inherited when you build a new camera.
//
// http://threejs.org/docs/index.html#Reference/Cameras/Camera
type Camera struct{ *Object3D }

// JSObject returns the underlying *js.Object.
func (c *Camera) JSObject() *js.Object { return c.p }

// Camera returns a Camera JavaScript class.
func (t *Three) Camera() *Camera {
	p := t.ctx.Get("Camera")
	return &Camera{&Object3D{p: p}}
}

// NewCamera returns a new Camera object.
func (t *Three) NewCamera() *Camera {
	p := t.ctx.Get("Camera").New()
	return &Camera{&Object3D{p: p}}
}

// Copy TODO description.
func (c *Camera) Copy(source *Camera) *Camera {
	c.p.Call("copy", source.p)
	return c
}

// Type returns the property of the same name.
func (c *Camera) Type() string {
	return c.p.Get("type").String()
}

// MatrixWorldInverse returns the property of the same name.
func (c *Camera) MatrixWorldInverse() *Matrix4 {
	return &Matrix4{p: c.p.Get("matrixWorldInverse")}
}

// ProjectionMatrix returns the property of the same name.
func (c *Camera) ProjectionMatrix() *Matrix4 {
	return &Matrix4{p: c.p.Get("projectionMatrix")}
}

// SetNear sets the property of the same name.
// Note that this function was added to make the SpotLight example work.
func (c *Camera) SetNear(value float64) *Camera {
	c.p.Set("near", value)
	return c
}

// SetFar sets the property of the same name.
// Note that this function was added to make the SpotLight example work.
func (c *Camera) SetFar(value float64) *Camera {
	c.p.Set("far", value)
	return c
}

// SetFOV sets the property of the same name.
// Note that this function was added to make the SpotLight example work.
func (c *Camera) SetFOV(value float64) *Camera {
	c.p.Set("fov", value)
	return c
}
