// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three_test

import (
	"github.com/gmlewis/go-threejs/three"
)

func ExampleCubeCamera() {
	t := three.New()

	camera := t.PerspectiveCamera().New(75, 1.25, 1, 10000)
	renderer := t.WebGLRenderer().New(nil)

	// Create cube camera.
	cubeCamera := t.CubeCamera().New(1, 100000, 128)
	scene := t.Scene().New()
	scene.Add(cubeCamera)

	// Create car.
	opts := three.MeshLambertMaterialOpts{
		"color":  0xffffff,
		"envMap": cubeCamera.RenderTarget(),
	}
	chromeMaterial := t.MeshLambertMaterial().New(opts)
	carGeometry := t.Object3D().New() // Load car geometry here.
	car := t.Mesh().New(carGeometry, chromeMaterial)
	scene.Add(car)

	// Update the render target cube.
	car.SetVisible(false)
	cubeCamera.Position().Copy(car.Position())
	cubeCamera.UpdateCubeMap(renderer, scene)

	// Render the scene.
	car.SetVisible(true)
	renderer.Render(scene, camera, nil)
}
