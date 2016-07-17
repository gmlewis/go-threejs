// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three_test

import (
	"github.com/gmlewis/go-threejs/three"
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
