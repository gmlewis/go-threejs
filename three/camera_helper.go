// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CameraHelper represents a camerahelper.
type CameraHelper struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (c *CameraHelper) JSObject() *js.Object { return c.p }

// CameraHelper returns a CameraHelper JavaScript class.
func (t *Three) CameraHelper() *CameraHelper {
	p := t.ctx.Get("CameraHelper")
	return CameraHelperFromJSObject(p)
}

// CameraHelperFromJSObject returns a wrapped CameraHelper JavaScript class.
func CameraHelperFromJSObject(p *js.Object) *CameraHelper {
	return &CameraHelper{p: p}
}

// NewCameraHelper returns a new CameraHelper object.
func (t *Three) NewCameraHelper(camera float64) *CameraHelper {
	p := t.ctx.Get("CameraHelper").New(camera)
	return CameraHelperFromJSObject(p)
}
