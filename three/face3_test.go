// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three_test

import (
	"github.com/gmlewis/go-threejs/three"
)

func ExampleFace3() {
	t := three.New()

	normal := t.Vector3().New(0, 1, 0)
	color := t.Color().New(0xffaa00)
	/* face := */ t.Face3().New(0, 1, 2, normal, color, 0)
	//...
}