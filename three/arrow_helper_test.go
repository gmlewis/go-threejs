// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three_test

import (
	"github.com/gmlewis/go-threejs/three"
)

func ExampleArrowHelper() {
	t := three.New()
	scene := t.NewScene()

	dir := t.NewVector3(1, 0, 0)
	origin := t.NewVector3(0, 0, 0)
	opts := &three.ArrowHelperOpts{
		Length: three.Float64(1),
		Color:  t.NewColor(0xffff00),
	}
	arrowHelper := t.NewArrowHelper(dir, origin, opts)
	scene.Add(arrowHelper)
}
