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
	return SceneFromJSObject(p)
}

// SceneFromJSObject returns a wrapped Scene JavaScript class.
func SceneFromJSObject(p *js.Object) *Scene {
	return &Scene{Object3DFromJSObject(p)}
}

// NewScene returns a new Scene object.
//
// http://threejs.org/docs/index.html#Reference/Scenes/Scene
func (t *Three) NewScene() *Scene {
	p := t.ctx.Get("Scene").New()
	return SceneFromJSObject(p)
}

// Copy TODO description.
func (s *Scene) Copy(source *Scene, recursive bool) *Scene {
	s.p.Call("copy", source.p, recursive)
	return s
}
