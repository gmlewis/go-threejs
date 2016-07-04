// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Scene represents a three.js scene.
type Scene struct{ *Object3D }

// JSObject returns the underlying *js.Object.
func (s *Scene) JSObject() *js.Object { return s.p }

// Scene returns a Scene JavaScript class.
func (t *Three) Scene() *Scene {
	p := t.ctx.Get("Scene")
	return scene(p)
}

// scene returns a wrapped Scene JavaScript class.
func scene(p *js.Object) *Scene {
	return &Scene{object3D(p)}
}

// NewScene returns a new Scene object.
//
// http://threejs.org/docs/index.html#Reference/Scenes/Scene
func (t *Three) NewScene() *Scene {
	p := t.ctx.Get("Scene").New()
	return scene(p)
}

// Copy TODO description.
func (s *Scene) Copy(source *Scene, recursive bool) *Scene {
	s.p.Call("copy", source.p, recursive)
	return s
}

// Add adds an object to a scene.
func (s *Scene) Add(obj JSObject) *Scene {
	s.p.Call("add", obj.JSObject())
	return s
}
