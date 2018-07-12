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

// Shape defines an arbitrary 2d shape plane using paths with optional holes. It can be used with
// ExtrudeGeometry, ShapeGeometry, to get points, or to get triangulated faces.
type Shape struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (s *Shape) JSObject() *js.Object { return s.p }

// Shape returns a Shape JavaScript class.
func (t *Three) Shape() *Shape {
	p := t.ctx.Get("Shape")
	return ShapeFromJSObject(p)
}

// ShapeFromJSObject returns a wrapped Shape JavaScript class.
func ShapeFromJSObject(p *js.Object) *Shape {
	return &Shape{p: p}
}

// NewShape returns a new Shape object.
func (t *Three) NewShape() *Shape {
	p := t.ctx.Get("Shape").New()
	return ShapeFromJSObject(p)
}

// Extrude convenience method to return ExtrudeGeometry
//
//     options -- This is passed as the second argument to ExtrudeGeometry
func (s *Shape) Extrude(options float64) *Shape {
	s.p.Call("extrude", options)
	return s
}

// MakeGeometry convenience method to return ShapeGeometry
//
//     options -- This is passed as the second argument to ShapeGeometry
func (s *Shape) MakeGeometry(options float64) *Shape {
	s.p.Call("makeGeometry", options)
	return s
}

// GetPointsHoles gets an array of [page Vector2 Vector2s] that represent the holes in the shape.
//
//     divisions -- The number of divisions to create on the shape
func (s *Shape) GetPointsHoles(divisions float64) *Shape {
	s.p.Call("getPointsHoles", divisions)
	return s
}

// ExtractAllPoints get points of shape and holes (keypoints based on segments parameter).
//
//     divisions -- The number of divisions to create on the shape
func (s *Shape) ExtractAllPoints(divisions float64) *Shape {
	s.p.Call("extractAllPoints", divisions)
	return s
}

// ExtractPoints returns an object with a shape and holes property that each reference an array of Vector2s.
//
//     divisions -- The number of divisions to create on the shape
func (s *Shape) ExtractPoints(divisions float64) *Shape {
	s.p.Call("extractPoints", divisions)
	return s
}
