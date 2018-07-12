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

// Triangle is a geometric triangle as defined by three vectors.
//
// http://threejs.org/docs/index.html#Reference/Math/Triangle
type Triangle struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *Triangle) JSObject() *js.Object { return t.p }

// Triangle returns a Triangle JavaScript class.
func (t *Three) Triangle() *Triangle {
	p := t.ctx.Get("Triangle")
	return TriangleFromJSObject(p)
}

// TriangleFromJSObject returns a wrapped Triangle JavaScript class.
func TriangleFromJSObject(p *js.Object) *Triangle {
	return &Triangle{p: p}
}

// NewTriangle returns a new Triangle object.
func (t *Three) NewTriangle(a, b, c float64) *Triangle {
	p := t.ctx.Get("Triangle").New(a, b, c)
	return TriangleFromJSObject(p)
}

// Set sets the triangle's vectors to the passed vectors.
//
//     a -- Vector3
//     b -- Vector3
//     c -- Vector3
func (t *Triangle) Set(a, b, c float64) *Triangle {
	t.p.Call("set", a, b, c)
	return t
}

// SetFromPointsAndIndices sets the triangle's vectors to the vectors in the array.
//
//     points -- Array of Vector3s
//     i0 -- Integer index
//     i1 -- Integer index
//     i2 -- Integer index
func (t *Triangle) SetFromPointsAndIndices(points, i0, i1, i2 float64) *Triangle {
	t.p.Call("setFromPointsAndIndices", points, i0, i1, i2)
	return t
}

// Clone returns a new copy of this triangle.
func (t *Triangle) Clone() *Triangle {
	t.p.Call("clone")
	return t
}

// Copy copies the values of the vertices of the passed triangle to this triangle.
//
//     triangle -- Triangle
func (t *Triangle) Copy(triangle *Triangle) *Triangle {
	t.p.Call("copy", triangle.p)
	return t
}

// Area returns the area of the triangle.
func (t *Triangle) Area() *Triangle {
	t.p.Call("area")
	return t
}

// Midpoint returns the midpoint of the triangle. Optionally sets a target vector.
//
//     optionalTarget -- Optional Vector3 target to set the result.
func (t *Triangle) Midpoint(optionalTarget float64) *Triangle {
	t.p.Call("midpoint", optionalTarget)
	return t
}

// Normal returns the calculated normal of the triangle.
//
//     optionalTarget -- Optional Vector3 target to set the result.
func (t *Triangle) Normal(optionalTarget float64) *Triangle {
	t.p.Call("normal", optionalTarget)
	return t
}

// Plane returns a plane based on the triangle. Optionally sets a target plane.
//
//     optionalTarget -- Optional Plane target to set the result.
func (t *Triangle) Plane(optionalTarget float64) *Triangle {
	t.p.Call("plane", optionalTarget)
	return t
}

// BarycoordFromPoint returns a barycentric coordinate from the given vector.
//
//     point -- Vector3
//     optionalTarget -- Optional Vector3 target to set the result.
// http://commons.wikimedia.org/wiki/File:Barycentric_coordinates_1.png (An image of barycentric coordinates)
func (t *Triangle) BarycoordFromPoint(point, optionalTarget float64) *Triangle {
	t.p.Call("barycoordFromPoint", point, optionalTarget)
	return t
}

// ContainsPoint checks to see if the passed vector is within the triangle.
//
//     point -- Vector3
func (t *Triangle) ContainsPoint(point float64) *Triangle {
	t.p.Call("containsPoint", point)
	return t
}

// Equals checks to see if two triangles are equal (share the same vectors).
//
//     triangle -- Triangle
func (t *Triangle) Equals(triangle float64) *Triangle {
	t.p.Call("equals", triangle)
	return t
}
