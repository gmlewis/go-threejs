// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Line3 represents a line3.
type Line3 struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *Line3) JSObject() *js.Object { return t.p }

// Line3 returns a Line3 object.
func (t *Three) Line3() *Line3 {
	p := t.ctx.Get("Line3")
	return &Line3{p: p}
}

// New returns a new Line3 object.
func (t *Line3) New(start, end float64) *Line3 {
	p := t.p.New(start, end)
	return &Line3{p: p}
}

// Set TODO description.
func (l *Line3) Set(start, end float64) *Line3 {
	l.p.Call("set", start, end)
	return l
}

// Clone TODO description.
func (l *Line3) Clone() *Line3 {
	l.p.Call("clone")
	return l
}

// Copy TODO description.
func (l *Line3) Copy(line float64) *Line3 {
	l.p.Call("copy", line)
	return l
}

// Center TODO description.
func (l *Line3) Center(optionalTarget float64) *Line3 {
	l.p.Call("center", optionalTarget)
	return l
}

// Delta TODO description.
func (l *Line3) Delta(optionalTarget float64) *Line3 {
	l.p.Call("delta", optionalTarget)
	return l
}

// DistanceSq TODO description.
func (l *Line3) DistanceSq() *Line3 {
	l.p.Call("distanceSq")
	return l
}

// Distance TODO description.
func (l *Line3) Distance() *Line3 {
	l.p.Call("distance")
	return l
}

// At TODO description.
func (l *Line3) At(t, optionalTarget float64) *Line3 {
	l.p.Call("at", t, optionalTarget)
	return l
}

// ClosestPointToPointParameter TODO description.
func (l *Line3) ClosestPointToPointParameter() *Line3 {
	l.p.Call("closestPointToPointParameter")
	return l
}

// ClosestPointToPoint TODO description.
func (l *Line3) ClosestPointToPoint(point, clampToLine, optionalTarget float64) *Line3 {
	l.p.Call("closestPointToPoint", point, clampToLine, optionalTarget)
	return l
}

// ApplyMatrix4 TODO description.
func (l *Line3) ApplyMatrix4(matrix float64) *Line3 {
	l.p.Call("applyMatrix4", matrix)
	return l
}

// Equals TODO description.
func (l *Line3) Equals(line float64) *Line3 {
	l.p.Call("equals", line)
	return l
}
