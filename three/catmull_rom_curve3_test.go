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
	"github.com/gmlewis/go-threejs/v2/three"
	"github.com/gopherjs/gopherjs/js"
)

func ExampleCatmullRomCurve3() {
	t := three.New()

	// Create a closed wavey loop
	pts := []*js.Object{
		t.NewVector3(-10, 0, 10).JSObject(),
		t.NewVector3(-5, 5, 5).JSObject(),
		t.NewVector3(0, 0, 0).JSObject(),
		t.NewVector3(5, -5, 5).JSObject(),
		t.NewVector3(10, 0, 10).JSObject(),
	}
	curve := t.NewCatmullRomCurve3(pts)

	geometry := t.NewGeometry()
	geometry.SetVertices(curve.GetPoints(50))
	opts := three.LineBasicMaterialOpts{
		"color": 0xff0000,
	}
	material := t.NewLineBasicMaterial(opts)
	/* mesh := */ t.NewMesh(geometry, material)
	//...
}
