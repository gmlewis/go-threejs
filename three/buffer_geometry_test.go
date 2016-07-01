// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three_test

import (
	"github.com/gmlewis/go-threejs/three"
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
