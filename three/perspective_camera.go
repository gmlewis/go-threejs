// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// PerspectiveCamera is a camera with perspective projection.
//
// http://threejs.org/docs/index.html#Reference/Cameras/PerspectiveCamera
type PerspectiveCamera struct{ *Camera }

// JSObject returns the underlying *js.Object.
func (t *PerspectiveCamera) JSObject() *js.Object { return t.p }

// PerspectiveCamera returns a PerspectiveCamera object.
func (t *Three) PerspectiveCamera() *PerspectiveCamera {
	p := t.ctx.Get("PerspectiveCamera")
	return &PerspectiveCamera{&Camera{&Object3D{p: p}}}
}

// New returns a new PerspectiveCamera object.
func (t *PerspectiveCamera) New(fov, aspect, near, far float64) *PerspectiveCamera {
	p := t.p.New(fov, aspect, near, far)
	return &PerspectiveCamera{&Camera{&Object3D{p: p}}}
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
func (p *PerspectiveCamera) Copy(source float64) *PerspectiveCamera {
	p.p.Call("copy", source)
	return p
}

// ToJSON TODO description.
func (p *PerspectiveCamera) ToJSON(meta float64) *PerspectiveCamera {
	p.p.Call("toJSON", meta)
	return p
}

// Position returns the Object3D base class member of the same name.
func (p *PerspectiveCamera) Position() *Vector3 {
	return &Vector3{p: p.p.Get("position")}
}

// Rotation returns the Object3D base class member of the same name.
func (p *PerspectiveCamera) Rotation() *Euler {
	return &Euler{p: p.p.Get("rotation")}
}

// Quaternion returns the Object3D base class member of the same name.
func (p *PerspectiveCamera) Quaternion() *Quaternion {
	return &Quaternion{p: p.p.Get("quaternion")}
}

// Scale returns the Object3D base class member of the same name.
func (p *PerspectiveCamera) Scale() *Vector3 {
	return &Vector3{p: p.p.Get("scale")}
}
