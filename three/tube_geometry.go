// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// TubeGeometry creates a tube that extrudes along a three-dimensional curve.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/TubeGeometry
type TubeGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *TubeGeometry) JSObject() *js.Object { return t.p }

// TubeGeometry returns a TubeGeometry JavaScript class.
func (t *Three) TubeGeometry() *TubeGeometry {
	p := t.ctx.Get("TubeGeometry")
	return TubeGeometryFromJSObject(p)
}

// TubeGeometryFromJSObject returns a wrapped TubeGeometry JavaScript class.
func TubeGeometryFromJSObject(p *js.Object) *TubeGeometry {
	return &TubeGeometry{p: p}
}

// NewTubeGeometry returns a new TubeGeometry object.
//
//     path — Curve - A path that inherits from the Curve base class
//     segments — Integer - The number of segments that make up the tube, default is 64
//     radius — Float - The radius of the tube, default is 1
//     radiusSegments — Integer - The number of segments that make up the cross-section, default is 8
//     closed — Boolean Is the tube open or closed, default is false
func (t *Three) NewTubeGeometry(path, segments, radius, radialSegments, closed, taper float64) *TubeGeometry {
	p := t.ctx.Get("TubeGeometry").New(path, segments, radius, radialSegments, closed, taper)
	return TubeGeometryFromJSObject(p)
}

/* TODO:
Parameters is an object with all of the parameters that were used to generate the geometry.
Tangents is an array of Vector3 tangents.
Normals is an array of Vector3 normals.
Binormals is an array of Vector3 binormals.
*/
