package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// ExtrudeGeometry represents an extrudegeometry.
type ExtrudeGeometry struct{ p *js.Object }

// ExtrudeGeometry returns an ExtrudeGeometry object.
func (t *Three) ExtrudeGeometry() *ExtrudeGeometry {
	p := t.ctx.Get("ExtrudeGeometry")
	return &ExtrudeGeometry{p: p}
}

// New returns a new ExtrudeGeometry object.
func (t *ExtrudeGeometry) New(shapes, options float64) *ExtrudeGeometry {
	p := t.p.New(shapes, options)
	return &ExtrudeGeometry{p: p}
}

// GenerateTopUV TODO description.
func (e *ExtrudeGeometry) GenerateTopUV(geometry, indexA, indexB, indexC float64) *ExtrudeGeometry {
	e.p.Call("generateTopUV", geometry, indexA, indexB, indexC)
	return e
}

// GenerateSideWallUV TODO description.
func (e *ExtrudeGeometry) GenerateSideWallUV(geometry, indexA, indexB, indexC, indexD float64) *ExtrudeGeometry {
	e.p.Call("generateSideWallUV", geometry, indexA, indexB, indexC, indexD)
	return e
}

// AddShapeList TODO description.
func (e *ExtrudeGeometry) AddShapeList(shapes, options float64) *ExtrudeGeometry {
	e.p.Call("addShapeList", shapes, options)
	return e
}

// AddShape TODO description.
func (e *ExtrudeGeometry) AddShape(shape, options float64) *ExtrudeGeometry {
	e.p.Call("addShape", shape, options)
	return e
}

