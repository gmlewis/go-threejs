// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// ShapeGeometry represents a shapegeometry.
type ShapeGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *ShapeGeometry) JSObject() *js.Object { return t.p }

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