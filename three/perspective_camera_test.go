// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three_test

import (
	"github.com/gmlewis/go-threejs/three"
)

func ExamplePerspectiveCamera() {
	const width, height = 800, 600
	t := three.New()
	scene := t.NewScene()

	camera := t.NewPerspectiveCamera(45, width/height, 1, 1000)
	scene.Add(camera)
}
