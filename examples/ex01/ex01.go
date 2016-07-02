// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

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
	camera = t.NewPerspectiveCamera(75, float64(window.Get("innerWidth").Int())/float64(window.Get("innerHeight").Int()), 1, 10000)
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
