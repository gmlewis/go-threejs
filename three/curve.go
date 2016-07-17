// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Curve represents a curve.
type Curve struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (c *Curve) JSObject() *js.Object { return c.p }

// Curve returns a Curve JavaScript class.
func (t *Three) Curve() *Curve {
	p := t.ctx.Get("Curve")
	return CurveFromJSObject(p)
}

// CurveFromJSObject returns a wrapped Curve JavaScript class.
func CurveFromJSObject(p *js.Object) *Curve {
	return &Curve{p: p}
}

// NewCurve returns a new Curve object.
func (t *Three) NewCurve() *Curve {
	p := t.ctx.Get("Curve").New()
	return CurveFromJSObject(p)
}

// GetPoint TODO description.
func (c *Curve) GetPoint(t int) *js.Object {
	return c.p.Call("getPoint", t)
}

// GetPointAt TODO description.
func (c *Curve) GetPointAt(u int) *js.Object {
	return c.p.Call("getPointAt", u)
}

// GetPoints TODO description.
func (c *Curve) GetPoints(divisions int) *js.Object {
	return c.p.Call("getPoints", divisions)
}

// GetSpacedPoints TODO description.
func (c *Curve) GetSpacedPoints(divisions int) *js.Object {
	return c.p.Call("getSpacedPoints", divisions)
}

// GetLength TODO description.
func (c *Curve) GetLength() *Curve {
	c.p.Call("getLength")
	return c
}

// GetLengths TODO description.
func (c *Curve) GetLengths(divisions float64) *Curve {
	c.p.Call("getLengths", divisions)
	return c
}

// UpdateArcLengths TODO description.
func (c *Curve) UpdateArcLengths() *Curve {
	c.p.Call("updateArcLengths")
	return c
}

// GetUtoTmapping TODO description.
func (c *Curve) GetUtoTmapping(u, distance float64) *Curve {
	c.p.Call("getUtoTmapping", u, distance)
	return c
}

// GetTangent TODO description.
func (c *Curve) GetTangent(t float64) *Curve {
	c.p.Call("getTangent", t)
	return c
}

// GetTangentAt TODO description.
func (c *Curve) GetTangentAt(u float64) *Curve {
	c.p.Call("getTangentAt", u)
	return c
}
