// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// HemisphereLight is a light that is positioned directly above the scene.
//
// http://threejs.org/docs/index.html#Reference/Lights/HemisphereLight
type HemisphereLight struct{ *Light }

// JSObject returns the underlying *js.Object.
func (h *HemisphereLight) JSObject() *js.Object { return h.p }

// HemisphereLight returns a HemisphereLight JavaScript class.
func (t *Three) HemisphereLight() *HemisphereLight {
	return hemisphereLight(t.ctx.Get("HemisphereLight"))
}

// hemisphereLight returns a wrapped HemisphereLight JavaScript class.
func hemisphereLight(p *js.Object) *HemisphereLight {
	return &HemisphereLight{&Light{&Object3D{p: p}}}
}

// NewHemisphereLight returns a new HemisphereLight object.
func (t *Three) NewHemisphereLight(skyColor, groundColor, intensity float64) *HemisphereLight {
	return hemisphereLight(t.ctx.Get("HemisphereLight").New(skyColor, groundColor, intensity))
}

// Copy TODO description.
func (h *HemisphereLight) Copy(source *HemisphereLight) *HemisphereLight {
	h.p.Call("copy", source.p)
	return h
}
