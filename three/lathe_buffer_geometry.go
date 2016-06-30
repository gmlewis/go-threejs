// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// LatheBufferGeometry represents a lathebuffergeometry.
type LatheBufferGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *LatheBufferGeometry) JSObject() *js.Object { return t.p }

// LatheBufferGeometry returns a LatheBufferGeometry object.
func (t *Three) LatheBufferGeometry() *LatheBufferGeometry {
	p := t.ctx.Get("LatheBufferGeometry")
	return &LatheBufferGeometry{p: p}
}

// New returns a new LatheBufferGeometry object.
func (t *LatheBufferGeometry) New(points, segments, phiStart, phiLength float64) *LatheBufferGeometry {
	p := t.p.New(points, segments, phiStart, phiLength)
	return &LatheBufferGeometry{p: p}
}
