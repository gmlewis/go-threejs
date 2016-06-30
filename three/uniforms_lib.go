// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// UniformsLib represents an uniformslib.
type UniformsLib struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *UniformsLib) JSObject() *js.Object { return t.p }

// UniformsLib returns an UniformsLib object.
func (t *Three) UniformsLib() *UniformsLib {
	p := t.ctx.Get("UniformsLib")
	return &UniformsLib{p: p}
}

// New returns a new UniformsLib object.
func (t *UniformsLib) New() *UniformsLib {
	p := t.p.New()
	return &UniformsLib{p: p}
}
