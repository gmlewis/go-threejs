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

// QuadraticBezierCurve represents a quadraticbeziercurve.
type QuadraticBezierCurve struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (q *QuadraticBezierCurve) JSObject() *js.Object { return q.p }

// QuadraticBezierCurve returns a QuadraticBezierCurve JavaScript class.
func (t *Three) QuadraticBezierCurve() *QuadraticBezierCurve {
	p := t.ctx.Get("QuadraticBezierCurve")
	return QuadraticBezierCurveFromJSObject(p)
}

// QuadraticBezierCurveFromJSObject returns a wrapped QuadraticBezierCurve JavaScript class.
func QuadraticBezierCurveFromJSObject(p *js.Object) *QuadraticBezierCurve {
	return &QuadraticBezierCurve{p: p}
}

// NewQuadraticBezierCurve returns a new QuadraticBezierCurve object.
func (t *Three) NewQuadraticBezierCurve(v0, v1, v2 float64) *QuadraticBezierCurve {
	p := t.ctx.Get("QuadraticBezierCurve").New(v0, v1, v2)
	return QuadraticBezierCurveFromJSObject(p)
}

// GetPoint TODO description.
func (q *QuadraticBezierCurve) GetPoint(t float64) *QuadraticBezierCurve {
	q.p.Call("getPoint", t)
	return q
}

// GetTangent TODO description.
func (q *QuadraticBezierCurve) GetTangent(t float64) *QuadraticBezierCurve {
	q.p.Call("getTangent", t)
	return q
}
