// Copyright 2016 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	return SplineCurveFromJSObject(p)
}

// SplineCurveFromJSObject returns a wrapped SplineCurve JavaScript class.
func SplineCurveFromJSObject(p *js.Object) *SplineCurve {
	return &SplineCurve{p: p}
}

// NewSplineCurve returns a new SplineCurve object.
func (t *Three) NewSplineCurve(points /* array of Vector2 */ float64) *SplineCurve {
	p := t.ctx.Get("SplineCurve").New(points /* array of Vector2 */)
	return SplineCurveFromJSObject(p)
}

// GetPoint TODO description.
func (s *SplineCurve) GetPoint(t float64) *SplineCurve {
	s.p.Call("getPoint", t)
	return s
}
