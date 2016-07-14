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
func (b *Bone) JSObject() *js.Object { return b.p }

// Bone returns a Bone JavaScript class.
func (t *Three) Bone() *Bone {
	p := t.ctx.Get("Bone")
	return BoneFromJSObject(p)
}

// BoneFromJSObject returns a wrapped Bone JavaScript class.
func BoneFromJSObject(p *js.Object) *Bone {
	return &Bone{p: p}
}

// NewBone returns a new Bone object.
func (t *Three) NewBone(skin float64) *Bone {
	p := t.ctx.Get("Bone").New(skin)
	return BoneFromJSObject(p)
}

// Copy TODO description.
func (b *Bone) Copy(source *Bone) *Bone {
	b.p.Call("copy", source.p)
	return b
}
