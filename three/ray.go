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

// Ray is a ray that emits from an origin in a certain direction.
//
// http://threejs.org/docs/index.html#Reference/Math/Ray
type Ray struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (r *Ray) JSObject() *js.Object { return r.p }

// Ray returns a Ray JavaScript class.
func (t *Three) Ray() *Ray {
	p := t.ctx.Get("Ray")
	return RayFromJSObject(p)
}

// RayFromJSObject returns a wrapped Ray JavaScript class.
func RayFromJSObject(p *js.Object) *Ray {
	return &Ray{p: p}
}

// NewRay returns a new Ray object. Initialises the origin and direction properties to the provided values.
//
//     origin -- Vector3 The origin of the Ray.
//     direction -- Vector3 The direction of the Ray. This must be normalized (with Vector3.normalize) for the
//     methods to operate properly.
func (t *Three) NewRay(origin, direction float64) *Ray {
	p := t.ctx.Get("Ray").New(origin, direction)
	return RayFromJSObject(p)
}

// Set copies the parameters to the origin and direction properties.
//
//     origin -- Vector3 The origin of the Ray.
//     direction -- Vector3 The direction of the Ray. This must be normalized (with Vector3.normalize) for the
//     methods to operate properly.
func (r *Ray) Set(origin, direction float64) *Ray {
	r.p.Call("set", origin, direction)
	return r
}

// Clone creates a clone of this Ray.
func (r *Ray) Clone() *Ray {
	r.p.Call("clone")
	return r
}

// Copy copies the properties of the provided Ray, then returns this Ray.
//
//     ray -- Ray The Ray to copy values from.
func (r *Ray) Copy(ray *Ray) *Ray {
	r.p.Call("copy", ray.p)
	return r
}

// At gets a Vector3 that is a given distance along this Ray.
//
//     t -- Float The distance along the Ray to retrieve a position for.
//     optionalTarget -- Vector3 Receives the position along the Ray if passed; otherwise a new Vector3 is created.
func (r *Ray) At(t, optionalTarget float64) *Ray {
	r.p.Call("at", t, optionalTarget)
	return r
}

// LookAt adjusts the direction of the ray to point at the vector in world coordinates.
//
//     v -- Vector3 The vector to look at.
func (r *Ray) LookAt(v float64) *Ray {
	r.p.Call("lookAt", v)
	return r
}

// Recast shifts the origin of this Ray along its direction by the distance given.
//
//     v -- Vector3 The vector to look at.
func (r *Ray) Recast() *Ray {
	r.p.Call("recast")
	return r
}

// ClosestPointToPoint gets the point along this Ray that is closest to the Vector3 provided.
//
//     point -- Vector3 The point to get the closest approach to.
//     optionalTarget -- Vector3 Receives the return value if passed; otherwise a new Vector3 is created.
func (r *Ray) ClosestPointToPoint(point, optionalTarget float64) *Ray {
	r.p.Call("closestPointToPoint", point, optionalTarget)
	return r
}

// DistanceToPoint gets the distance of the closest approach between the Ray and the Vector3.
//
//     point -- Vector3 The Vector3 to compute a distance to.
func (r *Ray) DistanceToPoint(point float64) *Ray {
	r.p.Call("distanceToPoint", point)
	return r
}

// DistanceSqToPoint gets the squared distance of the closest approach between the Ray and the Vector3.
//
//     point -- Vector3 The Vector3 to compute a distance to.
func (r *Ray) DistanceSqToPoint() *Ray {
	r.p.Call("distanceSqToPoint")
	return r
}

// DistanceSqToSegment gets the squared distance between this Ray and a line segment.
//
//     v0 -- Vector3 The start of the line segment.
//     v1 -- Vector3 The end of the line segment.
//     optionalPointOnRay -- Vector3 If this is provided, it receives the point on this Ray that is
//     closest to the segment.
//     optionalPointOnSegment -- Vector3 If this is provided, it receives the point on the line segment that is
//     closest to this Ray.
func (r *Ray) DistanceSqToSegment() *Ray {
	r.p.Call("distanceSqToSegment")
	return r
}

// IntersectSphere TODO description.
func (r *Ray) IntersectSphere() *Ray {
	r.p.Call("intersectSphere")
	return r
}

// IntersectsSphere  returns whether or not this Ray intersects with the Sphere.
//
//     sphere -- Sphere The Sphere to intersect with.
func (r *Ray) IntersectsSphere(sphere float64) *Ray {
	r.p.Call("intersectsSphere", sphere)
	return r
}

// DistanceToPlane gets the distance from the origin to the Plane, or null if the Ray doesn't intersect the Plane.
//
//     plane -- Plane The Plane to get the distance to.
func (r *Ray) DistanceToPlane(plane float64) *Ray {
	r.p.Call("distanceToPlane", plane)
	return r
}

// IntersectPlane intersects this Ray with a Plane, returning the intersection point or null if there is no intersection.
//
//     plane -- Plane The Plane to intersect with.
//     optionalTarget -- Vector3 The Vector3 to store the result in, or null to create a new Vector3.
func (r *Ray) IntersectPlane(plane, optionalTarget float64) *Ray {
	r.p.Call("intersectPlane", plane, optionalTarget)
	return r
}

// IntersectsPlane returns whether or not this Ray intersects with the Plane.
//
//     plane -- Plane The Plane to intersect with.
func (r *Ray) IntersectsPlane(plane float64) *Ray {
	r.p.Call("intersectsPlane", plane)
	return r
}

// IntersectBox intersects this Ray with a Box3, returning the intersection point or null if there is no intersection.
//
//     box -- Box3 The Box3 to intersect with.
//     optionalTarget -- Vector3 The Vector3 to store the result in, or null to create a new Vector3.
func (r *Ray) IntersectBox(box, optionalTarget float64) *Ray {
	r.p.Call("intersectBox", box, optionalTarget)
	return r
}

// IntersectTriangle intersects this Ray with a triangle, returning the intersection point or null if there is no
// intersection.
//
//     a, b, c -- Vector3 The Vector3 points on the triangle.
//     backfaceCulling -- Boolean Whether to use backface culling.
//     optionalTarget -- Vector3 The Vector3 to store the result in, or null to create a new Vector3.
func (r *Ray) IntersectTriangle() *Ray {
	r.p.Call("intersectTriangle")
	return r
}

// ApplyMatrix4 transforms this Ray by the Matrix4.
//
//     matrix4 -- Matrix4 The Matrix4 to transform this Ray by.
func (r *Ray) ApplyMatrix4(matrix4 float64) *Ray {
	r.p.Call("applyMatrix4", matrix4)
	return r
}

// Equals returns whether this and the other Ray have equal offsets and directions.
//
//     ray -- Ray The Ray to compare to.
func (r *Ray) Equals(ray float64) *Ray {
	r.p.Call("equals", ray)
	return r
}
