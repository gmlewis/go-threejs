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
	"github.com/gmlewis/go-threejs/v2/three"
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
