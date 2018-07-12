// Copyright 2016 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
