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

import "github.com/gopherjs/gopherjs/js"

// Geometry is a Javascript base class for geometries.
// A geometry holds all data necessary to describe a 3D model.
//
// http://threejs.org/docs/index.html#Reference/Core/Geometry
type Geometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (g *Geometry) JSObject() *js.Object { return g.p }

// Geometry returns a Geometry JavaScript class.
func (t *Three) Geometry() *Geometry {
	p := t.ctx.Get("Geometry")
	return GeometryFromJSObject(p)
}

// GeometryFromJSObject returns a wrapped Geometry JavaScript class.
func GeometryFromJSObject(p *js.Object) *Geometry {
	return &Geometry{p: p}
}

// NewGeometry returns a new Geometry object.
func (t *Three) NewGeometry() *Geometry {
	p := t.ctx.Get("Geometry").New()
	return GeometryFromJSObject(p)
}

// ApplyMatrix TODO description.
func (g *Geometry) ApplyMatrix(matrix float64) *Geometry {
	g.p.Call("applyMatrix", matrix)
	return g
}

// RotateX TODO description.
func (g *Geometry) RotateX() *Geometry {
	g.p.Call("rotateX")
	return g
}

// RotateY TODO description.
func (g *Geometry) RotateY() *Geometry {
	g.p.Call("rotateY")
	return g
}

// RotateZ TODO description.
func (g *Geometry) RotateZ() *Geometry {
	g.p.Call("rotateZ")
	return g
}

// Translate TODO description.
func (g *Geometry) Translate() *Geometry {
	g.p.Call("translate")
	return g
}

// Scale TODO description.
func (g *Geometry) Scale() *Geometry {
	g.p.Call("scale")
	return g
}

// LookAt TODO description.
func (g *Geometry) LookAt() *Geometry {
	g.p.Call("lookAt")
	return g
}

// FromBufferGeometry TODO description.
func (g *Geometry) FromBufferGeometry(geometry float64) *Geometry {
	g.p.Call("fromBufferGeometry", geometry)
	return g
}

// Center TODO description.
func (g *Geometry) Center() *Geometry {
	g.p.Call("center")
	return g
}

// Normalize TODO description.
func (g *Geometry) Normalize() *Geometry {
	g.p.Call("normalize")
	return g
}

// ComputeFaceNormals TODO description.
func (g *Geometry) ComputeFaceNormals() *Geometry {
	g.p.Call("computeFaceNormals")
	return g
}

// ComputeVertexNormals TODO description.
func (g *Geometry) ComputeVertexNormals(areaWeighted bool) *Geometry {
	g.p.Call("computeVertexNormals", areaWeighted)
	return g
}

// ComputeMorphNormals TODO description.
func (g *Geometry) ComputeMorphNormals() *Geometry {
	g.p.Call("computeMorphNormals")
	return g
}

// ComputeTangents TODO description.
func (g *Geometry) ComputeTangents() *Geometry {
	g.p.Call("computeTangents")
	return g
}

// ComputeLineDistances TODO description.
func (g *Geometry) ComputeLineDistances() *Geometry {
	g.p.Call("computeLineDistances")
	return g
}

// ComputeBoundingBox TODO description.
func (g *Geometry) ComputeBoundingBox() *Geometry {
	g.p.Call("computeBoundingBox")
	return g
}

// ComputeBoundingSphere TODO description.
func (g *Geometry) ComputeBoundingSphere() *Geometry {
	g.p.Call("computeBoundingSphere")
	return g
}

// Merge TODO description.
func (g *Geometry) Merge(geometry, matrix, materialIndexOffset float64) *Geometry {
	g.p.Call("merge", geometry, matrix, materialIndexOffset)
	return g
}

// MergeMesh TODO description.
func (g *Geometry) MergeMesh(mesh float64) *Geometry {
	g.p.Call("mergeMesh", mesh)
	return g
}

// MergeVertices TODO description.
func (g *Geometry) MergeVertices() *Geometry {
	g.p.Call("mergeVertices")
	return g
}

// SortFacesByMaterialIndex TODO description.
func (g *Geometry) SortFacesByMaterialIndex() *Geometry {
	g.p.Call("sortFacesByMaterialIndex")
	return g
}

// ToJSON TODO description.
func (g *Geometry) ToJSON() *Geometry {
	g.p.Call("toJSON")
	return g
}

// Clone TODO description.
func (g *Geometry) Clone() *Geometry {
	g.p.Call("clone")
	return g
}

// Copy TODO description.
func (g *Geometry) Copy(source *Geometry) *Geometry {
	g.p.Call("copy", source.p)
	return g
}

// Dispose TODO description.
func (g *Geometry) Dispose() *Geometry {
	g.p.Call("dispose")
	return g
}

// Faces returns a slice of the geometry's faces.
func (g *Geometry) Faces() []*Face3 {
	var result []*Face3
	faces := g.p.Get("faces")
	for i := 0; i < faces.Length(); i++ {
		result = append(result, Face3FromJSObject(faces.Index(i)))
	}
	return result
}

// Vertex returns a vertex at the given index.
func (g *Geometry) Vertex(index int) *Vector3 {
	return Vector3FromJSObject(g.p.Get("vertices").Index(index))
}

// VerticesLength returns a the length of the geometry's vertices array.
func (g *Geometry) VerticesLength() int {
	return g.p.Get("vertices").Length()
}

// Vertices returns a slice of the geometry's vertices.
func (g *Geometry) Vertices() []*Vector3 {
	var result []*Vector3
	vertices := g.p.Get("vertices")
	for i := 0; i < vertices.Length(); i++ {
		result = append(result, Vector3FromJSObject(vertices.Index(i)))
	}
	return result
}

// SetVertices sets the geometry's vertices.
func (g *Geometry) SetVertices(pts *js.Object) *Geometry {
	g.p.Set("vertices", pts)
	return g
}

// BoundingBox returns the bounding box of the Geometry.
func (g *Geometry) BoundingBox() *Box3 {
	return Box3FromJSObject(g.p.Get("boundingBox"))
}
