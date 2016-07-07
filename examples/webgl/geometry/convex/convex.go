// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// See: http://threejs.org/examples/#webgl_geometry_colors
package main

import (
	"math"
	"time"

	"github.com/gmlewis/go-threejs/three"
	"github.com/gopherjs/gopherjs/js"
)

var (
	container *js.Object
	stats     *js.Object

	camera   *three.PerspectiveCamera
	scene    *three.Scene
	renderer *three.WebGLRenderer

	document = js.Global.Get("document")
	window   = js.Global.Get("window")
)

func init() {
	t := three.New()
	// if ( ! Detector.webgl ) Detector.addGetWebGLMessage()

	container = document.Call("createElement", "div")
	document.Get("body").Call("appendChild", container)

	camera = t.NewPerspectiveCamera(45, window.Get("innerWidth").Float()/window.Get("innerHeight").Float(), 1, 2000)
	camera.Position().SetY(400)

	scene = t.NewScene()

	scene.Add(t.NewAmbientLight(0x404040, 1))

	light := t.NewDirectionalLight(0xffffff, 1)
	light.Position().Set(0, 1, 0)
	scene.Add(light)

	mapTexture := t.NewTextureLoader().Load("http://threejs.org/examples/textures/UV_Grid_Sm.jpg", nil, nil, nil)
	mapTexture.SetWrapS(three.RepeatWrapping).SetWrapT(three.RepeatWrapping).SetAnisotropy(16)

	materials := []*three.Material{
		t.NewMeshLambertMaterial(three.MeshLambertMaterialOpts{"map": mapTexture.JSObject()}),
		t.NewMeshBasicMaterial(three.MeshBasicMaterialOpts{"color": 0xffffff, "wireframe": true, "transparent": true, "opacity": 0.1}),
	}

	// tetrahedron

	points := []*three.Vector3{
		t.NewVector3(100, 0, 0),
		t.NewVector3(0, 100, 0),
		t.NewVector3(0, 0, 100),
		t.NewVector3(0, 0, 0),
	}

	object := t.SceneUtils().CreateMultiMaterialObject(t.NewConvexGeometry(points), materials)
	object.Position().Set(0, 0, 0)
	scene.Add(object)

	// cube

	points = []*three.Vector3{
		t.NewVector3(50, 50, 50),
		t.NewVector3(50, 50, -50),
		t.NewVector3(-50, 50, -50),
		t.NewVector3(-50, 50, 50),
		t.NewVector3(50, -50, 50),
		t.NewVector3(50, -50, -50),
		t.NewVector3(-50, -50, -50),
		t.NewVector3(-50, -50, 50),
	}

	object = t.SceneUtils().CreateMultiMaterialObject(t.NewConvexGeometry(points), materials)
	object.Position().Set(-200, 0, -200)
	scene.Add(object)

	// random convex

	points = []*three.Vector3{}
	for i := 0; i < 30; i++ {
		points = append(points, randomPointInSphere(50))
	}

	object = t.SceneUtils().CreateMultiMaterialObject(t.NewConvexGeometry(points), materials)
	object.Position().Set(-200, 0, 200)
	scene.Add(object)

	object = t.NewAxisHelper(50)
	object.Position().Set(200, 0, -200)
	scene.Add(object)

	object = t.NewArrowHelper(t.NewVector3(0, 1, 0), t.NewVector3(0, 0, 0), 50)
	object.Position().Set(200, 0, 400)
	scene.Add(object)

	renderer = t.NewWebGLRenderer(&three.WebGLRendererOpts{Antialias: three.Bool(true)})
	renderer.SetPixelRatio(window.Get("devicePixelRatio").Float())
	renderer.SetSize(window.Get("innerWidth").Float(), window.Get("innerHeight").Float(), true)
	container.Call("appendChild", renderer.DOMElement())

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
}

func randomPointInSphere(radius float64) *three.Vector3 {
	t := three.New()
	return t.NewVector3(
		(math.Random()-0.5)*2*radius,
		(math.Random()-0.5)*2*radius,
		(math.Random()-0.5)*2*radius,
	)
}

func animate() {
	js.Global.Call("requestAnimationFrame", animate)

	render()
	stats.Call("update")
}

func render() {
	timer := float64(time.Now().Nanosecond()) * 1e5

	camera.Position().SetX(math.Cos(timer) * 800)
	camera.Position().SetY(math.Sin(timer) * 800)

	camera.LookAt(scene.Position())

	children := scene.Children()
	for i := 0; i < len(children); i++ {
		object := children[i]
		object.Rotation().SetX(timer * 5)
		object.Rotation().SetY(timer * 2.5)
	}

	renderer.Render(scene, camera, nil)
}
