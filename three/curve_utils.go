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

// CurveUtils represents a curveutils.
type CurveUtils struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (c *CurveUtils) JSObject() *js.Object { return c.p }

// CurveUtils returns a CurveUtils JavaScript class.
func (t *Three) CurveUtils() *CurveUtils {
	p := t.ctx.Get("CurveUtils")
	return CurveUtilsFromJSObject(p)
}

// CurveUtilsFromJSObject returns a wrapped CurveUtils JavaScript class.
func CurveUtilsFromJSObject(p *js.Object) *CurveUtils {
	return &CurveUtils{p: p}
}

// NewCurveUtils returns a new CurveUtils object.
func (t *Three) NewCurveUtils() *CurveUtils {
	p := t.ctx.Get("CurveUtils").New()
	return CurveUtilsFromJSObject(p)
}

// TangentQuadraticBezier TODO description.
func (c *CurveUtils) TangentQuadraticBezier(t, p0, p1, p2 float64) *CurveUtils {
	c.p.Call("tangentQuadraticBezier", t, p0, p1, p2)
	return c
}

// TangentCubicBezier TODO description.
func (c *CurveUtils) TangentCubicBezier(t, p0, p1, p2, p3 float64) *CurveUtils {
	c.p.Call("tangentCubicBezier", t, p0, p1, p2, p3)
	return c
}

// TangentSpline TODO description.
func (c *CurveUtils) TangentSpline(t, p0, p1, p2, p3 float64) *CurveUtils {
	c.p.Call("tangentSpline", t, p0, p1, p2, p3)
	return c
}

// Interpolate TODO description.
func (c *CurveUtils) Interpolate(p0, p1, p2, p3, t float64) *CurveUtils {
	c.p.Call("interpolate", p0, p1, p2, p3, t)
	return c
}
