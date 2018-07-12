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

// Curve represents a a curve object that contains methods
// for interpolation. For an array of Curves see CurvePath.
//
// http://threejs.org/docs/index.html#Reference/Extras.Core/Curve
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
