// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Uniform represents an uniform.
type Uniform struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *Uniform) JSObject() *js.Object { return t.p }

// Uniform returns an Uniform object.
func (t *Three) Uniform() *Uniform {
	p := t.ctx.Get("Uniform")
	return &Uniform{p: p}
}

// New returns a new Uniform object.
func (t *Uniform) New(typ, value float64) *Uniform {
	p := t.p.New(typ, value)
	return &Uniform{p: p}
}

// OnUpdate TODO description.
func (u *Uniform) OnUpdate(callback float64) *Uniform {
	u.p.Call("onUpdate", callback)
	return u
}
