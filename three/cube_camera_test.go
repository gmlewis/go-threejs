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

func ExampleCubeCamera() {
	t := three.New()

	camera := t.NewPerspectiveCamera(75, 1.25, 1, 10000)
	renderer := t.NewWebGLRenderer(nil)

	// Create cube camera.
	cubeCamera := t.NewCubeCamera(1, 100000, 128)
	scene := t.NewScene()
	scene.Add(cubeCamera)

	// Create car.
	opts := three.MeshLambertMaterialOpts{
		"color":  0xffffff,
		"envMap": cubeCamera.RenderTarget(),
	}
	chromeMaterial := t.NewMeshLambertMaterial(opts)
	carGeometry := t.NewObject3D() // Load car geometry here.
	car := t.NewMesh(carGeometry, chromeMaterial)
	scene.Add(car)

	// Update the render target cube.
	car.SetVisible(false)
	cubeCamera.Position().Copy(car.Position())
	cubeCamera.UpdateCubeMap(renderer, scene)

	// Render the scene.
	car.SetVisible(true)
	renderer.Render(scene, camera, nil)
}
