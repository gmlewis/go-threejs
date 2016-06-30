// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// BoxBufferGeometry represents a boxbuffergeometry.
type BoxBufferGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *BoxBufferGeometry) JSObject() *js.Object { return t.p }

// BoxBufferGeometry returns a BoxBufferGeometry object.
func (t *Three) BoxBufferGeometry() *BoxBufferGeometry {
	p := t.ctx.Get("BoxBufferGeometry")
	return &BoxBufferGeometry{p: p}
}

// New returns a new BoxBufferGeometry object.
func (t *BoxBufferGeometry) New(width, height, depth, widthSegments, heightSegments, depthSegments float64) *BoxBufferGeometry {
	p := t.p.New(width, height, depth, widthSegments, heightSegments, depthSegments)
	return &BoxBufferGeometry{p: p}
}
