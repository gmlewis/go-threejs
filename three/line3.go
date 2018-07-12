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

// Line3 represents a line3.
type Line3 struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (l *Line3) JSObject() *js.Object { return l.p }

// Line3 returns a Line3 JavaScript class.
func (t *Three) Line3() *Line3 {
	p := t.ctx.Get("Line3")
	return Line3FromJSObject(p)
}

// Line3FromJSObject returns a wrapped Line3 JavaScript class.
func Line3FromJSObject(p *js.Object) *Line3 {
	return &Line3{p: p}
}

// NewLine3 returns a new Line3 object.
func (t *Three) NewLine3(start, end float64) *Line3 {
	p := t.ctx.Get("Line3").New(start, end)
	return Line3FromJSObject(p)
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
func (l *Line3) Copy(line *Line3) *Line3 {
	l.p.Call("copy", line.p)
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
