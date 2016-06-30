// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Bone represents a bone which is part of a skeleton.
//
// http://threejs.org/docs/index.html#Reference/Objects/Bone
type Bone struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *Bone) JSObject() *js.Object { return t.p }

// Bone returns a Bone object.
func (t *Three) Bone() *Bone {
	p := t.ctx.Get("Bone")
	return &Bone{p: p}
}

// New returns a new Bone object.
func (t *Bone) New(skin float64) *Bone {
	p := t.p.New(skin)
	return &Bone{p: p}
}

// Copy TODO description.
func (b *Bone) Copy(source float64) *Bone {
	b.p.Call("copy", source)
	return b
}
