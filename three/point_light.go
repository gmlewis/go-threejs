// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// PointLight represents a pointlight.
type PointLight struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (p *PointLight) JSObject() *js.Object { return p.p }

// PointLight returns a PointLight JavaScript class.
func (t *Three) PointLight() *PointLight {
	p := t.ctx.Get("PointLight")
	return &PointLight{p: p}
}

// NewPointLight returns a new PointLight object.
func (t *Three) NewPointLight(color, intensity, distance, decay float64) *PointLight {
	p := t.ctx.Get("PointLight").New(color, intensity, distance, decay)
	return &PointLight{p: p}
}

// Get TODO description.
func (p *PointLight) Get() *PointLight {
	p.p.Call("get")
	return p
}

// Set TODO description.
func (p *PointLight) Set(power float64) *PointLight {
	p.p.Call("set", power)
	return p
}

// Copy TODO description.
func (p *PointLight) Copy(source *PointLight) *PointLight {
	p.p.Call("copy", source.p)
	return p
}
