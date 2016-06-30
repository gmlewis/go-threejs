// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// BufferGeometry represents a buffergeometry.
type BufferGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (b *BufferGeometry) JSObject() *js.Object { return b.p }

// BufferGeometry returns a BufferGeometry object.
func (t *Three) BufferGeometry() *BufferGeometry {
	p := t.ctx.Get("BufferGeometry")
	return &BufferGeometry{p: p}
}

// New returns a new BufferGeometry object.
func (b *BufferGeometry) New() *BufferGeometry {
	p := b.p.New()
	return &BufferGeometry{p: p}
}

// GetIndex TODO description.
func (b *BufferGeometry) GetIndex() *BufferGeometry {
	b.p.Call("getIndex")
	return b
}

// SetIndex TODO description.
func (b *BufferGeometry) SetIndex(index float64) *BufferGeometry {
	b.p.Call("setIndex", index)
	return b
}

// AddAttribute TODO description.
func (b *BufferGeometry) AddAttribute(name, attribute float64) *BufferGeometry {
	b.p.Call("addAttribute", name, attribute)
	return b
}

// GetAttribute TODO description.
func (b *BufferGeometry) GetAttribute(name float64) *BufferGeometry {
	b.p.Call("getAttribute", name)
	return b
}

// RemoveAttribute TODO description.
func (b *BufferGeometry) RemoveAttribute(name float64) *BufferGeometry {
	b.p.Call("removeAttribute", name)
	return b
}

// AddGroup TODO description.
func (b *BufferGeometry) AddGroup(start, count, materialIndex float64) *BufferGeometry {
	b.p.Call("addGroup", start, count, materialIndex)
	return b
}

// ClearGroups TODO description.
func (b *BufferGeometry) ClearGroups() *BufferGeometry {
	b.p.Call("clearGroups")
	return b
}

// SetDrawRange TODO description.
func (b *BufferGeometry) SetDrawRange(start, count float64) *BufferGeometry {
	b.p.Call("setDrawRange", start, count)
	return b
}

// ApplyMatrix TODO description.
func (b *BufferGeometry) ApplyMatrix(matrix float64) *BufferGeometry {
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
func (b *BufferGeometry) SetFromObject(object float64) *BufferGeometry {
	b.p.Call("setFromObject", object)
	return b
}

// UpdateFromObject TODO description.
func (b *BufferGeometry) UpdateFromObject(object float64) *BufferGeometry {
	b.p.Call("updateFromObject", object)
	return b
}

// FromGeometry TODO description.
func (b *BufferGeometry) FromGeometry(geometry float64) *BufferGeometry {
	b.p.Call("fromGeometry", geometry)
	return b
}

// FromDirectGeometry TODO description.
func (b *BufferGeometry) FromDirectGeometry(geometry float64) *BufferGeometry {
	b.p.Call("fromDirectGeometry", geometry)
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
func (b *BufferGeometry) Merge(geometry, offset float64) *BufferGeometry {
	b.p.Call("merge", geometry, offset)
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
func (b *BufferGeometry) Copy(source float64) *BufferGeometry {
	b.p.Call("copy", source)
	return b
}

// Dispose TODO description.
func (b *BufferGeometry) Dispose() *BufferGeometry {
	b.p.Call("dispose")
	return b
}
