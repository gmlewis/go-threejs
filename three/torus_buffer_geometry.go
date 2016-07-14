// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// TorusBufferGeometry represents a torusbuffergeometry.
type TorusBufferGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *TorusBufferGeometry) JSObject() *js.Object { return t.p }

// TorusBufferGeometry returns a TorusBufferGeometry JavaScript class.
func (t *Three) TorusBufferGeometry() *TorusBufferGeometry {
	p := t.ctx.Get("TorusBufferGeometry")
	return TorusBufferGeometryFromJSObject(p)
}

// TorusBufferGeometryFromJSObject returns a wrapped TorusBufferGeometry JavaScript class.
func TorusBufferGeometryFromJSObject(p *js.Object) *TorusBufferGeometry {
	return &TorusBufferGeometry{p: p}
}

// NewTorusBufferGeometry returns a new TorusBufferGeometry object.
func (t *Three) NewTorusBufferGeometry(radius, tube, radialSegments, tubularSegments, arc float64) *TorusBufferGeometry {
	p := t.ctx.Get("TorusBufferGeometry").New(radius, tube, radialSegments, tubularSegments, arc)
	return TorusBufferGeometryFromJSObject(p)
}
