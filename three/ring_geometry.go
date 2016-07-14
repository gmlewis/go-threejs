// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// RingGeometry represents two-dimensional ring geometry.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/RingGeometry
type RingGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (r *RingGeometry) JSObject() *js.Object { return r.p }

// RingGeometry returns a RingGeometry JavaScript class.
func (t *Three) RingGeometry() *RingGeometry {
	p := t.ctx.Get("RingGeometry")
	return RingGeometryFromJSObject(p)
}

// RingGeometryFromJSObject returns a wrapped RingGeometry JavaScript class.
func RingGeometryFromJSObject(p *js.Object) *RingGeometry {
	return &RingGeometry{p: p}
}

// NewRingGeometry returns a new RingGeometry object.
//
//     innerRadius — Default is 0, but it doesn't work right when innerRadius is set to 0.
//     outerRadius — Default is 50.
//     thetaSegments — Number of segments. A higher number means the ring will
//         be more round. Minimum is 3. Default is 8.
//     phiSegments — Minimum is 1. Default is 8.
//     thetaStart — Starting angle. Default is 0.
//     thetaLength — Central angle. Default is Math.PI * 2.
func (t *Three) NewRingGeometry(innerRadius, outerRadius float64, thetaSegments, phiSegments int, thetaStart, thetaLength float64) *RingGeometry {
	p := t.ctx.Get("RingGeometry").New(innerRadius, outerRadius, thetaSegments, phiSegments, thetaStart, thetaLength)
	return RingGeometryFromJSObject(p)
}
