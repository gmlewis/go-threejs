// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// ArcCurve represents an arc curve which is an EllipseCurve.
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
