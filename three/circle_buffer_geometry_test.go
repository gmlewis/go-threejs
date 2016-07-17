// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three_test

import "github.com/gmlewis/go-threejs/three"

func ExampleCircleBufferGeometry() {
	t := three.New()
	scene := t.NewScene()

	geometry := t.NewCircleBufferGeometry(5, 32, nil)
	opts := three.MeshBasicMaterialOpts{
		"color": 0xffff00,
	}
	material := t.NewMeshBasicMaterial(opts)
	circle := t.NewMesh(geometry, material)
	scene.Add(circle)
}
