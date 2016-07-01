// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three_test

import (
	"github.com/gmlewis/go-threejs/three"
)

func ExamplePointLight() {
	t := three.New()
	scene := t.NewScene()

	light := t.NewPointLight(0xff0000, 1, 100, 0)
	light.Position().Set(50, 50, 50)
	scene.Add(light)
}
