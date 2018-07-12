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
