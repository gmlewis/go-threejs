// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

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
