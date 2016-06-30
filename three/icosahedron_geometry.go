// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// IcosahedronGeometry represents an icosahedron.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/IcosahedronGeometry
type IcosahedronGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (i *IcosahedronGeometry) JSObject() *js.Object { return i.p }

// IcosahedronGeometry returns an IcosahedronGeometry object.
func (t *Three) IcosahedronGeometry() *IcosahedronGeometry {
	p := t.ctx.Get("IcosahedronGeometry")
	return &IcosahedronGeometry{p: p}
}

// New returns a new IcosahedronGeometry object.
//
//     radius — Default is 1.
//     detail — Default is 0. Setting this to a value greater than 0 adds more vertices making it no longer an icosahedron. When detail is greater than 1, it's effectively a sphere.
func (i *IcosahedronGeometry) New(radius, detail float64) *IcosahedronGeometry {
	p := i.p.New(radius, detail)
	return &IcosahedronGeometry{p: p}
}
