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
func (t *CameraHelper) JSObject() *js.Object { return t.p }

// CameraHelper returns a CameraHelper object.
func (t *Three) CameraHelper() *CameraHelper {
	p := t.ctx.Get("CameraHelper")
	return &CameraHelper{p: p}
}

// New returns a new CameraHelper object.
func (t *CameraHelper) New(camera float64) *CameraHelper {
	p := t.p.New(camera)
	return &CameraHelper{p: p}
}
