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

// Spline represents a spline.
//
// http://threejs.org/docs/index.html#Reference/Math/Spline
type Spline struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (s *Spline) JSObject() *js.Object { return s.p }

// Spline returns a Spline JavaScript class.
func (t *Three) Spline() *Spline {
	p := t.ctx.Get("Spline")
	return SplineFromJSObject(p)
}

// SplineFromJSObject returns a wrapped Spline JavaScript class.
func SplineFromJSObject(p *js.Object) *Spline {
	return &Spline{p: p}
}

// NewSpline returns a new Spline object.
func (t *Three) NewSpline(points float64) *Spline {
	p := t.ctx.Get("Spline").New(points)
	return SplineFromJSObject(p)
}
