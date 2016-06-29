package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Box2 represents a box2.
type Box2 struct{ p *js.Object }

// Box2 returns a Box2 object.
func (t *Three) Box2() *Box2 {
	p := t.ctx.Get("Box2")
	return &Box2{p: p}
}

// New returns a new Box2 object.
func (t *Box2) New(min, max float64) *Box2 {
	p := t.p.New(min, max)
	return &Box2{p: p}
}

// Set TODO description.
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
func (b *Box2) Copy(box float64) *Box2 {
	b.p.Call("copy", box)
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

// ExpandByPoint TODO description.
func (b *Box2) ExpandByPoint(point float64) *Box2 {
	b.p.Call("expandByPoint", point)
	return b
}

// ExpandByVector TODO description.
func (b *Box2) ExpandByVector(vector float64) *Box2 {
	b.p.Call("expandByVector", vector)
	return b
}

// ExpandByScalar TODO description.
func (b *Box2) ExpandByScalar(scalar float64) *Box2 {
	b.p.Call("expandByScalar", scalar)
	return b
}

// ContainsPoint TODO description.
func (b *Box2) ContainsPoint(point float64) *Box2 {
	b.p.Call("containsPoint", point)
	return b
}

// ContainsBox TODO description.
func (b *Box2) ContainsBox(box float64) *Box2 {
	b.p.Call("containsBox", box)
	return b
}

// GetParameter TODO description.
func (b *Box2) GetParameter(point, optionalTarget float64) *Box2 {
	b.p.Call("getParameter", point, optionalTarget)
	return b
}

// IntersectsBox TODO description.
func (b *Box2) IntersectsBox(box float64) *Box2 {
	b.p.Call("intersectsBox", box)
	return b
}

// ClampPoint TODO description.
func (b *Box2) ClampPoint(point, optionalTarget float64) *Box2 {
	b.p.Call("clampPoint", point, optionalTarget)
	return b
}

// DistanceToPoint TODO description.
func (b *Box2) DistanceToPoint() *Box2 {
	b.p.Call("distanceToPoint")
	return b
}

// Intersect TODO description.
func (b *Box2) Intersect(box float64) *Box2 {
	b.p.Call("intersect", box)
	return b
}

// Union TODO description.
func (b *Box2) Union(box float64) *Box2 {
	b.p.Call("union", box)
	return b
}

// Translate TODO description.
func (b *Box2) Translate(offset float64) *Box2 {
	b.p.Call("translate", offset)
	return b
}

// Equals TODO description.
func (b *Box2) Equals(box float64) *Box2 {
	b.p.Call("equals", box)
	return b
}

