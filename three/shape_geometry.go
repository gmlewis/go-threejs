// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// ShapeGeometry creates a one-sided polygonal geometry from one or more path shapes. Similar to ExtrudeGeometry.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/ShapeGeometry
type ShapeGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (s *ShapeGeometry) JSObject() *js.Object { return s.p }

// ShapeGeometry returns a ShapeGeometry JavaScript class.
func (t *Three) ShapeGeometry() *ShapeGeometry {
	p := t.ctx.Get("ShapeGeometry")
	return ShapeGeometryFromJSObject(p)
}

// ShapeGeometryFromJSObject returns a wrapped ShapeGeometry JavaScript class.
func ShapeGeometryFromJSObject(p *js.Object) *ShapeGeometry {
	return &ShapeGeometry{p: p}
}

// NewShapeGeometry returns a new ShapeGeometry object.
func (t *Three) NewShapeGeometry(shapes, options float64) *ShapeGeometry {
	p := t.ctx.Get("ShapeGeometry").New(shapes, options)
	return ShapeGeometryFromJSObject(p)
}

// AddShapeList adds a list of shapes to the geometry.
//
//     shapes — Array of shapes
//     options — See options in constructor
func (s *ShapeGeometry) AddShapeList(shapes, options float64) *ShapeGeometry {
	s.p.Call("addShapeList", shapes, options)
	return s
}

// AddShape adds a single shape to the geometry
//
//     shape — Shape
//     options — See options in constructor
func (s *ShapeGeometry) AddShape(shape, options float64) *ShapeGeometry {
	s.p.Call("addShape", shape, options)
	return s
}
