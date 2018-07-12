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

// CubicBezierCurve represents a smooth 2D cubic bezier curve.
//
// http://threejs.org/docs/index.html#Reference/Extras.Curves/CubicBezierCurve
type CubicBezierCurve struct{ *Curve }

// JSObject returns the underlying *js.Object.
func (c *CubicBezierCurve) JSObject() *js.Object { return c.p }

// CubicBezierCurve returns a CubicBezierCurve JavaScript class.
func (t *Three) CubicBezierCurve() *CubicBezierCurve {
	p := t.ctx.Get("CubicBezierCurve")
	return CubicBezierCurveFromJSObject(p)
}

// CubicBezierCurveFromJSObject returns a wrapped CubicBezierCurve JavaScript class.
func CubicBezierCurveFromJSObject(p *js.Object) *CubicBezierCurve {
	return &CubicBezierCurve{CurveFromJSObject(p)}
}

// NewCubicBezierCurve returns a new CubicBezierCurve object.
func (t *Three) NewCubicBezierCurve(v0, v1, v2, v3 *Vector3) *CubicBezierCurve {
	p := t.ctx.Get("CubicBezierCurve").New(v0.p, v1.p, v2.p, v3.p)
	return CubicBezierCurveFromJSObject(p)
}

// GetPoint TODO description.
func (c *CubicBezierCurve) GetPoint(t float64) *CubicBezierCurve {
	c.p.Call("getPoint", t)
	return c
}

// GetTangent TODO description.
func (c *CubicBezierCurve) GetTangent(t float64) *CubicBezierCurve {
	c.p.Call("getTangent", t)
	return c
}
