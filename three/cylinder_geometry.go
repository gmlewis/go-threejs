// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CylinderGeometry represents a cylinder geometry primitive.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/CylinderGeometry
type CylinderGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (c *CylinderGeometry) JSObject() *js.Object { return c.p }

// CylinderGeometry returns a CylinderGeometry object.
func (t *Three) CylinderGeometry() *CylinderGeometry {
	p := t.ctx.Get("CylinderGeometry")
	return &CylinderGeometry{p: p}
}

// New returns a new CylinderGeometry object.
//
//     radiusTop — Radius of the cylinder at the top. Default is 20.
//     radiusBottom — Radius of the cylinder at the bottom. Default is 20.
//     height — Height of the cylinder. Default is 100.
//     radiusSegments — Number of segmented faces around the circumference of the cylinder. Default is 8
//     heightSegments — Number of rows of faces along the height of the cylinder. Default is 1.
//     openEnded — A Boolean indicating whether the ends of the cylinder are open or capped.
//         Default is false, meaning capped.
//     thetaStart — Start angle for first segment, default = 0 (three o'clock position).
//     thetaLength — The central angle, often called theta, of the circular sector.
//         The default is 2*Pi, which makes for a complete cylinder.
func (c *CylinderGeometry) New(radiusTop, radiusBottom, height float64, radialSegments, heightSegments int, openEnded bool, thetaStart, thetaLength float64) *CylinderGeometry {
	p := c.p.New(radiusTop, radiusBottom, height, radialSegments, heightSegments, openEnded, thetaStart, thetaLength)
	return &CylinderGeometry{p: p}
}
