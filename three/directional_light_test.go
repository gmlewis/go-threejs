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

import (
	"github.com/gmlewis/go-threejs/three"
)

func ExampleDirectionalLight() {
	t := three.New()
	scene := t.NewScene()

	// White directional light at half intensity shining from the top.
	light := t.NewDirectionalLight(0xffffff, 0.5)
	light.Position().Set(0, 1, 0)
	scene.Add(light)
}
