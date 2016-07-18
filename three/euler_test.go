// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three_test

import "github.com/gmlewis/go-threejs/three"

func ExampleEuler() {
	t := three.New()

	a := t.NewEuler(0, 1, 1.57, "XYZ")
	b := t.NewVector3(1, 0, 1)
	b.ApplyEuler(a)
}
