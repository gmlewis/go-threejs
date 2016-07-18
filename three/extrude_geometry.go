// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// ExtrudeGeometry represents extruded geometry from a path shape.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/ExtrudeGeometry
type ExtrudeGeometry struct{ *Geometry }

// JSObject returns the underlying *js.Object.
func (e *ExtrudeGeometry) JSObject() *js.Object { return e.p }

// ExtrudeGeometry returns an ExtrudeGeometry JavaScript class.
func (t *Three) ExtrudeGeometry() *ExtrudeGeometry {
	p := t.ctx.Get("ExtrudeGeometry")
	return ExtrudeGeometryFromJSObject(p)
}

// ExtrudeGeometryFromJSObject returns a wrapped ExtrudeGeometry JavaScript class.
func ExtrudeGeometryFromJSObject(p *js.Object) *ExtrudeGeometry {
	return &ExtrudeGeometry{GeometryFromJSObject(p)}
}

// ExtrudeGeometryOpts represents options passed to the ExtrudeGeometry constructor.
//
//     CurveSegments — int. number of points on the curves
//     Steps — int. number of points used for subdividing segements of extrude spline
//     Amount — int. Depth to extrude the shape
//     BevelEnabled — bool. turn on bevel
//     BevelThickness — float. how deep into the original shape bevel goes
//     BevelSize — float. how far from shape outline is bevel
//     BevelSegments — int. number of bevel layers
//     ExtrudePath — THREE.CurvePath. 3d spline path to extrude shape along. (creates Frames if (frames aren't defined)
//     Frames — THREE.TubeGeometry.FrenetFrames. containing arrays of tangents, normals, binormals
//     Material — int. material index for front and back faces
//     ExtrudeMaterial — int. material index for extrusion and beveled faces
//     UVGenerator — Object. object that provides UV generator functions
type ExtrudeGeometryOpts struct {
	CurveSegments   int
	Steps           int
	Amount          int
	BevelEnabled    bool
	BevelThickness  float64
	BevelSize       float64
	BevelSegments   int
	ExtrudePath     *CurvePath
	Frames          *TubeGeometry
	Material        int
	ExtrudeMaterial int
	UVGenerator     *Object3D
}

// NewExtrudeGeometry returns a new ExtrudeGeometry object.
//
//     shapes — Shape or an array of shapes.
func (t *Three) NewExtrudeGeometry(shapes JSObject, opts *ExtrudeGeometryOpts) *ExtrudeGeometry {
	p := t.ctx.Get("ExtrudeGeometry").New(shapes.JSObject, opts)
	return ExtrudeGeometryFromJSObject(p)
}

// GenerateTopUV TODO description.
func (e *ExtrudeGeometry) GenerateTopUV(geometry, indexA, indexB, indexC float64) *ExtrudeGeometry {
	e.p.Call("generateTopUV", geometry, indexA, indexB, indexC)
	return e
}

// GenerateSideWallUV TODO description.
func (e *ExtrudeGeometry) GenerateSideWallUV(geometry, indexA, indexB, indexC, indexD float64) *ExtrudeGeometry {
	e.p.Call("generateSideWallUV", geometry, indexA, indexB, indexC, indexD)
	return e
}

// AddShapeList TODO description.
func (e *ExtrudeGeometry) AddShapeList(shapes, options float64) *ExtrudeGeometry {
	e.p.Call("addShapeList", shapes, options)
	return e
}

// AddShape TODO description.
func (e *ExtrudeGeometry) AddShape(shape, options float64) *ExtrudeGeometry {
	e.p.Call("addShape", shape, options)
	return e
}
