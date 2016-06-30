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
type OrthographicCamera struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *OrthographicCamera) JSObject() *js.Object { return t.p }

// OrthographicCamera returns an OrthographicCamera object.
func (t *Three) OrthographicCamera() *OrthographicCamera {
	p := t.ctx.Get("OrthographicCamera")
	return &OrthographicCamera{p: p}
}

// New returns a new OrthographicCamera object.
func (t *OrthographicCamera) New(left, right, top, bottom, near, far float64) *OrthographicCamera {
	p := t.p.New(left, right, top, bottom, near, far)
	return &OrthographicCamera{p: p}
}

// Copy TODO description.
func (o *OrthographicCamera) Copy(source float64) *OrthographicCamera {
	o.p.Call("copy", source)
	return o
}

// ToJSON TODO description.
func (o *OrthographicCamera) ToJSON(meta float64) *OrthographicCamera {
	o.p.Call("toJSON", meta)
	return o
}
