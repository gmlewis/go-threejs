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

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Scene represents a three.js scene. Scenes allow you to set up what and where is to be
// rendered by three.js. This is where you place objects, lights and cameras.
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
