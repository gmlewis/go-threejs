// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Line represents a continuous line.
//
// http://threejs.org/docs/index.html#Reference/Objects/Line
type Line struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (l *Line) JSObject() *js.Object { return l.p }

// Line returns a Line object.
func (t *Three) Line() *Line {
	p := t.ctx.Get("Line")
	return &Line{p: p}
}

// New returns a new Line object.
//
//     geometry — Vertices representing the line segment(s).
//     material — Material for the line. Default is LineBasicMaterial.
func (l *Line) New(geometry []*js.Object, material *js.Object, mode float64) *Line {
	p := l.p.New(geometry, material, mode)
	return &Line{p: p}
}
