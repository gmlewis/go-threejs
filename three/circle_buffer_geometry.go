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

// CircleBufferGeometry is the BufferGeometry port of CircleGeometry.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/CircleBufferGeometry
type CircleBufferGeometry struct{ *BufferGeometry }

// JSObject returns the underlying *js.Object.
func (c *CircleBufferGeometry) JSObject() *js.Object { return c.p }

// CircleBufferGeometry returns a CircleBufferGeometry JavaScript class.
func (t *Three) CircleBufferGeometry() *CircleBufferGeometry {
	p := t.ctx.Get("CircleBufferGeometry")
	return CircleBufferGeometryFromJSObject(p)
}

// CircleBufferGeometryFromJSObject returns a wrapped CircleBufferGeometry JavaScript class.
func CircleBufferGeometryFromJSObject(p *js.Object) *CircleBufferGeometry {
	return &CircleBufferGeometry{BufferGeometryFromJSObject(p)}
}

// CircleBufferGeometryOpts represents options for NewCircleBufferGeometry.
//
//     thetaStart — Start angle for first segment, default = 0 (three o'clock position).
//     thetaLength — The central angle, often called theta, of the circular sector. The default is 2*Pi, which makes for a complete circle.
type CircleBufferGeometryOpts struct {
	thetaStart  float64
	thetaLength float64
}

// NewCircleBufferGeometry returns a new CircleBufferGeometry object.
//
//     radius — Radius of the circle, default = 50.
//     segments — Number of segments (triangles), minimum = 3, default = 8.
func (t *Three) NewCircleBufferGeometry(radius float64, segments int, opts *CircleBufferGeometryOpts) *CircleBufferGeometry {
	var p *js.Object
	if opts != nil {
		p = t.ctx.Get("CircleBufferGeometry").New(radius, segments, opts.thetaStart, opts.thetaLength)
	} else {
		p = t.ctx.Get("CircleBufferGeometry").New(radius, segments)
	}
	return CircleBufferGeometryFromJSObject(p)
}
