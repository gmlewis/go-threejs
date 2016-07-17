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
type CircleGeometry struct{ *Geometry }

// JSObject returns the underlying *js.Object.
func (c *CircleGeometry) JSObject() *js.Object { return c.p }

// CircleGeometry returns a CircleGeometry JavaScript class.
func (t *Three) CircleGeometry() *CircleGeometry {
	p := t.ctx.Get("CircleGeometry")
	return CircleGeometryFromJSObject(p)
}

// CircleGeometryFromJSObject returns a wrapped CircleGeometry JavaScript class.
func CircleGeometryFromJSObject(p *js.Object) *CircleGeometry {
	return &CircleGeometry{GeometryFromJSObject(p)}
}

// CircleGeometryOpts represents options for NewCircleGeometry.
//
//     thetaStart — Start angle for first segment, default = 0 (three o'clock position).
//     thetaLength — The central angle, often called theta, of the circular sector.
//         The default is 2*Pi, which makes for a complete circle.
type CircleGeometryOpts struct {
	thetaStart  float64
	thetaLength float64
}

// NewCircleGeometry returns a new CircleGeometry object.
//
//     radius — Radius of the circle, default = 50.
//     segments — Number of segments (triangles), minimum = 3, default = 8.
func (t *Three) NewCircleGeometry(radius float64, segments int, opts *CircleGeometryOpts) *CircleGeometry {
	var p *js.Object
	if opts != nil {
		p = t.ctx.Get("CircleGeometry").New(radius, segments, opts.thetaStart, opts.thetaLength)
	} else {
		p = t.ctx.Get("CircleGeometry").New(radius, segments)
	}
	return CircleGeometryFromJSObject(p)
}
