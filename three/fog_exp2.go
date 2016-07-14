// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// FogExp2 represents exponential fog.
//
// http://threejs.org/docs/index.html#Reference/Scenes/FogExp2
type FogExp2 struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (f *FogExp2) JSObject() *js.Object { return f.p }

// FogExp2 returns a FogExp2 JavaScript class.
func (t *Three) FogExp2() *FogExp2 {
	p := t.ctx.Get("FogExp2")
	return FogExp2FromJSObject(p)
}

// FogExp2FromJSObject returns a wrapped FogExp2 JavaScript class.
func FogExp2FromJSObject(p *js.Object) *FogExp2 {
	return &FogExp2{p: p}
}

// NewFogExp2 returns a new FogExp2 object.
func (t *Three) NewFogExp2(color, density float64) *FogExp2 {
	p := t.ctx.Get("FogExp2").New(color, density)
	return FogExp2FromJSObject(p)
}
