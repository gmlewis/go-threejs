// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// EllipseCurve represents an ellipsecurve.
type EllipseCurve struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (e *EllipseCurve) JSObject() *js.Object { return e.p }

// EllipseCurve returns an EllipseCurve object.
func (t *Three) EllipseCurve() *EllipseCurve {
	p := t.ctx.Get("EllipseCurve")
	return &EllipseCurve{p: p}
}

// New returns a new EllipseCurve object.
func (e *EllipseCurve) New(aX, aY, xRadius, yRadius, aStartAngle, aEndAngle, aClockwise, aRotation float64) *EllipseCurve {
	p := e.p.New(aX, aY, xRadius, yRadius, aStartAngle, aEndAngle, aClockwise, aRotation)
	return &EllipseCurve{p: p}
}

// GetPoint TODO description.
func (e *EllipseCurve) GetPoint(t float64) *EllipseCurve {
	e.p.Call("getPoint", t)
	return e
}
