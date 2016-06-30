// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Spline represents a spline.
type Spline struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (s *Spline) JSObject() *js.Object { return s.p }

// Spline returns a Spline object.
func (t *Three) Spline() *Spline {
	p := t.ctx.Get("Spline")
	return &Spline{p: p}
}

// New returns a new Spline object.
func (s *Spline) New(points float64) *Spline {
	p := s.p.New(points)
	return &Spline{p: p}
}
