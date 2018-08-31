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
)

func ExampleBufferGeometry() {
	t := three.New()

	geometry := t.NewBufferGeometry()
	// create a simple square shape. We duplicate the top left and bottom right
	// vertices because each vertex needs to appear once per triangle.
	vertices := []float64{
		-1.0, -1.0, 1.0,
		1.0, -1.0, 1.0,
		1.0, 1.0, 1.0,

		1.0, 1.0, 1.0,
		-1.0, 1.0, 1.0,
		-1.0, -1.0, 1.0,
	}

	// itemSize = 3 because there are 3 values (components) per vertex
	geometry.AddAttribute("position", t.NewBufferAttribute(vertices, 3).JSObject())
	opts := three.MeshBasicMaterialOpts{"color": 0xff0000}
	material := t.NewMeshBasicMaterial(opts)
	/* mesh := */ t.NewMesh(geometry, material)
	//...
}
