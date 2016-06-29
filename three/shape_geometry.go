package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// ShapeGeometry represents a shapegeometry.
type ShapeGeometry struct{ p *js.Object }

// ShapeGeometry returns a ShapeGeometry object.
func (t *Three) ShapeGeometry() *ShapeGeometry {
	p := t.ctx.Get("ShapeGeometry")
	return &ShapeGeometry{p: p}
}

// New returns a new ShapeGeometry object.
func (t *ShapeGeometry) New(shapes, options float64) *ShapeGeometry {
	p := t.p.New(shapes, options)
	return &ShapeGeometry{p: p}
}

// AddShapeList TODO description.
func (s *ShapeGeometry) AddShapeList(shapes, options float64) *ShapeGeometry {
	s.p.Call("addShapeList", shapes, options)
	return s
}

// AddShape TODO description.
func (s *ShapeGeometry) AddShape(shape, options float64) *ShapeGeometry {
	s.p.Call("addShape", shape, options)
	return s
}

