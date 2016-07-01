// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three_test

import (
	"github.com/gmlewis/go-threejs/three"
)

func ExampleSpotLight() {
	t := three.New()
	scene := t.NewScene()

	// white spotlight shining from the side, casting shadow

	light := t.NewSpotLight(0xffffff, 0, 0, 0, 0, 0)
	light.Position().Set(100, 1000, 100)

	light.SetCastShadow(true)

	light.Shadow().MapSize().
		SetWidth(1024).
		SetHeight(1024)

	light.Shadow().Camera().
		SetNear(500).
		SetFar(4000).
		SetFOV(30)

	scene.Add(light)
}
