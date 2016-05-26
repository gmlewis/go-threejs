package three

import "github.com/gopherjs/gopherjs/js"

// FogExp2 creates exponential fog.
//
// http://threejs.org/docs/index.html#Reference/Scenes/FogExp2
func (t *Three) FogExp2(hex int, density float64) *js.Object {
	return t.ctx.Get("FogExp2").New(hex, density)
}

// Fog creates linear fog.
//
// http://threejs.org/docs/index.html#Reference/Scenes/Fog
func (t *Three) Fog(hex int, near, far float64) *js.Object {
	return t.ctx.Get("Fog").New(hex, near, far)
}

// Scene represents a three.js Scene.
type Scene struct {
	obj *js.Object
}

// Scene creates a new scene object.
//
// http://threejs.org/docs/index.html#Reference/Scenes/Scene
func (t *Three) Scene() *Scene { return &Scene{obj: t.ctx.Get("Scene").New()} }

// Add adds an objects to a scene.
func (s *Scene) Add(obj interface{}) {
	s.obj.Call("add", obj)
}
