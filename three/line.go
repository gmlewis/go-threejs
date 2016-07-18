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
type Line struct{ *Object3D }

// JSObject returns the underlying *js.Object.
func (l *Line) JSObject() *js.Object { return l.p }

// Line returns a Line JavaScript class.
func (t *Three) Line() *Line {
	p := t.ctx.Get("Line")
	return LineFromJSObject(p)
}

// LineFromJSObject returns a wrapper Line JavaScript class.
func LineFromJSObject(p *js.Object) *Line {
	return &Line{Object3DFromJSObject(p)}
}

// NewLine returns a new Line object.
//
//     geometry — Vertices representing the line segment(s).
//     material — Material for the line. Default is LineBasicMaterial.
func (t *Three) NewLine(geometry JSObject, material JSObject) *Line {
	p := t.ctx.Get("Line").New(geometry.JSObject(), material.JSObject())
	return LineFromJSObject(p)
}
