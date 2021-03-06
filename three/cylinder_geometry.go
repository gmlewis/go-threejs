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

// CylinderGeometry represents a cylinder geometry primitive.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/CylinderGeometry
type CylinderGeometry struct{ *Geometry }

// JSObject returns the underlying *js.Object.
func (c *CylinderGeometry) JSObject() *js.Object { return c.p }

// CylinderGeometry returns a CylinderGeometry JavaScript class.
func (t *Three) CylinderGeometry() *CylinderGeometry {
	p := t.ctx.Get("CylinderGeometry")
	return CylinderGeometryFromJSObject(p)
}

// CylinderGeometryFromJSObject returns a wrapped CylinderGeometry JavaScript class.
func CylinderGeometryFromJSObject(p *js.Object) *CylinderGeometry {
	return &CylinderGeometry{GeometryFromJSObject(p)}
}

// CylinderGeometryOpts represents options for NewCylinderGeometry.
//
//     thetaStart — Start angle for first segment, default = 0 (three o'clock position).
//     thetaLength — The central angle, often called theta, of the circular sector.
//         The default is 2*Pi, which makes for a complete cylinder.
type CylinderGeometryOpts struct {
	thetaStart  float64
	thetaLength float64
}

// NewCylinderGeometry returns a new CylinderGeometry object.
//
//     radiusTop — Radius of the cylinder at the top. Default is 20.
//     radiusBottom — Radius of the cylinder at the bottom. Default is 20.
//     height — Height of the cylinder. Default is 100.
//     radiusSegments — Number of segmented faces around the circumference of the cylinder. Default is 8
//     heightSegments — Number of rows of faces along the height of the cylinder. Default is 1.
//     openEnded — A Boolean indicating whether the ends of the cylinder are open or capped.
//         Default is false, meaning capped.
func (t *Three) NewCylinderGeometry(radiusTop, radiusBottom, height float64, radialSegments, heightSegments int, openEnded bool, opts *CylinderGeometryOpts) *CylinderGeometry {
	var p *js.Object
	if opts != nil {
		p = t.ctx.Get("CylinderGeometry").New(radiusTop, radiusBottom, height, radialSegments, heightSegments, openEnded, opts.thetaStart, opts.thetaLength)
	} else {
		p = t.ctx.Get("CylinderGeometry").New(radiusTop, radiusBottom, height, radialSegments, heightSegments, openEnded)
	}
	return CylinderGeometryFromJSObject(p)
}
