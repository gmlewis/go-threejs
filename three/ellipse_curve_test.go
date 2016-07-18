// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three_test

import (
	"math"

	"github.com/gmlewis/go-threejs/three"
)

func ExampleEllipseCurve() {
	t := three.New()
	scene := t.NewScene()

	ellipseOpts := &three.EllipseCurveOpts{
		StartAngle: 0,
		EndAngle:   2 * math.Pi,
		Clockwise:  false,
		Rotation:   0,
	}
	curve := t.NewEllipseCurve(
		0, 0, // ax, aY
		10, 10, // xRadius, yRadius
		ellipseOpts)

	path := t.NewPath(curve.GetPoints(50))
	geometry := path.CreatePointsGeometry(50)
	opts := three.LineBasicMaterialOpts{
		"color": 0xff0000,
	}
	material := t.NewLineBasicMaterial(opts)

	// Create the final Object3d to add to the scene
	ellipse := t.NewLine(geometry, material)
	scene.Add(ellipse)
}
