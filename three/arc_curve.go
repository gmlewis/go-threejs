// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// ArcCurve represents an arccurve.
type ArcCurve struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *ArcCurve) JSObject() *js.Object { return t.p }

// ArcCurve returns an ArcCurve object.
func (t *Three) ArcCurve() *ArcCurve {
	p := t.ctx.Get("ArcCurve")
	return &ArcCurve{p: p}
}

// New returns a new ArcCurve object.
func (t *ArcCurve) New(aX, aY, aRadius, aStartAngle, aEndAngle, aClockwise float64) *ArcCurve {
	p := t.p.New(aX, aY, aRadius, aStartAngle, aEndAngle, aClockwise)
	return &ArcCurve{p: p}
}
