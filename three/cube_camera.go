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
type CubeCamera struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *CubeCamera) JSObject() *js.Object { return t.p }

// CubeCamera returns a CubeCamera object.
func (t *Three) CubeCamera() *CubeCamera {
	p := t.ctx.Get("CubeCamera")
	return &CubeCamera{p: p}
}

// New returns a new CubeCamera object.
func (t *CubeCamera) New(near, far, cubeResolution float64) *CubeCamera {
	p := t.p.New(near, far, cubeResolution)
	return &CubeCamera{p: p}
}
