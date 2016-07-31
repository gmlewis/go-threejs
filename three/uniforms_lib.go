// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// UniformsLib is a library for shared webgl shaders.
//
// http://threejs.org/docs/index.html#Reference/Renderers.Shaders/UniformsLib
type UniformsLib struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (u *UniformsLib) JSObject() *js.Object { return u.p }

// UniformsLib returns an UniformsLib JavaScript class.
func (t *Three) UniformsLib() *UniformsLib {
	p := t.ctx.Get("UniformsLib")
	return UniformsLibFromJSObject(p)
}

// UniformsLibFromJSObject returns a wrapped UniformsLib JavaScript class.
func UniformsLibFromJSObject(p *js.Object) *UniformsLib {
	return &UniformsLib{p: p}
}

// NewUniformsLib returns a new UniformsLib object.
func (t *Three) NewUniformsLib() *UniformsLib {
	p := t.ctx.Get("UniformsLib").New()
	return UniformsLibFromJSObject(p)
}
