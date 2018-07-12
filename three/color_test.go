// Copyright 2016 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
