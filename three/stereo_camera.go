// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// StereoCamera is a stereo projection camera with left and right eyes.
type StereoCamera struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (s *StereoCamera) JSObject() *js.Object { return s.p }

// StereoCamera returns a StereoCamera JavaScript class.
func (t *Three) StereoCamera() *StereoCamera {
	p := t.ctx.Get("StereoCamera")
	return StereoCameraFromJSObject(p)
}

// StereoCameraFromJSObject returns a wrapped StereoCamera JavaScript class.
func StereoCameraFromJSObject(p *js.Object) *StereoCamera {
	return &StereoCamera{p: p}
}

// NewStereoCamera returns a new StereoCamera object.
func (t *Three) NewStereoCamera() *StereoCamera {
	p := t.ctx.Get("StereoCamera").New()
	return StereoCameraFromJSObject(p)
}

// Type returns the property of the same name.
func (s *StereoCamera) Type() string {
	return s.p.Get("type").String()
}

// Aspect returns the property of the same name.
func (s *StereoCamera) Aspect() float64 {
	return s.p.Get("aspect").Float()
}

// CameraL returns the property of the same name.
func (s *StereoCamera) CameraL() *PerspectiveCamera {
	return &PerspectiveCamera{&Camera{&Object3D{p: s.p.Get("cameraL")}}}
}

// CameraR returns the property of the same name.
func (s *StereoCamera) CameraR() *PerspectiveCamera {
	return &PerspectiveCamera{&Camera{&Object3D{p: s.p.Get("cameraR")}}}
}
