// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// InstancedBufferGeometry represents an instancedbuffergeometry.
type InstancedBufferGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *InstancedBufferGeometry) JSObject() *js.Object { return t.p }

// InstancedBufferGeometry returns an InstancedBufferGeometry object.
func (t *Three) InstancedBufferGeometry() *InstancedBufferGeometry {
	p := t.ctx.Get("InstancedBufferGeometry")
	return &InstancedBufferGeometry{p: p}
}

// New returns a new InstancedBufferGeometry object.
func (t *InstancedBufferGeometry) New() *InstancedBufferGeometry {
	p := t.p.New()
	return &InstancedBufferGeometry{p: p}
}

// AddGroup TODO description.
func (i *InstancedBufferGeometry) AddGroup(start, count, instances float64) *InstancedBufferGeometry {
	i.p.Call("addGroup", start, count, instances)
	return i
}

// Copy TODO description.
func (i *InstancedBufferGeometry) Copy(source float64) *InstancedBufferGeometry {
	i.p.Call("copy", source)
	return i
}