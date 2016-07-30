// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// RingBufferGeometry is the BufferGeometry port of RingGeometry.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/RingBufferGeometry
type RingBufferGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (r *RingBufferGeometry) JSObject() *js.Object { return r.p }

// RingBufferGeometry returns a RingBufferGeometry JavaScript class.
func (t *Three) RingBufferGeometry() *RingBufferGeometry {
	p := t.ctx.Get("RingBufferGeometry")
	return &RingBufferGeometry{p: p}
}

// NewRingBufferGeometry returns a new RingBufferGeometry object.
//
//     innerRadius — Default is 0, but it doesn't work right when innerRadius is set to 0.
//     outerRadius — Default is 50.
//     thetaSegments — Number of segments. A higher number means the ring will be more round. Minimum is 3. Default is 8.
//     phiSegments — Minimum is 1. Default is 8.
//     thetaStart — Starting angle. Default is 0.
//     thetaLength — Central angle. Default is Math.PI * 2.
func (t *Three) NewRingBufferGeometry(innerRadius, outerRadius, thetaSegments, phiSegments, thetaStart, thetaLength float64) *RingBufferGeometry {
	p := t.ctx.Get("RingBufferGeometry").New(innerRadius, outerRadius, thetaSegments, phiSegments, thetaStart, thetaLength)
	return &RingBufferGeometry{p: p}
}
