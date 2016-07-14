// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CurvePath represents a curvepath.
type CurvePath struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (c *CurvePath) JSObject() *js.Object { return c.p }

// CurvePath returns a CurvePath JavaScript class.
func (t *Three) CurvePath() *CurvePath {
	p := t.ctx.Get("CurvePath")
	return CurvePathFromJSObject(p)
}

// CurvePathFromJSObject returns a wrapped CurvePath JavaScript class.
func CurvePathFromJSObject(p *js.Object) *CurvePath {
	return &CurvePath{p: p}
}

// NewCurvePath returns a new CurvePath object.
func (t *Three) NewCurvePath() *CurvePath {
	p := t.ctx.Get("CurvePath").New()
	return CurvePathFromJSObject(p)
}

// Add TODO description.
func (c *CurvePath) Add(curve float64) *CurvePath {
	c.p.Call("add", curve)
	return c
}

// GetPoint TODO description.
func (c *CurvePath) GetPoint(t float64) *CurvePath {
	c.p.Call("getPoint", t)
	return c
}

// GetTangent TODO description.
func (c *CurvePath) GetTangent(t float64) *CurvePath {
	c.p.Call("getTangent", t)
	return c
}

// CreatePointsGeometry TODO description.
func (c *CurvePath) CreatePointsGeometry(divisions float64) *CurvePath {
	c.p.Call("createPointsGeometry", divisions)
	return c
}

// CreateSpacedPointsGeometry TODO description.
func (c *CurvePath) CreateSpacedPointsGeometry(divisions float64) *CurvePath {
	c.p.Call("createSpacedPointsGeometry", divisions)
	return c
}

// CreateGeometry TODO description.
func (c *CurvePath) CreateGeometry(points float64) *CurvePath {
	c.p.Call("createGeometry", points)
	return c
}
