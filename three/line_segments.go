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
func (l *LineSegments) JSObject() *js.Object { return l.p }

// LineSegments returns a LineSegments JavaScript class.
func (t *Three) LineSegments() *LineSegments {
	p := t.ctx.Get("LineSegments")
	return &LineSegments{p: p}
}

// NewLineSegments returns a new LineSegments object.
//
//     geometry — Vertices representing the line segment(s).
//     material — Material for the line. Default is LineBasicMaterial.
func (t *Three) NewLineSegments(geometry []*js.Object, material *js.Object) *LineSegments {
	p := t.ctx.Get("LineSegments").New(geometry, material)
	return &LineSegments{p: p}
}
