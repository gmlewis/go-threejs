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
	"github.com/gmlewis/go-threejs/v2/three"
)

func ExampleEdgesHelper() {
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

	edges := t.NewEdgesHelper(object, 0x00ff00, 0.1)

	scene.Add(object)
	scene.Add(edges)
}
