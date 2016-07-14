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
func (c *CircleBufferGeometry) JSObject() *js.Object { return c.p }

// CircleBufferGeometry returns a CircleBufferGeometry JavaScript class.
func (t *Three) CircleBufferGeometry() *CircleBufferGeometry {
	p := t.ctx.Get("CircleBufferGeometry")
	return CircleBufferGeometryFromJSObject(p)
}

// CircleBufferGeometryFromJSObject returns a wrapped CircleBufferGeometry JavaScript class.
func CircleBufferGeometryFromJSObject(p *js.Object) *CircleBufferGeometry {
	return &CircleBufferGeometry{p: p}
}

// NewCircleBufferGeometry returns a new CircleBufferGeometry object.
func (t *Three) NewCircleBufferGeometry(radius, segments, thetaStart, thetaLength float64) *CircleBufferGeometry {
	p := t.ctx.Get("CircleBufferGeometry").New(radius, segments, thetaStart, thetaLength)
	return CircleBufferGeometryFromJSObject(p)
}
