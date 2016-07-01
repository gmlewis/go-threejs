// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// LatheGeometry represents a mesh with axial symmetry.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/LatheGeometry
type LatheGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (l *LatheGeometry) JSObject() *js.Object { return l.p }

// LatheGeometry returns a LatheGeometry JavaScript class.
func (t *Three) LatheGeometry() *LatheGeometry {
	p := t.ctx.Get("LatheGeometry")
	return &LatheGeometry{p: p}
}

// NewLatheGeometry returns a new LatheGeometry object.
//
//     points — Slice of Vector2s.
//     segments — the number of circumference segments to generate. Default is 12.
//     phiStart — the starting angle in radians. Default is 0.
//     phiLength — the radian (0 to 2PI) range of the lathed section 2PI is a
//         closed lathe, less than 2PI is a portion. Default is 2*PI
func (t *Three) NewLatheGeometry(points []*js.Object, segments int, phiStart, phiLength float64) *LatheGeometry {
	p := t.ctx.Get("LatheGeometry").New(points, segments, phiStart, phiLength)
	return &LatheGeometry{p: p}
}
