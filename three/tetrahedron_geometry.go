// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// TetrahedronGeometry represents a tetrahedron.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/TetrahedronGeometry
type TetrahedronGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *TetrahedronGeometry) JSObject() *js.Object { return t.p }

// TetrahedronGeometry returns a TetrahedronGeometry JavaScript class.
func (t *Three) TetrahedronGeometry() *TetrahedronGeometry {
	p := t.ctx.Get("TetrahedronGeometry")
	return TetrahedronGeometryFromJSObject(p)
}

// TetrahedronGeometryFromJSObject returns a wrapped TetrahedronGeometry JavaScript class.
func TetrahedronGeometryFromJSObject(p *js.Object) *TetrahedronGeometry {
	return &TetrahedronGeometry{p: p}
}

// NewTetrahedronGeometry returns a new TetrahedronGeometry object.
//
//     radius — Radius of the tetrahedron. Default is 1.
//     detail — Default is 0. Setting this to a value greater than 0 adds vertices making it no longer a tetrahedron.
func (t *Three) NewTetrahedronGeometry(radius, detail float64) *TetrahedronGeometry {
	p := t.ctx.Get("TetrahedronGeometry").New(radius, detail)
	return TetrahedronGeometryFromJSObject(p)
}
