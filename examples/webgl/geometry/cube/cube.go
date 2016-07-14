// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// See: http://threejs.org/examples/#webgl_geometry_colors
package main

import (
	"github.com/gmlewis/go-threejs/three"
	"github.com/gopherjs/gopherjs/js"
)

var (
	camera   *three.PerspectiveCamera
	scene    *three.Scene
	renderer *three.WebGLRenderer
	mesh     *three.Mesh

	document = js.Global.Get("document")
	window   = js.Global.Get("window")
)

func init() {
	t := three.New()
	// if ( ! Detector.webgl ) Detector.addGetWebGLMessage()

	camera = t.NewPerspectiveCamera(70, window.Get("innerWidth").Float()/window.Get("innerHeight").Float(), 1, 1000)
	camera.Position().SetZ(400)

	scene = t.NewScene()

	texture := t.NewTextureLoader().Load("crate.gif", nil, nil, nil)

	geometry := t.NewBoxBufferGeometry(200, 200, 200, nil)
	material := t.NewMeshBasicMaterial(three.MeshBasicMaterialOpts{"map": texture.JSObject()})

	mesh = t.NewMesh(geometry, material)
	scene.Add(mesh)

	renderer = t.NewWebGLRenderer(nil)
	renderer.SetPixelRatio(window.Get("devicePixelRatio").Float())
	renderer.SetSize(window.Get("innerWidth").Float(), window.Get("innerHeight").Float(), true)
	document.Get("body").Call("appendChild", renderer.DOMElement())

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

func animate() {
	js.Global.Call("requestAnimationFrame", animate)

	mesh.Rotation().SetX(mesh.Rotation().X() + 0.005)
	mesh.Rotation().SetY(mesh.Rotation().Y() + 0.01)

	renderer.Render(scene, camera, nil)
}
