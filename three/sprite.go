// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Sprite represents a plane in an 3d scene which faces always towards the camera.
//
// http://threejs.org/docs/index.html#Reference/Objects/Sprite
type Sprite struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *Sprite) JSObject() *js.Object { return t.p }

// Sprite returns a Sprite object.
func (t *Three) Sprite() *Sprite {
	p := t.ctx.Get("Sprite")
	return &Sprite{p: p}
}

// New returns a new Sprite object.
//
//     material — An instance of Material (optional).
func (t *Sprite) New(material *js.Object) *Sprite {
	p := t.p.New(material)
	return &Sprite{p: p}
}
