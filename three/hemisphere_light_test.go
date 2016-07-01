// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three_test

import (
	"github.com/gmlewis/go-threejs/three"
)

func ExampleHemisphereLight() {
	t := three.New()
	scene := t.NewScene()

	light := t.NewHemisphereLight(0xffffbb, 0x080820, 1)
	scene.Add(light)
}
