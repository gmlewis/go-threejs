// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Scene represents a three.js scene.
type Scene struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *Scene) JSObject() *js.Object { return t.p }

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
func (s *Scene) Add(obj JSObject) *Scene {
	s.p.Call("add", obj.JSObject())
	return s
}
