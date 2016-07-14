// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// OctahedronGeometry represents an octahedron.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/OctahedronGeometry
type OctahedronGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (o *OctahedronGeometry) JSObject() *js.Object { return o.p }

// OctahedronGeometry returns an OctahedronGeometry JavaScript class.
func (t *Three) OctahedronGeometry() *OctahedronGeometry {
	p := t.ctx.Get("OctahedronGeometry")
	return OctahedronGeometryFromJSObject(p)
}

// OctahedronGeometryFromJSObject returns a wrapped OctahedronGeometry JavaScript class.
func OctahedronGeometryFromJSObject(p *js.Object) *OctahedronGeometry {
	return &OctahedronGeometry{p: p}
}

// NewOctahedronGeometry returns a new OctahedronGeometry object.
//
//     radius — Default is 1.
//     detail — Default is 0. Setting this to a value greater than 0 adds more
//         vertices making it no longer an icosahedron. When detail is greater
//         than 1, it's effectively a sphere.
func (t *Three) NewOctahedronGeometry(radius, detail float64) *OctahedronGeometry {
	p := t.ctx.Get("OctahedronGeometry").New(radius, detail)
	return OctahedronGeometryFromJSObject(p)
}
