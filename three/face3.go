// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Face3 represents a face3.
type Face3 struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (f *Face3) JSObject() *js.Object { return f.p }

// Face3 returns a Face3 object.
func (t *Three) Face3() *Face3 {
	p := t.ctx.Get("Face3")
	return &Face3{p: p}
}

// New returns a new Face3 object.
func (f *Face3) New(a, b, c, normal, color, materialIndex float64) *Face3 {
	p := f.p.New(a, b, c, normal, color, materialIndex)
	return &Face3{p: p}
}

// Clone TODO description.
func (f *Face3) Clone() *Face3 {
	f.p.Call("clone")
	return f
}

// Copy TODO description.
func (f *Face3) Copy(source float64) *Face3 {
	f.p.Call("copy", source)
	return f
}
