// Copyright 2016 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package three_test

import (
	"math"

	"github.com/gmlewis/go-threejs/v2/three"
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
