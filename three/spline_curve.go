// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// SplineCurve creates a smooth 2d spline curve from a series of points.
//
// http://threejs.org/docs/index.html#Reference/Extras.Curves/SplineCurve
type SplineCurve struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (s *SplineCurve) JSObject() *js.Object { return s.p }

// SplineCurve returns a SplineCurve JavaScript class.
func (t *Three) SplineCurve() *SplineCurve {
	p := t.ctx.Get("SplineCurve")
	return &SplineCurve{p: p}
}

// NewSplineCurve returns a new SplineCurve object.
func (t *Three) NewSplineCurve(points /* array of Vector2 */ float64) *SplineCurve {
	p := t.ctx.Get("SplineCurve").New(points /* array of Vector2 */)
	return &SplineCurve{p: p}
}

// GetPoint TODO description.
func (s *SplineCurve) GetPoint(t float64) *SplineCurve {
	s.p.Call("getPoint", t)
	return s
}
