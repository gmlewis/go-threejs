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

// UniformsUtils returns an UniformsUtils JavaScript class.
func (t *Three) UniformsUtils() *UniformsUtils {
	p := t.ctx.Get("UniformsUtils")
	return &UniformsUtils{p: p}
}

// NewUniformsUtils returns a new UniformsUtils object.
func (t *Three) NewUniformsUtils() *UniformsUtils {
	p := t.ctx.Get("UniformsUtils").New()
	return &UniformsUtils{p: p}
}

// Merge TODO description.
func (u *UniformsUtils) Merge(uniforms float64) *UniformsUtils {
	u.p.Call("merge", uniforms)
	return u
}

// Clone TODO description.
func (u *UniformsUtils) Clone(uniformsSrc JSObject) *UniformsUtils {
	u.p.Call("clone", uniformsSrc.JSObject())
	return u
}
