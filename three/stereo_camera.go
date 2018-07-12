// Copyright 2016 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
