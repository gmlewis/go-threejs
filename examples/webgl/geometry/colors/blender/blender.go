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

// See: http://threejs.org/examples/#webgl_geometry_colors_blender
package main

import (
	"github.com/gmlewis/go-threejs/v2/three"
	"github.com/gopherjs/gopherjs/js"
)

const (
	radius  = 200
	canvasW = 128
	canvasH = 128
)

var (
	container *js.Object
	stats     *js.Object

	camera   *three.PerspectiveCamera
	scene    *three.Scene
	renderer *three.WebGLRenderer

	mesh, mesh2, mesh3 *three.Mesh
	light              *three.DirectionalLight

	mouseX, mouseY = 0.0, 0.0

	document    = js.Global.Get("document")
	window      = js.Global.Get("window")
	windowHalfX = window.Get("innerWidth").Float() / 2
	windowHalfY = window.Get("innerHeight").Float() / 2
)

func init() {
	t := three.New()

	// if ( ! Detector.webgl ) Detector.addGetWebGLMessage()

	container = document.Call("getElementById", "container")

	camera = t.NewPerspectiveCamera(40, window.Get("innerWidth").Float()/window.Get("innerHeight").Float(), 1, 10000)
	camera.Position().SetZ(1800)

	scene = t.NewScene()

	light = t.NewDirectionalLight(0xffffff, 1)
	light.Position().Set(0, 0, 1).Normalize()
	scene.Add(light)

	loader := t.NewJSONLoader()

	loader.Load("http://threejs.org/examples/obj/cubecolors/cubecolors.js", createScene1, nil, nil)
	loader.Load("http://threejs.org/examples/obj/cubecolors/cube_fvc.js", createScene2, nil, nil)

	renderer = t.NewWebGLRenderer(&three.WebGLRendererOpts{Antialias: three.Bool(true)})
	renderer.SetPixelRatio(window.Get("devicePixelRatio").Float())
	renderer.SetSize(window.Get("innerWidth").Float(), window.Get("innerHeight").Float(), true)
	container.Call("appendChild", renderer.DOMElement())

	stats = js.Global.Get("Stats").New()
	stats.Get("domElement").Get("style").Set("position", "absolute")
	stats.Get("domElement").Get("style").Set("top", "0px")
	container.Call("appendChild", stats.Get("domElement"))

	document.Call("addEventListener", "mousemove", onDocumentMouseMove, false)

	window.Call("addEventListener", "resize", onWindowResize, false)
}

func createScene1(geometry *three.Geometry, materials []*three.Material) {
	materials[0].SetShading(three.FlatShading)

	t := three.New()
	mesh = t.NewMesh(geometry, t.NewMultiMaterial(materials))
	mesh.Position().SetX(400)
	mesh.Scale().SetX(250).SetY(250).SetZ(250)
	scene.Add(mesh)
}

func createScene2(geometry *three.Geometry, materials []*three.Material) {
	materials[0].SetShading(three.FlatShading)

	t := three.New()
	mesh2 = t.NewMesh(geometry, t.NewMultiMaterial(materials))
	mesh2.Position().SetX(-400)
	mesh2.Scale().SetX(250).SetY(250).SetZ(250)
	scene.Add(mesh2)
}

func main() {
	animate()
}

func onWindowResize() {
	windowX := window.Get("innerWidth").Float()
	windowY := window.Get("innerHeight").Float()
	windowHalfX = windowX / 2
	windowHalfY = windowY / 2

	camera.SetAspect(windowX / windowY)
	camera.UpdateProjectionMatrix()

	renderer.SetSize(windowX, windowY, true)
}

func onDocumentMouseMove(event *js.Object) {
	mouseX = (event.Get("clientX").Float() - windowHalfX)
	mouseY = (event.Get("clientY").Float() - windowHalfY)
}

func animate() {
	js.Global.Call("requestAnimationFrame", animate)

	render()
	stats.Call("update")
}

func render() {
	camera.Position().SetX(camera.Position().X() + (mouseX-camera.Position().X())*0.05)
	camera.Position().SetY(camera.Position().Y() + (-mouseY-camera.Position().Y())*0.05)

	camera.LookAt(scene.Position())

	if mesh != nil {
		mesh.Rotation().SetX(mesh.Rotation().X() + 0.01)
		mesh.Rotation().SetY(mesh.Rotation().Y() + 0.01)
	}
	if mesh2 != nil {
		mesh2.Rotation().SetX(mesh.Rotation().X() + 0.01)
		mesh2.Rotation().SetY(mesh.Rotation().Y() + 0.01)
	}

	renderer.Render(scene, camera, nil)
}
