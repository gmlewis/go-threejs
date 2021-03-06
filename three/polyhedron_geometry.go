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

// PolyhedronGeometry represents a solid with flat faces.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/PolyhedronGeometry
type PolyhedronGeometry struct{ *Geometry }

// JSObject returns the underlying *js.Object.
func (p *PolyhedronGeometry) JSObject() *js.Object { return p.p }

// PolyhedronGeometry returns a PolyhedronGeometry JavaScript class.
func (t *Three) PolyhedronGeometry() *PolyhedronGeometry {
	p := t.ctx.Get("PolyhedronGeometry")
	return PolyhedronGeometryFromJSObject(p)
}

// PolyhedronGeometryFromJSObject returns a wrapped PolyhedronGeometry JavaScript class.
func PolyhedronGeometryFromJSObject(p *js.Object) *PolyhedronGeometry {
	return &PolyhedronGeometry{&Geometry{p: p}}
}

// NewPolyhedronGeometry returns a new PolyhedronGeometry object.
//
//     vertices — Array of points of the form [1,1,1, -1,-1,-1, ... ]
//     indices — Array of indices that make up the faces of the form [0,1,2, 2,3,0, ... ]
//     radius — Float - The radius of the final shape
//     detail — Integer - How many levels to subdivide the geometry. The more detail, the smoother the shape.
func (t *Three) NewPolyhedronGeometry(vertices []float64, indices []int, radius float64, detail int) *PolyhedronGeometry {
	p := t.ctx.Get("PolyhedronGeometry").New(vertices, indices, radius, detail)
	return PolyhedronGeometryFromJSObject(p)
}
