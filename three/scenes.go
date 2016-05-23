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

// Scene creates a new scene object.
//
// http://threejs.org/docs/index.html#Reference/Scenes/Scene
func (t *Three) Scene() *js.Object { return t.ctx.Get("Scene").New() }
