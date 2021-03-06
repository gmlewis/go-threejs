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

// PerspectiveCamera is a camera with perspective projection.
//
// http://threejs.org/docs/index.html#Reference/Cameras/PerspectiveCamera
type PerspectiveCamera struct{ *Camera }

// JSObject returns the underlying *js.Object.
func (p *PerspectiveCamera) JSObject() *js.Object { return p.p }

// PerspectiveCamera returns a PerspectiveCamera JavaScript class.
func (t *Three) PerspectiveCamera() *PerspectiveCamera {
	return PerspectiveCameraFromJSObject(t.ctx.Get("PerspectiveCamera"))
}

// PerspectiveCameraFromJSObject returns a wrapped PerspectiveCamera JavaScript class.
func PerspectiveCameraFromJSObject(p *js.Object) *PerspectiveCamera {
	return &PerspectiveCamera{&Camera{&Object3D{p: p}}}
}

// NewPerspectiveCamera returns a new PerspectiveCamera object.
func (t *Three) NewPerspectiveCamera(fov, aspect, near, far float64) *PerspectiveCamera {
	return PerspectiveCameraFromJSObject(t.ctx.Get("PerspectiveCamera").New(fov, aspect, near, far))
}

// SetLens TODO description.
func (p *PerspectiveCamera) SetLens(focalLength, frameHeight float64) *PerspectiveCamera {
	p.p.Call("setLens", focalLength, frameHeight)
	return p
}

// SetViewOffset TODO description.
func (p *PerspectiveCamera) SetViewOffset(fullWidth, fullHeight, x, y, width, height float64) *PerspectiveCamera {
	p.p.Call("setViewOffset", fullWidth, fullHeight, x, y, width, height)
	return p
}

// Copy TODO description.
func (p *PerspectiveCamera) Copy(source *PerspectiveCamera) *PerspectiveCamera {
	p.p.Call("copy", source.p)
	return p
}

// ToJSON TODO description.
func (p *PerspectiveCamera) ToJSON(meta float64) *PerspectiveCamera {
	p.p.Call("toJSON", meta)
	return p
}

// Zoom returns the property of the same name.
func (p *PerspectiveCamera) Zoom() float64 {
	return p.p.Get("zoom").Float()
}

// FocalLength returns the property of the same name.
func (p *PerspectiveCamera) FocalLength() float64 {
	return p.p.Get("focalLength").Float()
}

// FOV returns the property of the same name.
func (p *PerspectiveCamera) FOV() float64 {
	return p.p.Get("fov").Float()
}

// Aspect returns the property of the same name.
func (p *PerspectiveCamera) Aspect() float64 {
	return p.p.Get("aspect").Float()
}

// SetAspect set the aspect property.
func (p *PerspectiveCamera) SetAspect(value float64) *PerspectiveCamera {
	p.p.Set("aspect", value)
	return p
}

// Near returns the property of the same name.
func (p *PerspectiveCamera) Near() float64 {
	return p.p.Get("near").Float()
}

// Far returns the property of the same name.
func (p *PerspectiveCamera) Far() float64 {
	return p.p.Get("far").Float()
}
