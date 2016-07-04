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
	return camera(t.ctx.Get("Camera"))
}

// camera returns a wrapped Camera JavaScript class.
func camera(p *js.Object) *Camera {
	return &Camera{object3D(p)}
}

// NewCamera returns a new Camera object.
//
// This constructor sets the following properties to the correct
// type: matrixWorldInverse and projectionMatrix.
func (t *Three) NewCamera() *Camera {
	return camera(t.ctx.Get("Camera").New())
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

// MatrixWorldInverse is the inverse of matrixWorld. MatrixWorld
// contains the Matrix which has the world transform of the Camera.
func (c *Camera) MatrixWorldInverse() *Matrix4 {
	return &Matrix4{p: c.p.Get("matrixWorldInverse")}
}

// ProjectionMatrix is the matrix which contains the projection.
func (c *Camera) ProjectionMatrix() *Matrix4 {
	return &Matrix4{p: c.p.Get("projectionMatrix")}
}

// UpdateProjectionMatrix updates the projectionMatrix.
func (c *Camera) UpdateProjectionMatrix() *Camera {
	c.p.Call("updateProjectionMatrix")
	return c
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

// GetWorldDirection returns a vector representing the direction in
// which the camera is looking, in world space.
func (c *Camera) GetWorldDirection(vector *Vector3) *Vector3 {
	if vector != nil {
		return vector3(c.p.Call("getWorldDirection", vector.p))
	}
	return vector3(c.p.Call("getWorldDirection"))
}

// LookAt makes the camera look at the vector position in the global
// space as long as the parent of this camera is the scene or at
// position (0,0,0).
func (c *Camera) LookAt(vector *Vector3) *Camera {
	c.p.Call("lookAt", vector.p)
	return c
}

// Clone returns a clone of camera.
func (c *Camera) Clone(cam *Camera) *Camera {
	return camera(c.p.Call("clone", cam))
}
