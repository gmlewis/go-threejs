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
func (l *LatheBufferGeometry) JSObject() *js.Object { return l.p }

// LatheBufferGeometry returns a LatheBufferGeometry JavaScript class.
func (t *Three) LatheBufferGeometry() *LatheBufferGeometry {
	p := t.ctx.Get("LatheBufferGeometry")
	return LatheBufferGeometryFromJSObject(p)
}

// LatheBufferGeometryFromJSObject returns a wrapped LatheBufferGeometry JavaScript class.
func LatheBufferGeometryFromJSObject(p *js.Object) *LatheBufferGeometry {
	return &LatheBufferGeometry{p: p}
}

// NewLatheBufferGeometry returns a new LatheBufferGeometry object.
func (t *Three) NewLatheBufferGeometry(points, segments, phiStart, phiLength float64) *LatheBufferGeometry {
	p := t.ctx.Get("LatheBufferGeometry").New(points, segments, phiStart, phiLength)
	return LatheBufferGeometryFromJSObject(p)
}
