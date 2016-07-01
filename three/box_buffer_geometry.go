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
func (b *BoxBufferGeometry) JSObject() *js.Object { return b.p }

// BoxBufferGeometry returns a BoxBufferGeometry JavaScript class.
func (t *Three) BoxBufferGeometry() *BoxBufferGeometry {
	p := t.ctx.Get("BoxBufferGeometry")
	return &BoxBufferGeometry{p: p}
}

// NewBoxBufferGeometry returns a new BoxBufferGeometry object.
func (t *Three) NewBoxBufferGeometry(width, height, depth, widthSegments, heightSegments, depthSegments float64) *BoxBufferGeometry {
	p := t.ctx.Get("BoxBufferGeometry").New(width, height, depth, widthSegments, heightSegments, depthSegments)
	return &BoxBufferGeometry{p: p}
}
