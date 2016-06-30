// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// UniformsUtils represents an uniformsutils.
type UniformsUtils struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (u *UniformsUtils) JSObject() *js.Object { return u.p }

// UniformsUtils returns an UniformsUtils object.
func (t *Three) UniformsUtils() *UniformsUtils {
	p := t.ctx.Get("UniformsUtils")
	return &UniformsUtils{p: p}
}

// New returns a new UniformsUtils object.
func (u *UniformsUtils) New() *UniformsUtils {
	p := u.p.New()
	return &UniformsUtils{p: p}
}

// Merge TODO description.
func (u *UniformsUtils) Merge(uniforms float64) *UniformsUtils {
	u.p.Call("merge", uniforms)
	return u
}

// Clone TODO description.
func (u *UniformsUtils) Clone(uniforms_src float64) *UniformsUtils {
	u.p.Call("clone", uniforms_src)
	return u
}
