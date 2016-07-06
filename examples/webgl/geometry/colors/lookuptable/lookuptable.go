// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// See: http://threejs.org/examples/#webgl_geometry_colors_lookuptable
package main

import (
	"fmt"
	"math"
	"strconv"

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

	camera       *three.PerspectiveCamera
	scene        *three.Scene
	renderer     *three.WebGLRenderer
	lut          *three.LUT
	legendLayout = "vertical"

	position *three.Vector3

	rotWorldMatrix *three.Matrix4

	mesh           *three.Mesh
	colorMap       = "rainbow"
	numberOfColors = 512

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
	camera.SetName("camera")

	scene = t.NewScene()
	scene.Add(camera)

	stats = js.Global.Get("Stats").New()
	d := stats.Get("domElement")
	style := d.Get("style")
	style.Set("position", "absolute")
	style.Set("bottom", "0px")
	style.Set("zIndex", 100)
	container.Call("appendChild", d)

	ambientLight := t.NewAmbientLight(0x444444, 1)
	ambientLight.SetName("ambientLight")
	scene.Add(ambientLight)

	loadModel(colorMap, numberOfColors, legendLayout)

	camera.Position().Set(17, 9, 32)

	directionalLight := t.NewDirectionalLight(0xffffff, 0.7)
	directionalLight.SetName("directionalLight").Position().Set(17, 9, 30)
	scene.Add(directionalLight)

	renderer = t.NewWebGLRenderer(&three.WebGLRendererOpts{Antialias: three.Bool(true)})
	renderer.SetClearColor(t.NewColor(0xffffff), 1)
	renderer.SetPixelRatio(window.Get("devicePixelRatio").Float())
	renderer.SetSize(window.Get("innerWidth").Float(), window.Get("innerHeight").Float(), true)
	container.Call("appendChild", renderer.DOMElement())

	window.Call("addEventListener", "resize", onWindowResize, false)

	window.Call("addEventListener", "keydown", onKeyDown, true)
}

func main() {
	animate()
}

func rotateAroundWorldAxis(object *three.Mesh, axis *three.Vector3, radians float64) {
	if axis == nil {
		return
	}
	t := three.New()
	rotWorldMatrix = t.NewMatrix4()
	rotWorldMatrix.MakeRotationAxis(axis.Normalize(), radians)
	rotWorldMatrix.Multiply(object.Matrix())

	object.SetMatrix(rotWorldMatrix)
	object.Rotation().SetFromRotationMatrix(object.Matrix(), "", true)
}

func onWindowResize() {
	windowX := window.Get("innerWidth").Float()
	windowY := window.Get("innerHeight").Float()

	camera.SetAspect(windowX / windowY)
	camera.UpdateProjectionMatrix()

	renderer.SetSize(windowX, windowY, true)
	render()
}

func animate() {
	js.Global.Call("requestAnimationFrame", animate)

	render()
	stats.Call("update")
}

func render() {
	rotateAroundWorldAxis(mesh, position, math.Pi/180)

	renderer.Render(scene, camera, nil)
}

func loadModel(colorMap string, numberOfColors int, legendLayout string) {
	t := three.New()
	loader := t.NewBufferGeometryLoader()

	loader.Load("http://threejs.org/examples/models/json/pressure.json", func(geometry *three.BufferGeometry, materials []*three.Material) {

		geometry.ComputeVertexNormals()
		geometry.NormalizeNormals()

		material := t.NewMeshLambertMaterial(three.MeshLambertMaterialOpts{
			"side":         three.DoubleSide,
			"color":        0xF5F5F5,
			"vertexColors": three.VertexColors,
		})

		lut = t.NewLUT(colorMap, numberOfColors)

		lut.SetMax(2000).SetMin(0)

		pressure := geometry.Attributes().Get("pressure").Get("array")
		length := pressure.Get("length").Int()
		lutColors := make([]float32, 3*length)
		for i := 0; i < length; i++ {
			colorValue := pressure.Index(i).Float()
			if color := lut.GetColor(colorValue); color == nil {
				println("ERROR: ", colorValue)
			} else {
				lutColors[3*i] = float32(color.R())
				lutColors[3*i+1] = float32(color.G())
				lutColors[3*i+2] = float32(color.B())
			}
		}

		geometry.AddAttribute("color", t.NewBufferAttribute(lutColors, 3).JSObject())

		mesh = t.NewMesh(geometry, material)

		geometry.ComputeBoundingBox()
		center := geometry.BoundingBox().Center(nil)

		if position == nil {
			position = t.NewVector3(center.X(), center.Y(), center.Z())
		}

		scene.Add(mesh)

		if legendLayout != "" {
			var legend *three.Mesh
			if legendLayout == "horizontal" {
				legend = lut.SetLegendOn(three.LUTLegendOpts{"layout": "horizontal", "position": map[string]interface{}{"x": 21, "y": 6, "z": 5}})
			} else {
				legend = lut.SetLegendOn(nil)
			}

			scene.Add(legend)

			labels := lut.SetLegendLabels(three.LUTLegendOpts{"title": "Pressure", "um": "Pa", "ticks": 5}, nil)

			scene.Add(labels.Get("title"))

			ticks := labels.Get("ticks")
			for i := 0; i < len(js.Keys(ticks)); i++ {
				scene.Add(ticks.Index(i))
				scene.Add(labels.Get("lines").Index(i))
			}
		}
	}, nil, nil)
}

func cleanScene() {
	children := scene.Children()
	elementsInTheScene := len(children)
	for i := elementsInTheScene - 1; i >= 0; i-- {
		child := children[i]
		if child.Name() != "camera" &&
			child.Name() != "ambientLight" &&
			child.Name() != "directionalLight" {
			scene.Remove(child)
		}
	}
}

func indexOf(slice []string, s string) (int, bool) {
	for i := 0; i < len(slice); i++ {
		if slice[i] == s {
			return i, true
		}
	}
	return -1, false
}

func onKeyDown(e *js.Object) bool {
	maps := []string{"rainbow", "cooltowarm", "blackbody", "grayscale"}
	colorNumbers := []string{"16", "128", "256", "512"}

	switch e.Get("keyCode").Int() {
	case 65:
		cleanScene()
		index := 0
		if i, ok := indexOf(maps, colorMap); ok && i < len(maps)-1 {
			index = i + 1
		}
		colorMap = maps[index]
		loadModel(colorMap, numberOfColors, legendLayout)
	case 83:
		cleanScene()
		index := 0
		if i, ok := indexOf(colorNumbers, strconv.Itoa(numberOfColors)); ok && i < len(colorNumbers)-1 {
			index = i + 1
		}
		n, err := strconv.Atoi(colorNumbers[index])
		if err != nil {
			println(fmt.Sprintf("error parsing number %v: %v", colorNumbers[index], err))
			n = 512
		}
		numberOfColors = n
		loadModel(colorMap, numberOfColors, legendLayout)
	case 68:
		if legendLayout == "" {
			cleanScene()
			legendLayout = "vertical"
			loadModel(colorMap, numberOfColors, legendLayout)
		} else {
			cleanScene()
			legendLayout = lut.SetLegendOff()
			loadModel(colorMap, numberOfColors, legendLayout)
		}
	case 70:
		cleanScene()
		if legendLayout == "" {
			return false
		}
		lut.SetLegendOff()
		if legendLayout == "horizontal" {
			legendLayout = "vertical"
		} else {
			legendLayout = "horizontal"
		}
		loadModel(colorMap, numberOfColors, legendLayout)
	}
	return true
}
