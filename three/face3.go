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

// Face3 is a triangle face.
//
// http://threejs.org/docs/index.html#Reference/Core/Face3
type Face3 struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (f *Face3) JSObject() *js.Object { return f.p }

// Face3 returns a Face3 JavaScript class.
func (t *Three) Face3() *Face3 {
	p := t.ctx.Get("Face3")
	return Face3FromJSObject(p)
}

// Face3FromJSObject returns a wrapped Face3 JavaScript class.
func Face3FromJSObject(p *js.Object) *Face3 {
	return &Face3{p: p}
}

// NewFace3 returns a new Face3 object.
//
//     a — Vertex A index.
//     b — Vertex B index.
//     c — Vertex C index.
//     normal — Face normal or array of vertex normals.
//     color — Face color or array of vertex colors.
//     materialIndex — Material index.
func (t *Three) NewFace3(a, b, c float64, normal *Vector3, color *Color, materialIndex int) *Face3 {
	p := t.ctx.Get("Face3").New(a, b, c, normal, color, materialIndex)
	return Face3FromJSObject(p)
}

// Clone TODO description.
func (f *Face3) Clone() *Face3 {
	f.p.Call("clone")
	return f
}

// Copy TODO description.
func (f *Face3) Copy(source *Face3) *Face3 {
	f.p.Call("copy", source.p)
	return f
}

// Index returns the element (a vertex index) at this key.
// key must be "a", "b", or "c".
func (f *Face3) Index(key string) int {
	return f.p.Get(key).Int()
}

// SetVertexColor sets the colors of the indexed vertex.
// index is the vertex index and must be 0, 1, or 2.
// color is the color for the vertex.
func (f *Face3) SetVertexColor(index int, color *Color) *Face3 {
	f.p.Get("vertexColors").SetIndex(index, color.p)
	return f
}
