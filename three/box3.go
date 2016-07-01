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

// Box3 returns a Box3 object.
func (t *Three) Box3() *Box3 {
	p := t.ctx.Get("Box3")
	return &Box3{p: p}
}

// New returns a new Box3 object.
func (b *Box3) New(min, max float64) *Box3 {
	p := b.p.New(min, max)
	return &Box3{p: p}
}

// Set TODO description.
func (b *Box3) Set(min, max float64) *Box3 {
	b.p.Call("set", min, max)
	return b
}

// SetFromArray TODO description.
func (b *Box3) SetFromArray(array float64) *Box3 {
	b.p.Call("setFromArray", array)
	return b
}

// SetFromPoints TODO description.
func (b *Box3) SetFromPoints(points float64) *Box3 {
	b.p.Call("setFromPoints", points)
	return b
}

// SetFromCenterAndSize TODO description.
func (b *Box3) SetFromCenterAndSize() *Box3 {
	b.p.Call("setFromCenterAndSize")
	return b
}

// SetFromObject TODO description.
func (b *Box3) SetFromObject() *Box3 {
	b.p.Call("setFromObject")
	return b
}

// Clone TODO description.
func (b *Box3) Clone() *Box3 {
	b.p.Call("clone")
	return b
}

// Copy TODO description.
func (b *Box3) Copy(box *Box3) *Box3 {
	b.p.Call("copy", box.p)
	return b
}

// MakeEmpty TODO description.
func (b *Box3) MakeEmpty() *Box3 {
	b.p.Call("makeEmpty")
	return b
}

// IsEmpty TODO description.
func (b *Box3) IsEmpty() *Box3 {
	b.p.Call("isEmpty")
	return b
}

// Center TODO description.
func (b *Box3) Center(optionalTarget float64) *Box3 {
	b.p.Call("center", optionalTarget)
	return b
}

// Size TODO description.
func (b *Box3) Size(optionalTarget float64) *Box3 {
	b.p.Call("size", optionalTarget)
	return b
}

// ExpandByPoint TODO description.
func (b *Box3) ExpandByPoint(point float64) *Box3 {
	b.p.Call("expandByPoint", point)
	return b
}

// ExpandByVector TODO description.
func (b *Box3) ExpandByVector(vector float64) *Box3 {
	b.p.Call("expandByVector", vector)
	return b
}

// ExpandByScalar TODO description.
func (b *Box3) ExpandByScalar(scalar float64) *Box3 {
	b.p.Call("expandByScalar", scalar)
	return b
}

// ContainsPoint TODO description.
func (b *Box3) ContainsPoint(point float64) *Box3 {
	b.p.Call("containsPoint", point)
	return b
}

// ContainsBox TODO description.
func (b *Box3) ContainsBox(box float64) *Box3 {
	b.p.Call("containsBox", box)
	return b
}

// GetParameter TODO description.
func (b *Box3) GetParameter(point, optionalTarget float64) *Box3 {
	b.p.Call("getParameter", point, optionalTarget)
	return b
}

// IntersectsBox TODO description.
func (b *Box3) IntersectsBox(box float64) *Box3 {
	b.p.Call("intersectsBox", box)
	return b
}

// IntersectsPlane TODO description.
func (b *Box3) IntersectsPlane(plane float64) *Box3 {
	b.p.Call("intersectsPlane", plane)
	return b
}

// ClampPoint TODO description.
func (b *Box3) ClampPoint(point, optionalTarget float64) *Box3 {
	b.p.Call("clampPoint", point, optionalTarget)
	return b
}

// DistanceToPoint TODO description.
func (b *Box3) DistanceToPoint() *Box3 {
	b.p.Call("distanceToPoint")
	return b
}

// GetBoundingSphere TODO description.
func (b *Box3) GetBoundingSphere() *Box3 {
	b.p.Call("getBoundingSphere")
	return b
}

// Intersect TODO description.
func (b *Box3) Intersect(box float64) *Box3 {
	b.p.Call("intersect", box)
	return b
}

// Union TODO description.
func (b *Box3) Union(box float64) *Box3 {
	b.p.Call("union", box)
	return b
}

// ApplyMatrix4 TODO description.
func (b *Box3) ApplyMatrix4() *Box3 {
	b.p.Call("applyMatrix4")
	return b
}

// Translate TODO description.
func (b *Box3) Translate(offset float64) *Box3 {
	b.p.Call("translate", offset)
	return b
}

// Equals TODO description.
func (b *Box3) Equals(box float64) *Box3 {
	b.p.Call("equals", box)
	return b
}
