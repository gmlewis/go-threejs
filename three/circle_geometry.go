// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CircleGeometry represents a simple circular shape.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/CircleGeometry
type CircleGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (c *CircleGeometry) JSObject() *js.Object { return c.p }

// CircleGeometry returns a CircleGeometry object.
func (t *Three) CircleGeometry() *CircleGeometry {
	p := t.ctx.Get("CircleGeometry")
	return &CircleGeometry{p: p}
}

// New returns a new CircleGeometry object.
//
//     radius — Radius of the circle, default = 50.
//     segments — Number of segments (triangles), minimum = 3, default = 8.
//     thetaStart — Start angle for first segment, default = 0 (three o'clock position).
//     thetaLength — The central angle, often called theta, of the circular sector.
//         The default is 2*Pi, which makes for a complete circle.
func (c *CircleGeometry) New(radius float64, segments int, thetaStart, thetaLength float64) *CircleGeometry {
	p := c.p.New(radius, segments, thetaStart, thetaLength)
	return &CircleGeometry{p: p}
}
