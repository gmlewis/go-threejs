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

// Box2 represents a boundary box in 2D space.
//
// http://threejs.org/docs/index.html#Reference/Math/Box2
type Box2 struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (b *Box2) JSObject() *js.Object { return b.p }

// Box2 returns a Box2 JavaScript class.
func (t *Three) Box2() *Box2 {
	p := t.ctx.Get("Box2")
	return Box2FromJSObject(p)
}

// Box2FromJSObject returns a wrapped Box2 JavaScript class.
func Box2FromJSObject(p *js.Object) *Box2 {
	return &Box2{p: p}
}

// NewBox2 returns a new Box2 object bounded by min and max.
//
//     min -- Lower (x, y) boundary of the box.
//     max -- Upper (x, y) boundary of the box.
func (t *Three) NewBox2(min, max float64) *Box2 {
	p := t.ctx.Get("Box2").New(min, max)
	return Box2FromJSObject(p)
}

// Set sets the lower and upper (x, y) boundaries of this box.
//
//     min -- Lower (x, y) boundary of the box.
//     max -- Upper (x, y) boundary of the box.
func (b *Box2) Set(min, max float64) *Box2 {
	b.p.Call("set", min, max)
	return b
}

// SetFromPoints TODO description.
func (b *Box2) SetFromPoints(points float64) *Box2 {
	b.p.Call("setFromPoints", points)
	return b
}

// SetFromCenterAndSize TODO description.
func (b *Box2) SetFromCenterAndSize() *Box2 {
	b.p.Call("setFromCenterAndSize")
	return b
}

// Clone TODO description.
func (b *Box2) Clone() *Box2 {
	b.p.Call("clone")
	return b
}

// Copy TODO description.
func (b *Box2) Copy(box *Box2) *Box2 {
	b.p.Call("copy", box.p)
	return b
}

// MakeEmpty TODO description.
func (b *Box2) MakeEmpty() *Box2 {
	b.p.Call("makeEmpty")
	return b
}

// IsEmpty TODO description.
func (b *Box2) IsEmpty() *Box2 {
	b.p.Call("isEmpty")
	return b
}

// Center TODO description.
func (b *Box2) Center(optionalTarget float64) *Box2 {
	b.p.Call("center", optionalTarget)
	return b
}

// Size TODO description.
func (b *Box2) Size(optionalTarget float64) *Box2 {
	b.p.Call("size", optionalTarget)
	return b
}

// ExpandByPoint expands the boundaries of this box to include point.
//
//     point is a point that should be included in the box.
func (b *Box2) ExpandByPoint(point float64) *Box2 {
	b.p.Call("expandByPoint", point)
	return b
}

// ExpandByVector expands this box equilaterally by vector. The width of this box will be expanded by the x component of vector in both directions. The height of this box will be expanded by the y component of vector in both directions.
//
//     vector is the amount to expand this box in each dimension.
func (b *Box2) ExpandByVector(vector float64) *Box2 {
	b.p.Call("expandByVector", vector)
	return b
}

// ExpandByScalar expands each dimension of the box by scalar. If negative, the dimensions of the box will be contracted.
//
//     scalar is the distance to expand.
func (b *Box2) ExpandByScalar(scalar float64) *Box2 {
	b.p.Call("expandByScalar", scalar)
	return b
}

// ContainsPoint returns true if the specified point lies within the boundaries of this box.
//
//     point is a Vector2 object that is used to check for inclusion.
func (b *Box2) ContainsPoint(point float64) *Box2 {
	b.p.Call("containsPoint", point)
	return b
}

// ContainsBox returns true if this box includes the entirety of box. If this and box overlap exactly,
// this function also returns true.
//
//     box to test for inclusion.
func (b *Box2) ContainsBox(box float64) *Box2 {
	b.p.Call("containsBox", box)
	return b
}

// GetParameter returns a point as a proportion of this box's width and height.
func (b *Box2) GetParameter(point, optionalTarget *Vector2) *Box2 {
	b.p.Call("getParameter", point, optionalTarget)
	return b
}

// IntersectsBox determines whether or not this box intersects box.
//
//     box to check for intersection against.
func (b *Box2) IntersectsBox(box float64) *Box2 {
	b.p.Call("intersectsBox", box)
	return b
}

// ClampPoint clamps point within the bounds of this box.
//
//     point is the position to clamp.
//     optionalTarget, if specified, is where the clamped result will be copied.
func (b *Box2) ClampPoint(point, optionalTarget float64) *Box2 {
	b.p.Call("clampPoint", point, optionalTarget)
	return b
}

// DistanceToPoint returns the distance from any edge of this box to the specified point. If the point lies inside of this box, the distance will be 0.
//
//     Point is the point to measure distance to.
func (b *Box2) DistanceToPoint() *Box2 {
	b.p.Call("distanceToPoint")
	return b
}

// Intersect returns the intersection of this and box, setting the upper bound of this box to the lesser of the two boxes' upper bounds and the lower bound of this box to the greater of the two boxes' lower bounds.
//
//     box to intersect with.
func (b *Box2) Intersect(box float64) *Box2 {
	b.p.Call("intersect", box)
	return b
}

// Union unions this box with box setting the upper bound of this box to the greater of the two boxes' upper bounds and the lower bound of this box to the lesser of the two boxes' lower bounds.
//
//     box that will be unioned with this box.
func (b *Box2) Union(box float64) *Box2 {
	b.p.Call("union", box)
	return b
}

// Translate adds offset to both the upper and lower bounds of this box, effectively moving this box offset units in 2D space.
//
//     offset is th direction and distance of offset.
func (b *Box2) Translate(offset float64) *Box2 {
	b.p.Call("translate", offset)
	return b
}

// Equals returns true if this box and box share the same lower and upper bounds.
//
//     box to compare.
func (b *Box2) Equals(box float64) *Box2 {
	b.p.Call("equals", box)
	return b
}

// TODO: min, max
