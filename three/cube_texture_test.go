// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three_test

import "github.com/gmlewis/go-threejs/three"

func ExampleCubeTexture() {
	t := three.New()

	loader := t.NewCubeTextureLoader()
	loader.SetPath("textures/cube/pisa/")

	textureCube := loader.Load([]string{
		"px.png", "nx.png",
		"py.png", "ny.png",
		"pz.png", "nz.png",
	}, nil, nil, nil)

	opts := three.MeshBasicMaterialOpts{
		"color":  0xffffff,
		"envMap": textureCube,
	}
	/* material := */ t.NewMeshBasicMaterial(opts)
}
