// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// LineSegments represents a series of lines.
//
// http://threejs.org/docs/index.html#Reference/Objects/LineSegments
type LineSegments struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *LineSegments) JSObject() *js.Object { return t.p }

// LineSegments returns a LineSegments object.
func (t *Three) LineSegments() *LineSegments {
	p := t.ctx.Get("LineSegments")
	return &LineSegments{p: p}
}

// New returns a new LineSegments object.
//
//     geometry — Vertices representing the line segment(s).
//     material — Material for the line. Default is LineBasicMaterial.
func (t *LineSegments) New(geometry []*js.Object, material *js.Object) *LineSegments {
	p := t.p.New(geometry, material)
	return &LineSegments{p: p}
}
