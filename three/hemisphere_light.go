// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// HemisphereLight represents a hemispherelight.
type HemisphereLight struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *HemisphereLight) JSObject() *js.Object { return t.p }

// HemisphereLight returns a HemisphereLight object.
func (t *Three) HemisphereLight() *HemisphereLight {
	p := t.ctx.Get("HemisphereLight")
	return &HemisphereLight{p: p}
}

// New returns a new HemisphereLight object.
func (t *HemisphereLight) New(skyColor, groundColor, intensity float64) *HemisphereLight {
	p := t.p.New(skyColor, groundColor, intensity)
	return &HemisphereLight{p: p}
}

// Copy TODO description.
func (h *HemisphereLight) Copy(source float64) *HemisphereLight {
	h.p.Call("copy", source)
	return h
}