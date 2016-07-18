// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CylinderBufferGeometry represents a cylinderbuffergeometry.
type CylinderBufferGeometry struct{ *BufferGeometry }

// JSObject returns the underlying *js.Object.
func (c *CylinderBufferGeometry) JSObject() *js.Object { return c.p }

// CylinderBufferGeometry returns a CylinderBufferGeometry JavaScript class.
func (t *Three) CylinderBufferGeometry() *CylinderBufferGeometry {
	p := t.ctx.Get("CylinderBufferGeometry")
	return CylinderBufferGeometryFromJSObject(p)
}

// CylinderBufferGeometryFromJSObject returns a wrapped CylinderBufferGeometry JavaScript class.
func CylinderBufferGeometryFromJSObject(p *js.Object) *CylinderBufferGeometry {
	return &CylinderBufferGeometry{BufferGeometryFromJSObject(p)}
}

// CylinderBufferGeometryOpts represents options for NewCylinderBufferGeometry.
//
// thetaStart — Start angle for first segment, default = 0 (three o'clock position).
// thetaLength — The central angle, often called theta, of the circular sector. The default is 2*Pi, which makes for a complete cylinder.
type CylinderBufferGeometryOpts struct {
	thetaStart  float64
	thetaLength float64
}

// NewCylinderBufferGeometry returns a new CylinderBufferGeometry object.
//
// radiusTop — Radius of the cylinder at the top. Default is 20.
// radiusBottom — Radius of the cylinder at the bottom. Default is 20.
// height — Height of the cylinder. Default is 100.
// radiusSegments — Number of segmented faces around the circumference of the cylinder. Default is 8
// heightSegments — Number of rows of faces along the height of the cylinder. Default is 1.
// openEnded — A Boolean indicating whether the ends of the cylinder are open or capped. Default is false, meaning capped.
func (t *Three) NewCylinderBufferGeometry(radiusTop, radiusBottom, height float64, radialSegments, heightSegments int, openEnded bool, opts *CylinderBufferGeometryOpts) *CylinderBufferGeometry {
	var p *js.Object
	if opts != nil {
		p = t.ctx.Get("CylinderBufferGeometry").New(radiusTop, radiusBottom, height, radialSegments, heightSegments, openEnded, opts.thetaStart, opts.thetaLength)
	} else {
		p = t.ctx.Get("CylinderBufferGeometry").New(radiusTop, radiusBottom, height, radialSegments, heightSegments, openEnded)
	}
	return CylinderBufferGeometryFromJSObject(p)
}
