// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Spherical represents a spherical.
type Spherical struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (s *Spherical) JSObject() *js.Object { return s.p }

// Spherical returns a Spherical object.
func (t *Three) Spherical() *Spherical {
	p := t.ctx.Get("Spherical")
	return &Spherical{p: p}
}

// New returns a new Spherical object.
func (s *Spherical) New(radius, phi, theta float64) *Spherical {
	p := s.p.New(radius, phi, theta)
	return &Spherical{p: p}
}

// Set TODO description.
func (s *Spherical) Set(radius, phi, theta float64) *Spherical {
	s.p.Call("set", radius, phi, theta)
	return s
}

// Clone TODO description.
func (s *Spherical) Clone() *Spherical {
	s.p.Call("clone")
	return s
}

// Copy TODO description.
func (s *Spherical) Copy(other float64) *Spherical {
	s.p.Call("copy", other)
	return s
}

// MakeSafe TODO description.
func (s *Spherical) MakeSafe() *Spherical {
	s.p.Call("makeSafe")
	return s
}

// SetFromVector3 TODO description.
func (s *Spherical) SetFromVector3(vec3 float64) *Spherical {
	s.p.Call("setFromVector3", vec3)
	return s
}
