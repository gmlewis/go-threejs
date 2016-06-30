// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// AmbientLight represents an ambientlight.
type AmbientLight struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *AmbientLight) JSObject() *js.Object { return t.p }

// AmbientLight returns an AmbientLight object.
func (t *Three) AmbientLight() *AmbientLight {
	p := t.ctx.Get("AmbientLight")
	return &AmbientLight{p: p}
}

// New returns a new AmbientLight object.
func (t *AmbientLight) New(color, intensity float64) *AmbientLight {
	p := t.p.New(color, intensity)
	return &AmbientLight{p: p}
}
