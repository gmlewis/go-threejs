// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Camera represents a camera.
type Camera struct{ *Object3D }

// JSObject returns the underlying *js.Object.
func (t *Camera) JSObject() *js.Object { return t.p }

// Camera returns a Camera object.
func (t *Three) Camera() *Camera {
	p := t.ctx.Get("Camera")
	return &Camera{&Object3D{p: p}}
}

// New returns a new Camera object.
func (t *Camera) New() *Camera {
	p := t.p.New()
	return &Camera{&Object3D{p: p}}
}

// Copy TODO description.
func (c *Camera) Copy(source float64) *Camera {
	c.p.Call("copy", source)
	return c
}

// Type returns the property of the same name.
func (t *Camera) Type() string {
	return t.p.Get("type").String()
}

// MatrixWorldInverse returns the property of the same name.
func (t *Camera) MatrixWorldInverse() *Matrix4 {
	return &Matrix4{p: t.p.Get("matrixWorldInverse")}
}

// ProjectionMatrix returns the property of the same name.
func (t *Camera) ProjectionMatrix() *Matrix4 {
	return &Matrix4{p: t.p.Get("projectionMatrix")}
}
