package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Sphere represents a sphere.
type Sphere struct{ p *js.Object }

// Sphere returns a sphere object.
func (t *Three) Sphere() *Sphere {
	p := t.ctx.Get("Sphere")
	return &Sphere{p: p}
}

// NewSphere returns a new sphere object.
func (t *Sphere) New(center, radius float64) *Sphere {
	p := t.p.New(center, radius)
	return &Sphere{p: p}
}

// Set TODO description.
func (s *Sphere) Set(center, radius float64) *Sphere {
	s.p.Call("set", center, radius)
	return s
}

// SetFromPoints TODO description.
func (s *Sphere) SetFromPoints() *Sphere {
	s.p.Call("setFromPoints")
	return s
}

// Clone TODO description.
func (s *Sphere) Clone() *Sphere {
	s.p.Call("clone")
	return s
}

// Copy TODO description.
func (s *Sphere) Copy(sphere float64) *Sphere {
	s.p.Call("copy", sphere)
	return s
}

// Empty TODO description.
func (s *Sphere) Empty() *Sphere {
	s.p.Call("empty")
	return s
}

// ContainsPoint TODO description.
func (s *Sphere) ContainsPoint(point float64) *Sphere {
	s.p.Call("containsPoint", point)
	return s
}

// DistanceToPoint TODO description.
func (s *Sphere) DistanceToPoint(point float64) *Sphere {
	s.p.Call("distanceToPoint", point)
	return s
}

// IntersectsSphere TODO description.
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

// ClampPoint TODO description.
func (s *Sphere) ClampPoint(point, optionalTarget float64) *Sphere {
	s.p.Call("clampPoint", point, optionalTarget)
	return s
}

// GetBoundingBox TODO description.
func (s *Sphere) GetBoundingBox(optionalTarget float64) *Sphere {
	s.p.Call("getBoundingBox", optionalTarget)
	return s
}

// ApplyMatrix4 TODO description.
func (s *Sphere) ApplyMatrix4(matrix float64) *Sphere {
	s.p.Call("applyMatrix4", matrix)
	return s
}

// Translate TODO description.
func (s *Sphere) Translate(offset float64) *Sphere {
	s.p.Call("translate", offset)
	return s
}

// Equals TODO description.
func (s *Sphere) Equals(sphere float64) *Sphere {
	s.p.Call("equals", sphere)
	return s
}

