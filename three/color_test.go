// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three_test

import "github.com/gmlewis/go-threejs/three"

func ExampleColor() {
	t := three.New()

	color := t.NewColor()
	color = t.NewColor(0xff0000)
	color = t.NewColor("rgb(255, 0, 0)")
	color = t.NewColor("rgb(100%, 0%, 0%)")
	color = t.NewColor("hsl(0, 100%, 50%)")
	color = t.NewColor(1, 0, 0)
	color.SetHex(0x00ff00)
}
