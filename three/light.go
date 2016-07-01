// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Light represents a light.
type Light struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (l *Light) JSObject() *js.Object { return l.p }

// Light returns a Light object.
func (t *Three) Light() *Light {
	p := t.ctx.Get("Light")
	return &Light{p: p}
}

// New returns a new Light object.
func (l *Light) New(color, intensity float64) *Light {
	p := l.p.New(color, intensity)
	return &Light{p: p}
}

// Copy TODO description.
func (l *Light) Copy(source *Light) *Light {
	l.p.Call("copy", source.p)
	return l
}

// ToJSON TODO description.
func (l *Light) ToJSON(meta float64) *Light {
	l.p.Call("toJSON", meta)
	return l
}
