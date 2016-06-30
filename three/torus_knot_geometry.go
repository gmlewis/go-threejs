// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// TorusKnotGeometry represents a torus knot.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/TorusKnotGeometry
type TorusKnotGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *TorusKnotGeometry) JSObject() *js.Object { return t.p }

// TorusKnotGeometry returns a TorusKnotGeometry object.
func (t *Three) TorusKnotGeometry() *TorusKnotGeometry {
	p := t.ctx.Get("TorusKnotGeometry")
	return &TorusKnotGeometry{p: p}
}

// New returns a new TorusKnotGeometry object.
//
//     radius — Default is 100.
//     tube — Default is 40.
//     radialSegments — Default is 64.
//     tubularSegments — Default is 8.
//     p — Default is 2.
//     q — Default is 3.
//     heightScale — Default is 1.
func (t *TorusKnotGeometry) New(radius, tube float64, tubularSegments, radialSegments int, p, q, heightScale float64) *TorusKnotGeometry {
	s := t.p.New(radius, tube, tubularSegments, radialSegments, p, q, heightScale)
	return &TorusKnotGeometry{p: s}
}