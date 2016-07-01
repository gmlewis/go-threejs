// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// InstancedBufferAttribute represents an instancedbufferattribute.
type InstancedBufferAttribute struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (i *InstancedBufferAttribute) JSObject() *js.Object { return i.p }

// InstancedBufferAttribute returns an InstancedBufferAttribute object.
func (t *Three) InstancedBufferAttribute() *InstancedBufferAttribute {
	p := t.ctx.Get("InstancedBufferAttribute")
	return &InstancedBufferAttribute{p: p}
}

// New returns a new InstancedBufferAttribute object.
func (i *InstancedBufferAttribute) New(array, itemSize, meshPerAttribute float64) *InstancedBufferAttribute {
	p := i.p.New(array, itemSize, meshPerAttribute)
	return &InstancedBufferAttribute{p: p}
}

// Copy TODO description.
func (i *InstancedBufferAttribute) Copy(source *InstancedBufferAttribute) *InstancedBufferAttribute {
	i.p.Call("copy", source.p)
	return i
}
