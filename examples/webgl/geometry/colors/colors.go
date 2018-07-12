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

// See: http://threejs.org/examples/#webgl_geometry_colors
package main

import (
	"math"

	"github.com/gmlewis/go-threejs/three"
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

	mesh                   *three.Mesh
	group1, group2, group3 *three.Group
	light                  *three.DirectionalLight

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

	camera = t.NewPerspectiveCamera(20, window.Get("innerWidth").Float()/window.Get("innerHeight").Float(), 1, 10000)
	camera.Position().SetZ(1800)

	scene = t.NewScene()

	light = t.NewDirectionalLight(0xffffff, 1)
	light.Position().Set(0, 0, 1)
	scene.Add(light)

	// shadow

	canvas := document.Call("createElement", "canvas")
	canvas.Set("width", canvasW)
	canvas.Set("height", canvasH)

	context := canvas.Call("getContext", "2d")
	gradient := context.Call("createRadialGradient", canvasW/2, canvasH/2, 0, canvasW/2, canvasH/2, canvasW/2)
	gradient.Call("addColorStop", 0.1, "rgba(210,210,210,1)")
	gradient.Call("addColorStop", 1, "rgba(255,255,255,1)")

	context.Set("fillStyle", gradient)
	context.Call("fillRect", 0, 0, canvasW, canvasH)

	shadowTexture := t.NewTexture(canvas, nil)
	shadowTexture.SetNeedsUpdate(true)

	shadowMaterial := t.NewMeshBasicMaterial(three.MeshBasicMaterialOpts{"map": shadowTexture})
	shadowGeo := t.NewPlaneBufferGeometry(300, 300, 1, 1)

	mesh = t.NewMesh(shadowGeo, shadowMaterial)
	mesh.Position().SetY(-250)
	mesh.Rotation().SetX(-math.Pi / 2)
	scene.Add(mesh)

	mesh = t.NewMesh(shadowGeo, shadowMaterial)
	mesh.Position().SetY(-250)
	mesh.Position().SetX(-400)
	mesh.Rotation().SetX(-math.Pi / 2)
	scene.Add(mesh)

	mesh = t.NewMesh(shadowGeo, shadowMaterial)
	mesh.Position().SetY(-250)
	mesh.Position().SetX(400)
	mesh.Rotation().SetX(-math.Pi / 2)
	scene.Add(mesh)

	faceIndices := []string{"a", "b", "c"}

	geometry := t.NewIcosahedronGeometry(radius, 1)
	geometry2 := t.NewIcosahedronGeometry(radius, 1)
	geometry3 := t.NewIcosahedronGeometry(radius, 1)

	faces := geometry.Faces() // expensive operation.
	faces2 := geometry2.Faces()
	faces3 := geometry3.Faces()
	verts := geometry.Vertices()

	for i := 0; i < len(faces); i++ {
		f := faces[i]
		f2 := faces2[i]
		f3 := faces3[i]

		for j := 0; j < 3; j++ {
			vertexIndex := f.Index(faceIndices[j])

			p := verts[vertexIndex]

			color := t.NewColor(0xffffff)
			color.SetHSL((p.Y()/radius+1)/2, 1.0, 0.5)

			f.SetVertexColor(j, color)

			color = t.NewColor(0xffffff)
			color.SetHSL(0.0, (p.Y()/radius+1)/2, 0.5)

			f2.SetVertexColor(j, color)

			color = t.NewColor(0xffffff)
			color.SetHSL(0.125*float64(vertexIndex)/float64(len(verts)), 1.0, 0.5)

			f3.SetVertexColor(j, color)
		}
	}

	materials := []three.JSObject{
		t.NewMeshPhongMaterial(three.MeshPhongMaterialOpts{
			"color":        0xffffff,
			"shading":      three.FlatShading,
			"vertexColors": three.VertexColors, "shininess": 0,
		}),
		t.NewMeshBasicMaterial(three.MeshBasicMaterialOpts{
			"color":       0x000000,
			"shading":     three.FlatShading,
			"wireframe":   true,
			"transparent": true,
		}),
	}

	group1 = t.SceneUtils().CreateMultiMaterialObject(geometry, materials)
	group1.Position().SetX(-400)
	group1.Rotation().SetX(-1.87)
	scene.Add(group1)

	group2 = t.SceneUtils().CreateMultiMaterialObject(geometry2, materials)
	group2.Position().SetX(400)
	group2.Rotation().SetX(0)
	scene.Add(group2)

	group3 = t.SceneUtils().CreateMultiMaterialObject(geometry3, materials)
	group3.Position().SetX(0)
	group3.Rotation().SetX(0)
	scene.Add(group3)

	renderer = t.NewWebGLRenderer(&three.WebGLRendererOpts{Antialias: three.Bool(true)})
	renderer.SetClearColor(t.NewColor(0xffffff), 1)
	renderer.SetPixelRatio(window.Get("devicePixelRatio").Float())
	renderer.SetSize(window.Get("innerWidth").Float(), window.Get("innerHeight").Float(), true)
	container.Call("appendChild", renderer.DOMElement())

	stats = js.Global.Get("Stats").New()
	container.Call("appendChild", stats.Get("dom"))

	document.Call("addEventListener", "mousemove", onDocumentMouseMove, false)

	window.Call("addEventListener", "resize", onWindowResize, false)
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

	renderer.Render(scene, camera, nil)
}
