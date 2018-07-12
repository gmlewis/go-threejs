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

// CurvePath represents an array of connected curves,
// but retains the API of a curve.
//
// http://threejs.org/docs/index.html#Reference/Extras.Core/CurvePath
type CurvePath struct{ *Curve }

// JSObject returns the underlying *js.Object.
func (c *CurvePath) JSObject() *js.Object { return c.p }

// CurvePath returns a CurvePath JavaScript class.
func (t *Three) CurvePath() *CurvePath {
	p := t.ctx.Get("CurvePath")
	return CurvePathFromJSObject(p)
}

// CurvePathFromJSObject returns a wrapped CurvePath JavaScript class.
func CurvePathFromJSObject(p *js.Object) *CurvePath {
	return &CurvePath{CurveFromJSObject(p)}
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
func (c *CurvePath) CreatePointsGeometry(divisions float64) *Geometry {
	p := c.p.Call("createPointsGeometry", divisions)
	return GeometryFromJSObject(p)
}

// CreateSpacedPointsGeometry TODO description.
func (c *CurvePath) CreateSpacedPointsGeometry(divisions float64) *Geometry {
	p := c.p.Call("createSpacedPointsGeometry", divisions)
	return GeometryFromJSObject(p)
}

// CreateGeometry TODO description.
func (c *CurvePath) CreateGeometry(points JSObject) *Geometry {
	p := c.p.Call("createGeometry", points.JSObject())
	return GeometryFromJSObject(p)
}
