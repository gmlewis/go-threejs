// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// SpotLight represents a spotlight.
type SpotLight struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (s *SpotLight) JSObject() *js.Object { return s.p }

// SpotLight returns a SpotLight JavaScript class.
func (t *Three) SpotLight() *SpotLight {
	p := t.ctx.Get("SpotLight")
	return &SpotLight{p: p}
}

// NewSpotLight returns a new SpotLight object.
func (t *Three) NewSpotLight(color, intensity, distance, angle, penumbra, decay float64) *SpotLight {
	p := t.ctx.Get("SpotLight").New(color, intensity, distance, angle, penumbra, decay)
	return &SpotLight{p: p}
}

// Get TODO description.
func (s *SpotLight) Get() *SpotLight {
	s.p.Call("get")
	return s
}

// Set TODO description.
func (s *SpotLight) Set(power float64) *SpotLight {
	s.p.Call("set", power)
	return s
}

// Copy TODO description.
func (s *SpotLight) Copy(source *SpotLight) *SpotLight {
	s.p.Call("copy", source.p)
	return s
}
