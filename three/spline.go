// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Spline represents a spline.
//
// http://threejs.org/docs/index.html#Reference/Math/Spline
type Spline struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (s *Spline) JSObject() *js.Object { return s.p }

// Spline returns a Spline JavaScript class.
func (t *Three) Spline() *Spline {
	p := t.ctx.Get("Spline")
	return &Spline{p: p}
}

// NewSpline returns a new Spline object.
func (t *Three) NewSpline(points float64) *Spline {
	p := t.ctx.Get("Spline").New(points)
	return &Spline{p: p}
}
