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

import "github.com/gmlewis/go-threejs/three"

func ExampleCubicBezierCurve3() {
	t := three.New()
	scene := t.NewScene()

	curve := t.NewCubicBezierCurve3(
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
	curveObject := t.NewLine(geometry, material)
	scene.Add(curveObject)
}
