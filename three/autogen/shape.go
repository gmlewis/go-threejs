package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Shape represents a shape.
type Shape struct{ p *js.Object }

// Shape returns a shape object.
func (t *Three) Shape() *Shape {
	p := t.ctx.Get("Shape")
	return &Shape{p: p}
}

// NewShape returns a new shape object.
func (t *Shape) New() *Shape {
	p := t.p.New()
	return &Shape{p: p}
}

// Extrude TODO description.
func (s *Shape) Extrude(options float64) *Shape {
	s.p.Call("extrude", options)
	return s
}

// MakeGeometry TODO description.
func (s *Shape) MakeGeometry(options float64) *Shape {
	s.p.Call("makeGeometry", options)
	return s
}

// GetPointsHoles TODO description.
func (s *Shape) GetPointsHoles(divisions float64) *Shape {
	s.p.Call("getPointsHoles", divisions)
	return s
}

// ExtractAllPoints TODO description.
func (s *Shape) ExtractAllPoints(divisions float64) *Shape {
	s.p.Call("extractAllPoints", divisions)
	return s
}

// ExtractPoints TODO description.
func (s *Shape) ExtractPoints(divisions float64) *Shape {
	s.p.Call("extractPoints", divisions)
	return s
}

