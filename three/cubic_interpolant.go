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

// CubicInterpolant represents a cubicinterpolant.
type CubicInterpolant struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (c *CubicInterpolant) JSObject() *js.Object { return c.p }

// CubicInterpolant returns a CubicInterpolant JavaScript class.
func (t *Three) CubicInterpolant() *CubicInterpolant {
	p := t.ctx.Get("CubicInterpolant")
	return CubicInterpolantFromJSObject(p)
}

// CubicInterpolantFromJSObject returns a wrapped CubicInterpolant JavaScript class.
func CubicInterpolantFromJSObject(p *js.Object) *CubicInterpolant {
	return &CubicInterpolant{p: p}
}

// NewCubicInterpolant returns a new CubicInterpolant object.
func (t *Three) NewCubicInterpolant(parameterPositions, sampleValues, sampleSize, resultBuffer float64) *CubicInterpolant {
	p := t.ctx.Get("CubicInterpolant").New(parameterPositions, sampleValues, sampleSize, resultBuffer)
	return CubicInterpolantFromJSObject(p)
}

// IntervalChanged TODO description.
func (c *CubicInterpolant) IntervalChanged(i1, t0, t1 float64) *CubicInterpolant {
	c.p.Call("intervalChanged_", i1, t0, t1)
	return c
}

// Interpolate TODO description.
func (c *CubicInterpolant) Interpolate(i1, t0, t, t1 float64) *CubicInterpolant {
	c.p.Call("interpolate_", i1, t0, t, t1)
	return c
}
