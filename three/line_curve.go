// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// LineCurve represents a linecurve.
type LineCurve struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (l *LineCurve) JSObject() *js.Object { return l.p }

// LineCurve returns a LineCurve JavaScript class.
func (t *Three) LineCurve() *LineCurve {
	p := t.ctx.Get("LineCurve")
	return LineCurveFromJSObject(p)
}

// LineCurveFromJSObject returns a wrapped LineCurve JavaScript class.
func LineCurveFromJSObject(p *js.Object) *LineCurve {
	return &LineCurve{p: p}
}

// NewLineCurve returns a new LineCurve object.
func (t *Three) NewLineCurve(v1, v2 float64) *LineCurve {
	p := t.ctx.Get("LineCurve").New(v1, v2)
	return LineCurveFromJSObject(p)
}

// GetPoint TODO description.
func (l *LineCurve) GetPoint(t float64) *LineCurve {
	l.p.Call("getPoint", t)
	return l
}

// GetPointAt TODO description.
func (l *LineCurve) GetPointAt(u float64) *LineCurve {
	l.p.Call("getPointAt", u)
	return l
}

// GetTangent TODO description.
func (l *LineCurve) GetTangent(t float64) *LineCurve {
	l.p.Call("getTangent", t)
	return l
}
