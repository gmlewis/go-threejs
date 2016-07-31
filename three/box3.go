// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Box3 represents a box3.
type Box3 struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (b *Box3) JSObject() *js.Object { return b.p }

// Box3 returns a Box3 JavaScript class.
func (t *Three) Box3() *Box3 {
	p := t.ctx.Get("Box3")
	return &Box3{p: p}
}

// NewBox3 returns a new Box3 object.
func (t *Three) NewBox3(min, max float64) *Box3 {
	p := t.ctx.Get("Box3").New(min, max)
	return &Box3{p: p}
}

// Set sets the lower and upper (x, y) boundaries of this box.
//
//     min is lower (x, y) boundary of the box.
//     max is upper (x, y) boundary of the box.
func (b *Box3) Set(min, max float64) *Box3 {
	b.p.Call("set", min, max)
	return b
}

// SetFromArray TODO Description.
func (b *Box3) SetFromArray(array float64) *Box3 {
	b.p.Call("setFromArray", array)
	return b
}

// SetFromPoints sets the upper and lower bounds of this box to include all of the points in points.
//
//     points is the set of points that the resulting box will envelop.
func (b *Box3) SetFromPoints(points float64) *Box3 {
	b.p.Call("setFromPoints", points)
	return b
}

// SetFromCenterAndSize centers this box on center and sets this box's width and height to the values specified in size.
//
//     center is the desired center position of the box.
//     size is the desired x and y dimensions of the box.
func (b *Box3) SetFromCenterAndSize() *Box3 {
	b.p.Call("setFromCenterAndSize")
	return b
}

// SetFromObject TODO description.
func (b *Box3) SetFromObject() *Box3 {
	b.p.Call("setFromObject")
	return b
}

// Clone returns a copy of this box.
func (b *Box3) Clone() *Box3 {
	b.p.Call("clone")
	return b
}

// Copy copies the values of box to this box.
//
//     box is the box to copy.
func (b *Box3) Copy(box *Box3) *Box3 {
	b.p.Call("copy", box.p)
	return b
}

// MakeEmpty Makes this box empty.
func (b *Box3) MakeEmpty() *Box3 {
	b.p.Call("makeEmpty")
	return b
}

// IsEmpty returns true if this box includes zero points within its bounds. Note that a box with equal lower and upper bounds still includes one point, the one both bounds share.
func (b *Box3) IsEmpty() *Box3 {
	b.p.Call("isEmpty")
	return b
}

// Center returns the center point of this box.
//
//     optionalTarget, if specified, the result will be copied here.
func (b *Box3) Center(optionalTarget float64) *Box3 {
	b.p.Call("center", optionalTarget)
	return b
}

// Size returns the width and height of this box.
//
//     optionalTarget, if specified, the result will be copied here.
func (b *Box3) Size(optionalTarget float64) *Box3 {
	b.p.Call("size", optionalTarget)
	return b
}

// ExpandByPoint expands the boundaries of this box to include point.
//
//     point is the point that should be included in the box.
func (b *Box3) ExpandByPoint(point float64) *Box3 {
	b.p.Call("expandByPoint", point)
	return b
}

// ExpandByVector expands this box equilaterally by vector. The width of this box will be expanded by the x component of vector in both directions. The height of this box will be expanded by the y component of vector in both directions.
//
//     vector is the amount to expand this box in each dimension.
func (b *Box3) ExpandByVector(vector float64) *Box3 {
	b.p.Call("expandByVector", vector)
	return b
}

// ExpandByScalar expands each dimension of the box by scalar. If negative, the dimensions of the box will be contracted.
//
//     scalar is the distance to expand.
func (b *Box3) ExpandByScalar(scalar float64) *Box3 {
	b.p.Call("expandByScalar", scalar)
	return b
}

// ContainsPoint returns true if the specified point lies within the boundaries of this box.
//
//    point is Vector2 to be checked for inclusion.
func (b *Box3) ContainsPoint(point float64) *Box3 {
	b.p.Call("containsPoint", point)
	return b
}

// ContainsBox returns true if this box includes the entirety of box. If this and box overlap exactly, this function also returns true.
//
//     box to test for inclusion.
func (b *Box3) ContainsBox(box float64) *Box3 {
	b.p.Call("containsBox", box)
	return b
}

// GetParameter Returns a point as a proportion of this box's width and height.
//
//     point is Vector2
//     optionalTarget is Vector2
func (b *Box3) GetParameter(point, optionalTarget float64) *Box3 {
	b.p.Call("getParameter", point, optionalTarget)
	return b
}

// IntersectsBox Determines whether or not this box intersects box.
//
//     box to check for intersection against.
func (b *Box3) IntersectsBox(box float64) *Box3 {
	b.p.Call("intersectsBox", box)
	return b
}

// IntersectsPlane TODO description.
func (b *Box3) IntersectsPlane(plane float64) *Box3 {
	b.p.Call("intersectsPlane", plane)
	return b
}

// ClampPoint clamps point within the bounds of this box.
//
//     point is the position to clamp.
//     optionalTarget, if specified, the clamped result will be copied here.
func (b *Box3) ClampPoint(point, optionalTarget float64) *Box3 {
	b.p.Call("clampPoint", point, optionalTarget)
	return b
}

// DistanceToPoint returns the distance from any edge of this box to the specified point. If the point lies inside of this box, the distance will be 0.
//
//     point is the point to measure distance to.
func (b *Box3) DistanceToPoint() *Box3 {
	b.p.Call("distanceToPoint")
	return b
}

// GetBoundingSphere TODO description.
func (b *Box3) GetBoundingSphere() *Box3 {
	b.p.Call("getBoundingSphere")
	return b
}

// Intersect returns the intersection of this and box, setting the upper bound of this box to the lesser of the two boxes' upper bounds and the lower bound of this box to the greater of the two boxes' lower bounds.
//
//     box to intersect with.
func (b *Box3) Intersect(box float64) *Box3 {
	b.p.Call("intersect", box)
	return b
}

// Union nions this box with box setting the upper bound of this box to the greater of the two boxes' upper bounds and the lower bound of this box to the lesser of the two boxes' lower bounds.
//
//     box that will be unioned with this box.
func (b *Box3) Union(box float64) *Box3 {
	b.p.Call("union", box)
	return b
}

// ApplyMatrix4 TODO description.
func (b *Box3) ApplyMatrix4() *Box3 {
	b.p.Call("applyMatrix4")
	return b
}

// Translate adds offset to both the upper and lower bounds of this box, effectively moving this box offfset units in 2D space.
//
//     offset is the direction and distance of offset.
func (b *Box3) Translate(offset float64) *Box3 {
	b.p.Call("translate", offset)
	return b
}

// Equals returns true if this box and box share the same lower and upper bounds.
//
//     box to compare.
func (b *Box3) Equals(box float64) *Box3 {
	b.p.Call("equals", box)
	return b
}
