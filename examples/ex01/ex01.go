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

// See: https://github.com/mrdoob/three.js
package main

import (
	"github.com/gmlewis/go-threejs/three"
	"github.com/gopherjs/gopherjs/js"
)

var (
	scene    *three.Scene
	camera   *three.PerspectiveCamera
	renderer *three.WebGLRenderer
	mesh     *three.Mesh
)

func main() {
	t := three.New()

	scene = t.NewScene()

	window := js.Global.Get("window")
	camera = t.NewPerspectiveCamera(75, window.Get("innerWidth").Float()/window.Get("innerHeight").Float(), 1, 10000)
	camera.Position().SetZ(1000)

	geometry := t.NewBoxGeometry(200, 200, 200, nil)
	material := t.NewMeshBasicMaterial(map[string]interface{}{"color": 0xff0000, "wireframe": true})

	mesh = t.NewMesh(geometry, material)
	scene.Add(mesh)

	renderer = t.NewWebGLRenderer(nil)
	renderer.SetSize(window.Get("innerWidth").Float(), window.Get("innerHeight").Float(), false)

	js.Global.Get("document").Get("body").Call("appendChild", renderer.DOMElement())

	animate()
}

func animate() {
	js.Global.Call("requestAnimationFrame", animate)

	mesh.Rotation().SetX(mesh.Rotation().X() + 0.01)
	mesh.Rotation().SetY(mesh.Rotation().Y() + 0.02)

	renderer.Render(scene, camera, nil)
}
