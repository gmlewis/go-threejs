// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three_test

import (
	"github.com/gmlewis/go-threejs/three"
)

func ExampleAmbientLight() {
	t := three.New()
	scene := t.NewScene()

	light := t.NewAmbientLight(0x404040, 1) // soft white light
	scene.Add(light)
}
