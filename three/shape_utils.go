package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// ShapeUtils represents a shapeutils.
type ShapeUtils struct{ p *js.Object }

// ShapeUtils returns a ShapeUtils object.
func (t *Three) ShapeUtils() *ShapeUtils {
	p := t.ctx.Get("ShapeUtils")
	return &ShapeUtils{p: p}
}

// New returns a new ShapeUtils object.
func (t *ShapeUtils) New() *ShapeUtils {
	p := t.p.New()
	return &ShapeUtils{p: p}
}

// Area TODO description.
func (s *ShapeUtils) Area(contour float64) *ShapeUtils {
	s.p.Call("area", contour)
	return s
}

// TriangulateShape TODO description.
func (s *ShapeUtils) TriangulateShape(contour, holes float64) *ShapeUtils {
	s.p.Call("triangulateShape", contour, holes)
	return s
}

// IsClockWise TODO description.
func (s *ShapeUtils) IsClockWise(pts float64) *ShapeUtils {
	s.p.Call("isClockWise", pts)
	return s
}

