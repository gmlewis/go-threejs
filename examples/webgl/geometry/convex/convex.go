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
	"math/rand"
	"time"

	"github.com/gmlewis/go-threejs/v2/three"
	"github.com/gopherjs/gopherjs/js"
)

const uvImageURL = "http://threejs.org/examples/textures/UV_Grid_Sm.jpg"

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

	loader := t.NewTextureLoader()
	mapTexture := loader.Load(uvImageURL, nil, nil, nil)
	mapTexture.SetWrapS(three.RepeatWrapping).SetWrapT(three.RepeatWrapping).SetAnisotropy(16)

	materials := []three.JSObject{
		t.NewMeshLambertMaterial(three.MeshLambertMaterialOpts{"map": mapTexture.JSObject()}),
		t.NewMeshBasicMaterial(three.MeshBasicMaterialOpts{"color": 0xffffff, "wireframe": true, "transparent": true, "opacity": 0.1}),
	}

	// tetrahedron
	{
		points := []*js.Object{
			t.NewVector3(100, 0, 0).JSObject(),
			t.NewVector3(0, 100, 0).JSObject(),
			t.NewVector3(0, 0, 100).JSObject(),
			t.NewVector3(0, 0, 0).JSObject(),
		}
		geom := three.GeometryFromJSObject(t.JSObject().Get("ConvexGeometry").New(points))
		object := t.SceneUtils().CreateMultiMaterialObject(geom, materials)
		object.Position().Set(0, 0, 0)
		scene.Add(object)
	}

	// cube
	{
		points := []*js.Object{
			t.NewVector3(50, 50, 50).JSObject(),
			t.NewVector3(50, 50, -50).JSObject(),
			t.NewVector3(-50, 50, -50).JSObject(),
			t.NewVector3(-50, 50, 50).JSObject(),
			t.NewVector3(50, -50, 50).JSObject(),
			t.NewVector3(50, -50, -50).JSObject(),
			t.NewVector3(-50, -50, -50).JSObject(),
			t.NewVector3(-50, -50, 50).JSObject(),
		}

		geom := three.GeometryFromJSObject(t.JSObject().Get("ConvexGeometry").New(points))
		object := t.SceneUtils().CreateMultiMaterialObject(geom, materials)
		object.Position().Set(-200, 0, -200)
		scene.Add(object)
	}

	// random convex
	{
		points := []*js.Object{}
		for i := 0; i < 30; i++ {
			points = append(points, randomPointInSphere(50).JSObject())
		}

		geom := three.GeometryFromJSObject(t.JSObject().Get("ConvexGeometry").New(points))
		object := t.SceneUtils().CreateMultiMaterialObject(geom, materials)
		object.Position().Set(-200, 0, 200)
		scene.Add(object)
	}

	{
		object := t.NewAxisHelper(50)
		object.Position().Set(200, 0, -200)
		scene.Add(object)
	}

	{
		opts := &three.ArrowHelperOpts{Length: three.Float64(50)}
		object := t.NewArrowHelper(t.NewVector3(0, 1, 0), t.NewVector3(0, 0, 0), opts)
		object.Position().Set(200, 0, 400)
		scene.Add(object)
	}

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
		(rand.Float64()-0.5)*2*radius,
		(rand.Float64()-0.5)*2*radius,
		(rand.Float64()-0.5)*2*radius,
	)
}

func animate() {
	js.Global.Call("requestAnimationFrame", animate)

	render()
	stats.Call("update")
}

func render() {
	now := time.Now().UnixNano()
	timer := float64(now) / 1e10

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
