// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// DirectionalLight represents a directionallight.
type DirectionalLight struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (d *DirectionalLight) JSObject() *js.Object { return d.p }

// DirectionalLight returns a DirectionalLight JavaScript class.
func (t *Three) DirectionalLight() *DirectionalLight {
	p := t.ctx.Get("DirectionalLight")
	return &DirectionalLight{p: p}
}

// NewDirectionalLight returns a new DirectionalLight object.
func (t *Three) NewDirectionalLight(color, intensity float64) *DirectionalLight {
	p := t.ctx.Get("DirectionalLight").New(color, intensity)
	return &DirectionalLight{p: p}
}

// Copy TODO description.
func (d *DirectionalLight) Copy(source *DirectionalLight) *DirectionalLight {
	d.p.Call("copy", source.p)
	return d
}
