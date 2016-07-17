// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three_test

import "github.com/gmlewis/go-threejs/three"

func ExampleCubicBezierCurve() {
	t := three.New()
	scene := t.NewScene()

	curve := t.NewCubicBezierCurve(
		t.NewVector3(-10, 0, 0),
		t.NewVector3(-5, 15, 0),
		t.NewVector3(20, 15, 0),
		t.NewVector3(10, 0, 0),
	)

	path := t.NewPath(curve.GetPoints(50))

	geometry := path.CreatePointsGeometry(50)
	opts := three.LineBasicMaterialOpts{
		"color": 0xff0000,
	}
	material := t.NewLineBasicMaterial(opts)

	// Create the final Object3D to add to the scene
	curveObject := t.NewLine(geometry.Points(), material.JSObject())
	scene.Add(curveObject)
}
