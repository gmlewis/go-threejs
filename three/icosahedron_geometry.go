// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// IcosahedronGeometry represents an icosahedron geometry.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/IcosahedronGeometry
type IcosahedronGeometry struct{ *PolyhedronGeometry }

// JSObject returns the underlying *js.Object.
func (i *IcosahedronGeometry) JSObject() *js.Object { return i.p }

// IcosahedronGeometry returns an IcosahedronGeometry JavaScript class.
func (t *Three) IcosahedronGeometry() *IcosahedronGeometry {
	p := t.ctx.Get("IcosahedronGeometry")
	return icosahedronGeometry(p)
}

// icosahedronGeometry returns a wrapped IcosahedronGeometry JavaScript class.
func icosahedronGeometry(p *js.Object) *IcosahedronGeometry {
	return &IcosahedronGeometry{polyhedronGeometry(p)}
}

// NewIcosahedronGeometry returns a new IcosahedronGeometry object.
//
//     radius — Default is 1.
//     detail — Default is 0. Setting this to a value greater than 0 adds more
//         vertices making it no longer an icosahedron. When detail is greater
//         than 1, it's effectively a sphere.
func (t *Three) NewIcosahedronGeometry(radius, detail float64) *IcosahedronGeometry {
	p := t.ctx.Get("IcosahedronGeometry").New(radius, detail)
	return icosahedronGeometry(p)
}
