// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// DodecahedronGeometry represents a dodecagedron.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/DodecahedronGeometry
type DodecahedronGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (d *DodecahedronGeometry) JSObject() *js.Object { return d.p }

// DodecahedronGeometry returns a DodecahedronGeometry JavaScript class.
func (t *Three) DodecahedronGeometry() *DodecahedronGeometry {
	p := t.ctx.Get("DodecahedronGeometry")
	return DodecahedronGeometryFromJSObject(p)
}

// DodecahedronGeometryFromJSObject returns a wrapped DodecahedronGeometry JavaScript class.
func DodecahedronGeometryFromJSObject(p *js.Object) *DodecahedronGeometry {
	return &DodecahedronGeometry{p: p}
}

// NewDodecahedronGeometry returns a new DodecahedronGeometry object.
//
//     radius — Radius of the dodecahedron. Default is 1.
//     detail — Default is 0. Setting this to a value greater than 0 adds
//         vertices making it no longer a dodecahedron.
func (t *Three) NewDodecahedronGeometry(radius, detail float64) *DodecahedronGeometry {
	p := t.ctx.Get("DodecahedronGeometry").New(radius, detail)
	return DodecahedronGeometryFromJSObject(p)
}
