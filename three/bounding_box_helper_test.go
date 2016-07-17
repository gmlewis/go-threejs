// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three_test

import (
	"github.com/gmlewis/go-threejs/three"
)

func ExampleBoundingBoxHelper() {
	t := three.New()
	scene := t.NewScene()

	hex := 0xff0000

	opts := three.MeshLambertMaterialOpts{
		"color": 0x00ff00,
	}
	sphereMaterial := t.NewMeshLambertMaterial(opts)
	sphere := t.NewMesh(t.NewSphereGeometry(30, 12, 12, nil), sphereMaterial)
	scene.Add(sphere)

	bbox := t.NewBoundingBoxHelper(sphere, hex)
	bbox.Update()
	scene.Add(bbox)
}
