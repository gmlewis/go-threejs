// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Sphere represents a geometric sphere defined by a center position and radius.
//
// http://threejs.org/docs/index.html#Reference/Math/Sphere
type Sphere struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (s *Sphere) JSObject() *js.Object { return s.p }

// Sphere returns a Sphere JavaScript class.
func (t *Three) Sphere() *Sphere {
	p := t.ctx.Get("Sphere")
	return SphereFromJSObject(p)
}

// SphereFromJSObject returns a wrapped Sphere JavaScript class.
func SphereFromJSObject(p *js.Object) *Sphere {
	return &Sphere{p: p}
}

// NewSphere returns a new Sphere object.
//
//     center -- Vector3
//     radius -- Float
func (t *Three) NewSphere(center, radius float64) *Sphere {
	p := t.ctx.Get("Sphere").New(center, radius)
	return SphereFromJSObject(p)
}

// Set sets the center and radius.
//
//     center -- Vector3
//     radius -- Float
func (s *Sphere) Set(center, radius float64) *Sphere {
	s.p.Call("set", center, radius)
	return s
}

// SetFromPoints computes the minimum bounding sphere for points. If optionalCenter is given,
// it is used as the sphere's center. Otherwise, the center of the axis-aligned bounding box
// encompassing points is calculated.
//
//     points -- Array of Vector3 positions.
//     optionalCenter -- Optional Vector3 position for the sphere's center.
func (s *Sphere) SetFromPoints() *Sphere {
	s.p.Call("setFromPoints")
	return s
}

// Clone provides a new copy of the sphere.
func (s *Sphere) Clone() *Sphere {
	s.p.Call("clone")
	return s
}

// Copy copies the values of the passed sphere to this sphere.
//
//     sphere -- Sphere to copy
func (s *Sphere) Copy(sphere *Sphere) *Sphere {
	s.p.Call("copy", sphere.p)
	return s
}

// Empty checks to see if the sphere is empty (the radius set to 0).
func (s *Sphere) Empty() *Sphere {
	s.p.Call("empty")
	return s
}

// ContainsPoint checks to see if the sphere contains the provided point inclusive of the edge of the sphere.
//
//     point -- Vector3
func (s *Sphere) ContainsPoint(point float64) *Sphere {
	s.p.Call("containsPoint", point)
	return s
}

// DistanceToPoint Returns the closest distance from the boundary of the sphere to the point. If the sphere
// contains the point, the distance will be negative.
//
//     point -- Vector3
func (s *Sphere) DistanceToPoint(point float64) *Sphere {
	s.p.Call("distanceToPoint", point)
	return s
}

// IntersectsSphere checks to see if two spheres intersect.
//
//     sphere -- Sphere
func (s *Sphere) IntersectsSphere(sphere float64) *Sphere {
	s.p.Call("intersectsSphere", sphere)
	return s
}

// IntersectsBox TODO description.
func (s *Sphere) IntersectsBox(box float64) *Sphere {
	s.p.Call("intersectsBox", box)
	return s
}

// IntersectsPlane TODO description.
func (s *Sphere) IntersectsPlane(plane float64) *Sphere {
	s.p.Call("intersectsPlane", plane)
	return s
}

// ClampPoint clamps a point within the sphere. If the point is is outside the sphere, it will
// clamp it to the closets point on the edge of the sphere.
//
//     point -- Vector3 The point to clamp
//     optionalTarget -- Vector3 The optional target point to return
func (s *Sphere) ClampPoint(point, optionalTarget float64) *Sphere {
	s.p.Call("clampPoint", point, optionalTarget)
	return s
}

// GetBoundingBox Returns a bounding box for the sphere, optionally setting a provided box target.
//
//     optionalTarget -- Box
func (s *Sphere) GetBoundingBox(optionalTarget float64) *Sphere {
	s.p.Call("getBoundingBox", optionalTarget)
	return s
}

// ApplyMatrix4 transforms this sphere with the provided Matrix4.
//
//     matrix -- Matrix4
func (s *Sphere) ApplyMatrix4(matrix float64) *Sphere {
	s.p.Call("applyMatrix4", matrix)
	return s
}

// Translate translates the sphere's center by the provided offset vector.
//
//     offset -- Vector3
func (s *Sphere) Translate(offset float64) *Sphere {
	s.p.Call("translate", offset)
	return s
}

// Equals checks to see if the two spheres' centers and radii are equal.
//
//     sphere -- Sphere
func (s *Sphere) Equals(sphere float64) *Sphere {
	s.p.Call("equals", sphere)
	return s
}
