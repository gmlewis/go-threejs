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

// QuaternionLinearInterpolant represents a quaternionlinearinterpolant.
type QuaternionLinearInterpolant struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (q *QuaternionLinearInterpolant) JSObject() *js.Object { return q.p }

// QuaternionLinearInterpolant returns a QuaternionLinearInterpolant JavaScript class.
func (t *Three) QuaternionLinearInterpolant() *QuaternionLinearInterpolant {
	p := t.ctx.Get("QuaternionLinearInterpolant")
	return QuaternionLinearInterpolantFromJSObject(p)
}

// QuaternionLinearInterpolantFromJSObject returns a wrapped QuaternionLinearInterpolant JavaScript class.
func QuaternionLinearInterpolantFromJSObject(p *js.Object) *QuaternionLinearInterpolant {
	return &QuaternionLinearInterpolant{p: p}
}

// NewQuaternionLinearInterpolant returns a new QuaternionLinearInterpolant object.
func (t *Three) NewQuaternionLinearInterpolant(parameterPositions, sampleValues, sampleSize, resultBuffer float64) *QuaternionLinearInterpolant {
	p := t.ctx.Get("QuaternionLinearInterpolant").New(parameterPositions, sampleValues, sampleSize, resultBuffer)
	return QuaternionLinearInterpolantFromJSObject(p)
}

// Interpolate TODO description.
func (q *QuaternionLinearInterpolant) Interpolate(i1, t0, t, t1 float64) *QuaternionLinearInterpolant {
	q.p.Call("interpolate_", i1, t0, t, t1)
	return q
}
