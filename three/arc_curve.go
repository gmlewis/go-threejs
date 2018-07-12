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

// ArcCurve represents an arc curve which is an EllipseCurve.
//
// http://threejs.org/docs/index.html#Reference/Extras.Curves/ArcCurve
type ArcCurve struct{ *EllipseCurve }

// JSObject returns the underlying *js.Object.
func (a *ArcCurve) JSObject() *js.Object { return a.p }

// ArcCurve returns an ArcCurve JavaScript class.
func (t *Three) ArcCurve() *ArcCurve {
	p := t.ctx.Get("ArcCurve")
	return ArcCurveFromJSObject(p)
}

// ArcCurveFromJSObject returns a wrapped ArcCurve JavaScript class.
func ArcCurveFromJSObject(p *js.Object) *ArcCurve {
	return &ArcCurve{EllipseCurveFromJSObject(p)}
}

// NewArcCurve returns a new ArcCurve object.
func (t *Three) NewArcCurve(aX, aY, aRadius, aStartAngle, aEndAngle, aClockwise float64) *ArcCurve {
	p := t.ctx.Get("ArcCurve").New(aX, aY, aRadius, aStartAngle, aEndAngle, aClockwise)
	return ArcCurveFromJSObject(p)
}
