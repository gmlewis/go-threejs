// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// TorusKnotBufferGeometry is the BufferGeometry port of TorusKnotGeometry.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/TorusKnotBufferGeometry
type TorusKnotBufferGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *TorusKnotBufferGeometry) JSObject() *js.Object { return t.p }

// TorusKnotBufferGeometry returns a TorusKnotBufferGeometry JavaScript class.
func (t *Three) TorusKnotBufferGeometry() *TorusKnotBufferGeometry {
	p := t.ctx.Get("TorusKnotBufferGeometry")
	return &TorusKnotBufferGeometry{p: p}
}

// NewTorusKnotBufferGeometry returns a new TorusKnotBufferGeometry object.
//
//     radius — Default is 100.
//     tube — Diameter of the tube. Default is 40.
//     tubularSegments — Default is 64.
//     radialSegments — Default is 8.
//     p — This value determines, how many times the geometry winds around its axis of rotational symmetry. Default is 2.
//     q — This value determines, how many times the geometry winds around a circle in the interior of the torus. Default is 3.
func (t *Three) NewTorusKnotBufferGeometry(radius, tube, tubularSegments, radialSegments, p, q float64) *TorusKnotBufferGeometry {
	s := t.ctx.Get("TorusKnotBufferGeometry").New(radius, tube, tubularSegments, radialSegments, p, q)
	return &TorusKnotBufferGeometry{p: s}
}
