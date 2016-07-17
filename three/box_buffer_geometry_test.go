// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three_test

import (
	"github.com/gmlewis/go-threejs/three"
)

func ExampleBoxBufferGeometry() {
	t := three.New()
	scene := t.NewScene()

	geometry := t.NewBoxBufferGeometry(1, 1, 1, nil)
	opts := three.MeshBasicMaterialOpts{
		"color": 0x00ff00,
	}
	material := t.NewMeshBasicMaterial(opts)
	cube := t.NewMesh(geometry, material)
	scene.Add(cube)
}
