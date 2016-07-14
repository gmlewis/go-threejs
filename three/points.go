// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Points is a class for displaying particles in the form of variable size points.
//
// http://threejs.org/docs/index.html#Reference/Objects/Points
type Points struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (p *Points) JSObject() *js.Object { return p.p }

// Points returns a Points JavaScript class.
func (t *Three) Points() *Points {
	p := t.ctx.Get("Points")
	return PointsFromJSObject(p)
}

// PointsFromJSObject returns a wrapped Points JavaScript class.
func PointsFromJSObject(p *js.Object) *Points {
	return &Points{p: p}
}

// NewPoints returns a new Points object.
//
//     geometry — an instance of geometry.
//     material — an instance of material (optional).
func (t *Three) NewPoints(geometry, material *js.Object) *Points {
	p := t.ctx.Get("Points").New(geometry, material)
	return PointsFromJSObject(p)
}
