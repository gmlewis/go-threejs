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

// NOTE: This example has not been completed yet.

// See: http://threejs.org/examples/#webgl_geometry_dynamic
package main

import (
	"math"

	"github.com/gmlewis/go-threejs/three"
	"github.com/gopherjs/gopherjs/js"
)

const (
	waterURL = "http://threejs.org/examples/textures/water.jpg"
)

var (
	container *js.Object
	stats     *js.Object

	camera   *three.PerspectiveCamera
	controls *three.FirstPersonControls
	scene    *three.Scene
	renderer *three.WebGLRenderer

	mesh     *three.Mesh
	texture  *three.Texture
	geometry *three.PlaneGeometry
	material *three.Material

	worldWidth, worldDepth         = 128, 128
	worldHalfWidth, worldHalfDepth = worldWidth / 2, worldDepth / 2

	document = js.Global.Get("document")
	window   = js.Global.Get("window")
)

func init() {
	t := three.New()

	clock = t.NewClock(true)

	// if ( ! Detector.webgl ) {
	//     Detector.addGetWebGLMessage()
	//     document.getElementById( 'container' ).innerHTML = ""
	// }

	container = document.Call("getElementById", "container")

	camera = t.NewPerspectiveCamera(60, window.Get("innerWidth").Float()/window.Get("innerHeight").Float(), 1, 20000)
	camera.Position().SetY(200)

	controls = t.NewFirstPersonControls(camera)

	controls.SetMovementSpeed(500)
	controls.SetLookSpeed(0.1)

	scene = t.NewScene()
	scene.SetFog(t.NewFogExp2(xaaccff, 0.0007))

	geometry = t.NewPlaneGeometry(20000, 20000, worldWidth-1, worldDepth-1)
	geometry.SetRotateX(math.Pi / 2)

	for i, n := 0, len(geometry.Vertices()); i < n; i++ {
		geometry.Vertices(i).SetY(35 * math.Sin(float64(i)/2))
	}

	texture = t.NewTextureLoader().Load(waterURL)
	texture.SetWrapS(three.RepeatWrapping)
	texture.SetWrapT(three.RepeatWrapping)
	texture.Repeat().Set(5, 5)

	opts := three.MeshBasicMaterialOpts{
		"color": 0x0044ff,
		"map":   texture.JSObject(),
	}
	material = t.NewMeshBasicMaterial(opts)

	mesh = t.NewMesh(geometry, material)
	scene.Add(mesh)

	renderer = t.NewWebGLRenderer(nil)
	renderer.SetClearColor(0xaaccff)
	renderer.SetPixelRatio(window.Get("devicePixelRatio").Float())
	renderer.SetSize(window.Get("innerWidth").Float(), window.Get("innerHeight").Float(), true)

	container.Set("innerHTML", "")

	container.Call("appendChild", renderer.DOMElement.JSObject())

	stats = js.Global.Get("Stats").New()
	style := stats.Get("domElement").Get("style")
	style.Set("position", "absolute")
	style.Set("top", "0px")
	container.Call("appendChild", stats.Get("domElement"))

	window.Call("addEventListener", "resize", onWindowResize, false)

}

func main() {
	animate()
}

func onWindowResize() {
	windowX := window.Get("innerWidth").Float()
	windowY := window.Get("innerHeight").Float()

	camera.SetAspect(windowX / windowY)
	camera.UpdateProjectionMatrix()

	renderer.SetSize(windowX, windowY, true)

	controls.HandleResize()
}

func animate() {
	js.Global.Call("requestAnimationFrame", animate)

	render()
	stats.Call("update")
}

func render() {
	delta := clock.GetDelta()
	time := clock.GetElapsedTime() * 10

	for i, n := 0, len(geometry.Vertices()); i < n; i++ {
		geometry.Vertices(i).SetY(35 * math.Sin(i/5+(time+i)/7))
	}

	renderer.Render(scene, camera, nil)
}
