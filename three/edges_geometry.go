// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// EdgesGeometry represents an edgesgeometry.
type EdgesGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (e *EdgesGeometry) JSObject() *js.Object { return e.p }

// EdgesGeometry returns an EdgesGeometry JavaScript class.
func (t *Three) EdgesGeometry() *EdgesGeometry {
	p := t.ctx.Get("EdgesGeometry")
	return EdgesGeometryFromJSObject(p)
}

// EdgesGeometryFromJSObject returns a wrapped EdgesGeometry JavaScript class.
func EdgesGeometryFromJSObject(p *js.Object) *EdgesGeometry {
	return &EdgesGeometry{p: p}
}

// NewEdgesGeometry returns a new EdgesGeometry object.
func (t *Three) NewEdgesGeometry(geometry, thresholdAngle float64) *EdgesGeometry {
	p := t.ctx.Get("EdgesGeometry").New(geometry, thresholdAngle)
	return EdgesGeometryFromJSObject(p)
}
