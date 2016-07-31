// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// TorusKnotGeometry creates a torus knot, the particular shape of which is defined
// by a pair of coprime integers, p and q. If p and q are not coprime, the result will be a torus link.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/TorusKnotGeometry
type TorusKnotGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *TorusKnotGeometry) JSObject() *js.Object { return t.p }

// TorusKnotGeometry returns a TorusKnotGeometry JavaScript class.
func (t *Three) TorusKnotGeometry() *TorusKnotGeometry {
	p := t.ctx.Get("TorusKnotGeometry")
	return TorusKnotGeometryFromJSObject(p)
}

// TorusKnotGeometryFromJSObject returns a wrapped TorusKnotGeometry JavaScript class.
func TorusKnotGeometryFromJSObject(p *js.Object) *TorusKnotGeometry {
	return &TorusKnotGeometry{p: p}
}

// NewTorusKnotGeometry returns a new TorusKnotGeometry object.
//
//     radius — Default is 100.
//     tube — Default is 40.
//     radialSegments — Default is 64.
//     tubularSegments — Default is 8.
//     p — Default is 2.
//     q — Default is 3.
//     heightScale — Default is 1.
func (t *Three) NewTorusKnotGeometry(radius, tube float64, tubularSegments, radialSegments int, p, q, heightScale float64) *TorusKnotGeometry {
	s := t.ctx.Get("TorusKnotGeometry").New(radius, tube, tubularSegments, radialSegments, p, q, heightScale)
	return TorusKnotGeometryFromJSObject(s)
}
