// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three_test

import (
	"github.com/gmlewis/go-threejs/three"
)

func ExampleGeometry() {
	t := three.New()

	geometry := t.NewGeometry()
	/*
		geometry.vertices.push(
			t.NewVector3(-10, 10, 0),
			t.NewVector3(-10, -10, 0),
			t.NewVector3(10, -10, 0),
		)

		geometry.faces.push(t.NewFace3(0, 1, 2))
	*/
	geometry.ComputeBoundingSphere()
}
