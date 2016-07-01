// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// OrthographicCamera is a camera with orthographic projection.
//
// http://threejs.org/docs/index.html#Reference/Cameras/OrthographicCamera
type OrthographicCamera struct{ *Camera }

// JSObject returns the underlying *js.Object.
func (o *OrthographicCamera) JSObject() *js.Object { return o.p }

// OrthographicCamera returns an OrthographicCamera JavaScript class.
func (t *Three) OrthographicCamera() *OrthographicCamera {
	p := t.ctx.Get("OrthographicCamera")
	return &OrthographicCamera{&Camera{&Object3D{p: p}}}
}

// NewOrthographicCamera returns a new OrthographicCamera object.
func (t *Three) NewOrthographicCamera(left, right, top, bottom, near, far float64) *OrthographicCamera {
	p := t.ctx.Get("OrthographicCamera").New(left, right, top, bottom, near, far)
	return &OrthographicCamera{&Camera{&Object3D{p: p}}}
}

// Copy TODO description.
func (o *OrthographicCamera) Copy(source *OrthographicCamera) *OrthographicCamera {
	o.p.Call("copy", source.p)
	return o
}

// ToJSON TODO description.
func (o *OrthographicCamera) ToJSON(meta float64) *OrthographicCamera {
	o.p.Call("toJSON", meta)
	return o
}

// Zoom returns the property of the same name.
func (o *OrthographicCamera) Zoom() float64 {
	return o.p.Get("zoom").Float()
}

// Left returns the property of the same name.
func (o *OrthographicCamera) Left() float64 {
	return o.p.Get("left").Float()
}

// Right returns the property of the same name.
func (o *OrthographicCamera) Right() float64 {
	return o.p.Get("right").Float()
}

// Top returns the property of the same name.
func (o *OrthographicCamera) Top() float64 {
	return o.p.Get("top").Float()
}

// Bottom returns the property of the same name.
func (o *OrthographicCamera) Bottom() float64 {
	return o.p.Get("bottom").Float()
}

// Near returns the property of the same name.
func (o *OrthographicCamera) Near() float64 {
	return o.p.Get("near").Float()
}

// Far returns the property of the same name.
func (o *OrthographicCamera) Far() float64 {
	return o.p.Get("far").Float()
}
