// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// LightShadow represents a lightshadow.
type LightShadow struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (l *LightShadow) JSObject() *js.Object { return l.p }

// LightShadow returns a LightShadow object.
func (t *Three) LightShadow() *LightShadow {
	p := t.ctx.Get("LightShadow")
	return &LightShadow{p: p}
}

// New returns a new LightShadow object.
func (l *LightShadow) New(camera float64) *LightShadow {
	p := l.p.New(camera)
	return &LightShadow{p: p}
}

// Copy TODO description.
func (l *LightShadow) Copy(source *LightShadow) *LightShadow {
	l.p.Call("copy", source.p)
	return l
}

// Clone TODO description.
func (l *LightShadow) Clone() *LightShadow {
	l.p.Call("clone")
	return l
}
