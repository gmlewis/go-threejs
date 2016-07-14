// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CylinderBufferGeometry represents a cylinderbuffergeometry.
type CylinderBufferGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (c *CylinderBufferGeometry) JSObject() *js.Object { return c.p }

// CylinderBufferGeometry returns a CylinderBufferGeometry JavaScript class.
func (t *Three) CylinderBufferGeometry() *CylinderBufferGeometry {
	p := t.ctx.Get("CylinderBufferGeometry")
	return CylinderBufferGeometryFromJSObject(p)
}

// CylinderBufferGeometryFromJSObject returns a wrapped CylinderBufferGeometry JavaScript class.
func CylinderBufferGeometryFromJSObject(p *js.Object) *CylinderBufferGeometry {
	return &CylinderBufferGeometry{p: p}
}

// NewCylinderBufferGeometry returns a new CylinderBufferGeometry object.
func (t *Three) NewCylinderBufferGeometry(radiusTop, radiusBottom, height, radialSegments, heightSegments, openEnded, thetaStart, thetaLength float64) *CylinderBufferGeometry {
	p := t.ctx.Get("CylinderBufferGeometry").New(radiusTop, radiusBottom, height, radialSegments, heightSegments, openEnded, thetaStart, thetaLength)
	return CylinderBufferGeometryFromJSObject(p)
}
