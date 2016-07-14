// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// RingBufferGeometry represents a ringbuffergeometry.
type RingBufferGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (r *RingBufferGeometry) JSObject() *js.Object { return r.p }

// RingBufferGeometry returns a RingBufferGeometry JavaScript class.
func (t *Three) RingBufferGeometry() *RingBufferGeometry {
	p := t.ctx.Get("RingBufferGeometry")
	return RingBufferGeometryFromJSObject(p)
}

// RingBufferGeometryFromJSObject returns a wrapped RingBufferGeometry JavaScript class.
func RingBufferGeometryFromJSObject(p *js.Object) *RingBufferGeometry {
	return &RingBufferGeometry{p: p}
}

// NewRingBufferGeometry returns a new RingBufferGeometry object.
func (t *Three) NewRingBufferGeometry(innerRadius, outerRadius, thetaSegments, phiSegments, thetaStart, thetaLength float64) *RingBufferGeometry {
	p := t.ctx.Get("RingBufferGeometry").New(innerRadius, outerRadius, thetaSegments, phiSegments, thetaStart, thetaLength)
	return RingBufferGeometryFromJSObject(p)
}
