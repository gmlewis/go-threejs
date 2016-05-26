// See: https://github.com/mrdoob/three.js
package main

import (
	"github.com/gmlewis/go-threejs/three"
	"github.com/gopherjs/gopherjs/js"
)

var (
	scene                  *three.Scene
	camera, renderer, mesh *js.Object
)

func main() {
	t := three.New()

	scene = t.Scene()

	window := js.Global.Get("window")
	camera = t.PerspectiveCamera(75, float64(window.Get("innerWidth").Int())/float64(window.Get("innerHeight").Int()), 1, 10000)
	camera.Get("position").Set("z", 1000)

	geometry := t.BoxGeometry(200, 200, 200, nil)
	material := t.MeshBasicMaterial(map[string]interface{}{"color": 0xff0000, "wireframe": true})

	mesh = t.Mesh(geometry, material)
	scene.Add(mesh)

	renderer = t.WebGLRenderer(nil)
	renderer.Call("setSize", window.Get("innerWidth").Int(), window.Get("innerHeight").Int())

	js.Global.Get("document").Get("body").Call("appendChild", renderer.Get("domElement"))

	animate()
}

func animate() {
	js.Global.Call("requestAnimationFrame", animate)

	mesh.Get("rotation").Set("x", mesh.Get("rotation").Get("x").Float()+0.01)
	mesh.Get("rotation").Set("y", mesh.Get("rotation").Get("y").Float()+0.02)

	renderer.Call("render", scene, camera)
}
