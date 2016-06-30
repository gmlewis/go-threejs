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
func (t *RingBufferGeometry) JSObject() *js.Object { return t.p }

// RingBufferGeometry returns a RingBufferGeometry object.
func (t *Three) RingBufferGeometry() *RingBufferGeometry {
	p := t.ctx.Get("RingBufferGeometry")
	return &RingBufferGeometry{p: p}
}

// New returns a new RingBufferGeometry object.
func (t *RingBufferGeometry) New(innerRadius, outerRadius, thetaSegments, phiSegments, thetaStart, thetaLength float64) *RingBufferGeometry {
	p := t.p.New(innerRadius, outerRadius, thetaSegments, phiSegments, thetaStart, thetaLength)
	return &RingBufferGeometry{p: p}
}