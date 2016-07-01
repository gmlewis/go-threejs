// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three_test

import (
	"github.com/gmlewis/go-threejs/three"
)

func ExampleGeometry() {
	t := three.New()

	geometry := t.Geometry().New()

	geometry.vertices.push(
		t.Vector3().New(-10, 10, 0),
		t.Vector3().New(-10, -10, 0),
		t.Vector3().New(10, -10, 0),
	)

	geometry.faces.push(t.Face3().New(0, 1, 2))

	geometry.computeBoundingSphere()
}
