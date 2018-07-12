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

// BufferGeometry is an efficient alternative to Geometry, because it
// stores all data, including vertex positions, face indices,
// normals, colors, UVs, and custom attributes within buffers; this
// reduces the cost of passing all this data to the GPU. This also
// makes BufferGeometry harder to work with than Geometry; rather
// than accessing position data as Vector3 objects, color data as
// Color objects, and so on, you have to access the raw data from the
// appropriate attribute buffer. This makes BufferGeometry
// best-suited for static objects where you don't need to manipulate
// the geometry much after instantiating it.
//
// http://threejs.org/docs/index.html#Reference/Core/BufferGeometry
type BufferGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (b *BufferGeometry) JSObject() *js.Object { return b.p }

// BufferGeometry returns a BufferGeometry JavaScript class.
func (t *Three) BufferGeometry() *BufferGeometry {
	p := t.ctx.Get("BufferGeometry")
	return BufferGeometryFromJSObject(p)
}

// BufferGeometryFromJSObject returns a wrapped Geometry JavaScript class.
func BufferGeometryFromJSObject(p *js.Object) *BufferGeometry {
	return &BufferGeometry{p: p}
}

// NewBufferGeometry returns a new BufferGeometry object.
func (t *Three) NewBufferGeometry() *BufferGeometry {
	p := t.ctx.Get("BufferGeometry").New()
	return BufferGeometryFromJSObject(p)
}

// GetIndex TODO description.
func (b *BufferGeometry) GetIndex() *BufferGeometry {
	b.p.Call("getIndex")
	return b
}

// SetIndex TODO description.
func (b *BufferGeometry) SetIndex(index int) *BufferGeometry {
	b.p.Call("setIndex", index)
	return b
}

// AddAttribute TODO description.
func (b *BufferGeometry) AddAttribute(name string, attribute *js.Object) *BufferGeometry {
	b.p.Call("addAttribute", name, attribute)
	return b
}

// GetAttribute TODO description.
func (b *BufferGeometry) GetAttribute(name string) *BufferGeometry {
	b.p.Call("getAttribute", name)
	return b
}

// RemoveAttribute TODO description.
func (b *BufferGeometry) RemoveAttribute(name string) *BufferGeometry {
	b.p.Call("removeAttribute", name)
	return b
}

// AddGroup TODO description.
func (b *BufferGeometry) AddGroup(start, count, materialIndex int) *BufferGeometry {
	b.p.Call("addGroup", start, count, materialIndex)
	return b
}

// ClearGroups TODO description.
func (b *BufferGeometry) ClearGroups() *BufferGeometry {
	b.p.Call("clearGroups")
	return b
}

// SetDrawRange TODO description.
func (b *BufferGeometry) SetDrawRange(start, count int) *BufferGeometry {
	b.p.Call("setDrawRange", start, count)
	return b
}

// ApplyMatrix TODO description.
func (b *BufferGeometry) ApplyMatrix(matrix *Matrix4) *BufferGeometry {
	b.p.Call("applyMatrix", matrix)
	return b
}

// RotateX TODO description.
func (b *BufferGeometry) RotateX() *BufferGeometry {
	b.p.Call("rotateX")
	return b
}

// RotateY TODO description.
func (b *BufferGeometry) RotateY() *BufferGeometry {
	b.p.Call("rotateY")
	return b
}

// RotateZ TODO description.
func (b *BufferGeometry) RotateZ() *BufferGeometry {
	b.p.Call("rotateZ")
	return b
}

// Translate TODO description.
func (b *BufferGeometry) Translate() *BufferGeometry {
	b.p.Call("translate")
	return b
}

// Scale TODO description.
func (b *BufferGeometry) Scale() *BufferGeometry {
	b.p.Call("scale")
	return b
}

// LookAt TODO description.
func (b *BufferGeometry) LookAt() *BufferGeometry {
	b.p.Call("lookAt")
	return b
}

// Center TODO description.
func (b *BufferGeometry) Center() *BufferGeometry {
	b.p.Call("center")
	return b
}

// SetFromObject TODO description.
func (b *BufferGeometry) SetFromObject(object JSObject) *BufferGeometry {
	b.p.Call("setFromObject", object.JSObject())
	return b
}

// UpdateFromObject TODO description.
func (b *BufferGeometry) UpdateFromObject(object JSObject) *BufferGeometry {
	b.p.Call("updateFromObject", object.JSObject())
	return b
}

// FromGeometry TODO description.
func (b *BufferGeometry) FromGeometry(geometry JSObject) *BufferGeometry {
	b.p.Call("fromGeometry", geometry.JSObject())
	return b
}

// FromDirectGeometry TODO description.
func (b *BufferGeometry) FromDirectGeometry(geometry JSObject) *BufferGeometry {
	b.p.Call("fromDirectGeometry", geometry.JSObject())
	return b
}

// ComputeBoundingBox TODO description.
func (b *BufferGeometry) ComputeBoundingBox() *BufferGeometry {
	b.p.Call("computeBoundingBox")
	return b
}

// ComputeBoundingSphere TODO description.
func (b *BufferGeometry) ComputeBoundingSphere() *BufferGeometry {
	b.p.Call("computeBoundingSphere")
	return b
}

// ComputeFaceNormals TODO description.
func (b *BufferGeometry) ComputeFaceNormals() *BufferGeometry {
	b.p.Call("computeFaceNormals")
	return b
}

// ComputeVertexNormals TODO description.
func (b *BufferGeometry) ComputeVertexNormals() *BufferGeometry {
	b.p.Call("computeVertexNormals")
	return b
}

// Merge TODO description.
func (b *BufferGeometry) Merge(geometry JSObject, offset int) *BufferGeometry {
	b.p.Call("merge", geometry.JSObject(), offset)
	return b
}

// NormalizeNormals TODO description.
func (b *BufferGeometry) NormalizeNormals() *BufferGeometry {
	b.p.Call("normalizeNormals")
	return b
}

// ToNonIndexed TODO description.
func (b *BufferGeometry) ToNonIndexed() *BufferGeometry {
	b.p.Call("toNonIndexed")
	return b
}

// ToJSON TODO description.
func (b *BufferGeometry) ToJSON() *BufferGeometry {
	b.p.Call("toJSON")
	return b
}

// Clone TODO description.
func (b *BufferGeometry) Clone() *BufferGeometry {
	b.p.Call("clone")
	return b
}

// Copy TODO description.
func (b *BufferGeometry) Copy(source *BufferGeometry) *BufferGeometry {
	b.p.Call("copy", source.p)
	return b
}

// Dispose TODO description.
func (b *BufferGeometry) Dispose() *BufferGeometry {
	b.p.Call("dispose")
	return b
}

// UUID returns the property of the same name.
func (b *BufferGeometry) UUID() int {
	return b.p.Get("uuid").Int()
}

// Name returns the property of the same name.
func (b *BufferGeometry) Name() string {
	return b.p.Get("name").String()
}

// Type returns the property of the same name.
func (b *BufferGeometry) Type() string {
	return b.p.Get("type").String()
}

// Index returns the property of the same name.
func (b *BufferGeometry) Index() int {
	return b.p.Get("index").Int()
}

// Attributes returns the property of the same name.
func (b *BufferGeometry) Attributes() *js.Object {
	return b.p.Get("attributes")
}

// MortphAttributes returns the property of the same name.
func (b *BufferGeometry) MortphAttributes() *js.Object {
	return b.p.Get("morphAttributes")
}

// Groups returns the property of the same name.
func (b *BufferGeometry) Groups() *js.Object {
	return b.p.Get("groups")
}

// BoundingBox returns the property of the same name.
func (b *BufferGeometry) BoundingBox() *Box3 {
	return Box3FromJSObject(b.p.Get("boundingBox"))
}

// BoundingSphere returns the property of the same name.
func (b *BufferGeometry) BoundingSphere() *js.Object {
	return b.p.Get("boundingSphere")
}

// DrawRange returns the property of the same name.
func (b *BufferGeometry) DrawRange() (start, count int) {
	t := b.p.Get("drawRange")
	return t.Get("start").Int(), t.Get("count").Int()
}
