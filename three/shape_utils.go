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

// ShapeUtils represents a shapeutils.
type ShapeUtils struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (s *ShapeUtils) JSObject() *js.Object { return s.p }

// ShapeUtils returns a ShapeUtils JavaScript class.
func (t *Three) ShapeUtils() *ShapeUtils {
	p := t.ctx.Get("ShapeUtils")
	return ShapeUtilsFromJSObject(p)
}

// ShapeUtilsFromJSObject returns a wrapped ShapeUtils JavaScript class.
func ShapeUtilsFromJSObject(p *js.Object) *ShapeUtils {
	return &ShapeUtils{p: p}
}

// NewShapeUtils returns a new ShapeUtils object.
func (t *Three) NewShapeUtils() *ShapeUtils {
	p := t.ctx.Get("ShapeUtils").New()
	return ShapeUtilsFromJSObject(p)
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
