// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Fog represents linear fog.
//
// http://threejs.org/docs/index.html#Reference/Scenes/Fog
type Fog struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (f *Fog) JSObject() *js.Object { return f.p }

// Fog returns a Fog object.
func (t *Three) Fog() *Fog {
	p := t.ctx.Get("Fog")
	return &Fog{p: p}
}

// New returns a new Fog object.
func (f *Fog) New(color, near, far float64) *Fog {
	p := f.p.New(color, near, far)
	return &Fog{p: p}
}
