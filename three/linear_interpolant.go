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

// LinearInterpolant represents a linearinterpolant.
type LinearInterpolant struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (l *LinearInterpolant) JSObject() *js.Object { return l.p }

// LinearInterpolant returns a LinearInterpolant JavaScript class.
func (t *Three) LinearInterpolant() *LinearInterpolant {
	p := t.ctx.Get("LinearInterpolant")
	return LinearInterpolantFromJSObject(p)
}

// LinearInterpolantFromJSObject returns a wrapped LinearInterpolant JavaScript class.
func LinearInterpolantFromJSObject(p *js.Object) *LinearInterpolant {
	return &LinearInterpolant{p: p}
}

// NewLinearInterpolant returns a new LinearInterpolant object.
func (t *Three) NewLinearInterpolant(parameterPositions, sampleValues, sampleSize, resultBuffer float64) *LinearInterpolant {
	p := t.ctx.Get("LinearInterpolant").New(parameterPositions, sampleValues, sampleSize, resultBuffer)
	return LinearInterpolantFromJSObject(p)
}

// Interpolate TODO description.
func (l *LinearInterpolant) Interpolate(i1, t0, t, t1 float64) *LinearInterpolant {
	l.p.Call("interpolate_", i1, t0, t, t1)
	return l
}
