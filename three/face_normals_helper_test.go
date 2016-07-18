// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three_test

import "github.com/gmlewis/go-threejs/three"

func ExampleFaceNormalsHelper() {
	t := three.New()
	scene := t.NewScene()

	boxOpts := &three.BoxGeometryOpts{
		WidthSegments:  2,
		HeightSegments: 2,
		DepthSegments:  2,
	}
	geometry := t.NewBoxGeometry(10, 10, 10, boxOpts)
	opts := three.MeshBasicMaterialOpts{
		"color": 0xff0000,
	}
	material := t.NewMeshBasicMaterial(opts)
	object := t.NewMesh(geometry, material)

	edges := t.NewFaceNormalsHelper(object, 2, 0x00ff00, 1)

	scene.Add(object)
	scene.Add(edges)
}
