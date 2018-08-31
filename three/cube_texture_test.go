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

import "github.com/gmlewis/go-threejs/v2/three"

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
