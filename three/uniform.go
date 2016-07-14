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
func (u *Uniform) JSObject() *js.Object { return u.p }

// Uniform returns an Uniform JavaScript class.
func (t *Three) Uniform() *Uniform {
	p := t.ctx.Get("Uniform")
	return UniformFromJSObject(p)
}

// UniformFromJSObject returns a wrapped Uniform JavaScript class.
func UniformFromJSObject(p *js.Object) *Uniform {
	return &Uniform{p: p}
}

// NewUniform returns a new Uniform object.
func (t *Three) NewUniform(typ, value float64) *Uniform {
	p := t.ctx.Get("Uniform").New(typ, value)
	return UniformFromJSObject(p)
}

// OnUpdate TODO description.
func (u *Uniform) OnUpdate(callback float64) *Uniform {
	u.p.Call("onUpdate", callback)
	return u
}
