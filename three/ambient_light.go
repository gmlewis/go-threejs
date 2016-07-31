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
	return AmbientLightFromJSObject(t.ctx.Get("AmbientLight"))
}

// AmbientLightFromJSObject returns a wrapped AmbientLight JavaScript class.
func AmbientLightFromJSObject(p *js.Object) *AmbientLight {
	return &AmbientLight{&Light{&Object3D{p}}}
}

// NewAmbientLight creates an Ambientlight with a color and intensity.
//
//     color — Numeric value of the RGB component of the color.
//     intensity -- Numeric value of the light's strength/intensity.
func (t *Three) NewAmbientLight(color int, intensity float64) *AmbientLight {
	return AmbientLightFromJSObject(t.ctx.Get("AmbientLight").New(color, intensity))
}
