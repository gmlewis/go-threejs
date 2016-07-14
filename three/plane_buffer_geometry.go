// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// PlaneBufferGeometry represents a planebuffergeometry.
type PlaneBufferGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (p *PlaneBufferGeometry) JSObject() *js.Object { return p.p }

// PlaneBufferGeometry returns a PlaneBufferGeometry JavaScript class.
func (t *Three) PlaneBufferGeometry() *PlaneBufferGeometry {
	p := t.ctx.Get("PlaneBufferGeometry")
	return PlaneBufferGeometryFromJSObject(p)
}

// PlaneBufferGeometryFromJSObject returns a wrapped PlaneBufferGeometry JavaScript class.
func PlaneBufferGeometryFromJSObject(p *js.Object) *PlaneBufferGeometry {
	return &PlaneBufferGeometry{p: p}
}

// NewPlaneBufferGeometry returns a new PlaneBufferGeometry object.
func (t *Three) NewPlaneBufferGeometry(width, height, widthSegments, heightSegments float64) *PlaneBufferGeometry {
	p := t.ctx.Get("PlaneBufferGeometry").New(width, height, widthSegments, heightSegments)
	return PlaneBufferGeometryFromJSObject(p)
}
