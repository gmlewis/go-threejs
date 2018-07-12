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

// DiscreteInterpolant represents a discreteinterpolant.
type DiscreteInterpolant struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (d *DiscreteInterpolant) JSObject() *js.Object { return d.p }

// DiscreteInterpolant returns a DiscreteInterpolant JavaScript class.
func (t *Three) DiscreteInterpolant() *DiscreteInterpolant {
	p := t.ctx.Get("DiscreteInterpolant")
	return DiscreteInterpolantFromJSObject(p)
}

// DiscreteInterpolantFromJSObject returns a wrapped DiscreteInterpolant JavaScript class.
func DiscreteInterpolantFromJSObject(p *js.Object) *DiscreteInterpolant {
	return &DiscreteInterpolant{p: p}
}

// NewDiscreteInterpolant returns a new DiscreteInterpolant object.
func (t *Three) NewDiscreteInterpolant(parameterPositions, sampleValues, sampleSize, resultBuffer float64) *DiscreteInterpolant {
	p := t.ctx.Get("DiscreteInterpolant").New(parameterPositions, sampleValues, sampleSize, resultBuffer)
	return DiscreteInterpolantFromJSObject(p)
}

// Interpolate TODO description.
func (d *DiscreteInterpolant) Interpolate(i1, t0, t, t1 float64) *DiscreteInterpolant {
	d.p.Call("interpolate_", i1, t0, t, t1)
	return d
}
