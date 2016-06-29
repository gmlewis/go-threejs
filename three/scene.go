package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Scene represents a three.js scene.
type Scene struct{ p *js.Object }

// Scene returns a Scene object.
func (t *Three) Scene() *Scene {
	p := t.ctx.Get("Scene")
	return &Scene{p: p}
}

// New returns a new Scene object.
//
// http://threejs.org/docs/index.html#Reference/Scenes/Scene
func (t *Scene) New() *Scene {
	p := t.p.New()
	return &Scene{p: p}
}

// Copy TODO description.
func (s *Scene) Copy(source, recursive float64) *Scene {
	s.p.Call("copy", source, recursive)
	return s
}

// Add adds an object to a scene.
func (s *Scene) Add(obj *js.Object) *Scene {
	s.p.Call("add", obj)
	return s
}
