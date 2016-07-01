// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// AmbientLight is a color light that gets applied to all the objects in the
// scene globally.
//
// http://threejs.org/docs/index.html#Reference/Lights/AmbientLight
type AmbientLight struct{ *Light }

// JSObject returns the underlying *js.Object.
func (a *AmbientLight) JSObject() *js.Object { return a.p }

// AmbientLight returns an AmbientLight JavaScript class.
func (t *Three) AmbientLight() *AmbientLight {
	p := t.ctx.Get("AmbientLight")
	return &AmbientLight{&Light{&Object3D{p: p}}}
}

// NewAmbientLight returns a new AmbientLight object.
func (t *Three) NewAmbientLight(color int, intensity float64) *AmbientLight {
	p := t.ctx.Get("AmbientLight").New(color, intensity)
	return &AmbientLight{&Light{&Object3D{p: p}}}
}
