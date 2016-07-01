// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three_test

import (
	"github.com/gmlewis/go-threejs/three"
)

func ExampleDirectionalLight() {
	t := three.New()
	scene := t.NewScene()

	directionalLight := t.NewDirectionalLight(0xffffff, 0.5)
	directionalLight.Position().Set(0, 1, 0)
	scene.Add(directionalLight)
}
