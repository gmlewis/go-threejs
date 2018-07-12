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
