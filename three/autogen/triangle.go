package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Triangle represents a triangle.
type Triangle struct{ p *js.Object }

// Triangle returns a Triangle object.
func (t *Three) Triangle() *Triangle {
	p := t.ctx.Get("Triangle")
	return &Triangle{p: p}
}

// New returns a new Triangle object.
func (t *Triangle) New(a, b, c float64) *Triangle {
	p := t.p.New(a, b, c)
	return &Triangle{p: p}
}

// Set TODO description.
func (t *Triangle) Set(a, b, c float64) *Triangle {
	t.p.Call("set", a, b, c)
	return t
}

// SetFromPointsAndIndices TODO description.
func (t *Triangle) SetFromPointsAndIndices(points, i0, i1, i2 float64) *Triangle {
	t.p.Call("setFromPointsAndIndices", points, i0, i1, i2)
	return t
}

// Clone TODO description.
func (t *Triangle) Clone() *Triangle {
	t.p.Call("clone")
	return t
}

// Copy TODO description.
func (t *Triangle) Copy(triangle float64) *Triangle {
	t.p.Call("copy", triangle)
	return t
}

// Area TODO description.
func (t *Triangle) Area() *Triangle {
	t.p.Call("area")
	return t
}

// Midpoint TODO description.
func (t *Triangle) Midpoint(optionalTarget float64) *Triangle {
	t.p.Call("midpoint", optionalTarget)
	return t
}

// Normal TODO description.
func (t *Triangle) Normal(optionalTarget float64) *Triangle {
	t.p.Call("normal", optionalTarget)
	return t
}

// Plane TODO description.
func (t *Triangle) Plane(optionalTarget float64) *Triangle {
	t.p.Call("plane", optionalTarget)
	return t
}

// BarycoordFromPoint TODO description.
func (t *Triangle) BarycoordFromPoint(point, optionalTarget float64) *Triangle {
	t.p.Call("barycoordFromPoint", point, optionalTarget)
	return t
}

// ContainsPoint TODO description.
func (t *Triangle) ContainsPoint(point float64) *Triangle {
	t.p.Call("containsPoint", point)
	return t
}

// Equals TODO description.
func (t *Triangle) Equals(triangle float64) *Triangle {
	t.p.Call("equals", triangle)
	return t
}

