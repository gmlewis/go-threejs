package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Ray represents a ray.
type Ray struct{ p *js.Object }

// Ray returns a ray object.
func (t *Three) Ray() *Ray {
	p := t.ctx.Get("Ray")
	return &Ray{p: p}
}

// NewRay returns a new ray object.
func (t *Ray) New(origin, direction float64) *Ray {
	p := t.p.New(origin, direction)
	return &Ray{p: p}
}

// Set TODO description.
func (r *Ray) Set(origin, direction float64) *Ray {
	r.p.Call("set", origin, direction)
	return r
}

// Clone TODO description.
func (r *Ray) Clone() *Ray {
	r.p.Call("clone")
	return r
}

// Copy TODO description.
func (r *Ray) Copy(ray float64) *Ray {
	r.p.Call("copy", ray)
	return r
}

// At TODO description.
func (r *Ray) At(t, optionalTarget float64) *Ray {
	r.p.Call("at", t, optionalTarget)
	return r
}

// LookAt TODO description.
func (r *Ray) LookAt(v float64) *Ray {
	r.p.Call("lookAt", v)
	return r
}

// Recast TODO description.
func (r *Ray) Recast() *Ray {
	r.p.Call("recast")
	return r
}

// ClosestPointToPoint TODO description.
func (r *Ray) ClosestPointToPoint(point, optionalTarget float64) *Ray {
	r.p.Call("closestPointToPoint", point, optionalTarget)
	return r
}

// DistanceToPoint TODO description.
func (r *Ray) DistanceToPoint(point float64) *Ray {
	r.p.Call("distanceToPoint", point)
	return r
}

// DistanceSqToPoint TODO description.
func (r *Ray) DistanceSqToPoint() *Ray {
	r.p.Call("distanceSqToPoint")
	return r
}

// DistanceSqToSegment TODO description.
func (r *Ray) DistanceSqToSegment() *Ray {
	r.p.Call("distanceSqToSegment")
	return r
}

// IntersectSphere TODO description.
func (r *Ray) IntersectSphere() *Ray {
	r.p.Call("intersectSphere")
	return r
}

// IntersectsSphere TODO description.
func (r *Ray) IntersectsSphere(sphere float64) *Ray {
	r.p.Call("intersectsSphere", sphere)
	return r
}

// DistanceToPlane TODO description.
func (r *Ray) DistanceToPlane(plane float64) *Ray {
	r.p.Call("distanceToPlane", plane)
	return r
}

// IntersectPlane TODO description.
func (r *Ray) IntersectPlane(plane, optionalTarget float64) *Ray {
	r.p.Call("intersectPlane", plane, optionalTarget)
	return r
}

// IntersectsPlane TODO description.
func (r *Ray) IntersectsPlane(plane float64) *Ray {
	r.p.Call("intersectsPlane", plane)
	return r
}

// IntersectBox TODO description.
func (r *Ray) IntersectBox(box, optionalTarget float64) *Ray {
	r.p.Call("intersectBox", box, optionalTarget)
	return r
}

// IntersectTriangle TODO description.
func (r *Ray) IntersectTriangle() *Ray {
	r.p.Call("intersectTriangle")
	return r
}

// ApplyMatrix4 TODO description.
func (r *Ray) ApplyMatrix4(matrix4 float64) *Ray {
	r.p.Call("applyMatrix4", matrix4)
	return r
}

// Equals TODO description.
func (r *Ray) Equals(ray float64) *Ray {
	r.p.Call("equals", ray)
	return r
}

