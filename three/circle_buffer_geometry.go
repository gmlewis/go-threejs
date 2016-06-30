// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CircleBufferGeometry represents a circlebuffergeometry.
type CircleBufferGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *CircleBufferGeometry) JSObject() *js.Object { return t.p }

// CircleBufferGeometry returns a CircleBufferGeometry object.
func (t *Three) CircleBufferGeometry() *CircleBufferGeometry {
	p := t.ctx.Get("CircleBufferGeometry")
	return &CircleBufferGeometry{p: p}
}

// New returns a new CircleBufferGeometry object.
func (t *CircleBufferGeometry) New(radius, segments, thetaStart, thetaLength float64) *CircleBufferGeometry {
	p := t.p.New(radius, segments, thetaStart, thetaLength)
	return &CircleBufferGeometry{p: p}
}
