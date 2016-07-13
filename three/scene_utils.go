// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// SceneUtils represents a sceneutils.
type SceneUtils struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (s *SceneUtils) JSObject() *js.Object { return s.p }

// SceneUtils returns a SceneUtils JavaScript class.
func (t *Three) SceneUtils() *SceneUtils {
	p := t.ctx.Get("SceneUtils")
	return &SceneUtils{p: p}
}

// NewSceneUtils returns a new SceneUtils object.
func (t *Three) NewSceneUtils() *SceneUtils {
	p := t.ctx.Get("SceneUtils").New()
	return &SceneUtils{p: p}
}

// CreateMultiMaterialObject TODO description.
func (s *SceneUtils) CreateMultiMaterialObject(geometry JSObject, materials []JSObject) *Group {
	array := []*js.Object{}
	for i := 0; i < len(materials); i++ {
		array = append(array, materials[i].JSObject())
	}
	return GroupFromJSObject(s.p.Call("createMultiMaterialObject", geometry.JSObject(), array))
}

// Detach TODO description.
func (s *SceneUtils) Detach(child, parent, scene JSObject) *SceneUtils {
	s.p.Call("detach", child, parent, scene)
	return s
}

// Attach TODO description.
func (s *SceneUtils) Attach(child, scene, parent JSObject) *SceneUtils {
	s.p.Call("attach", child, scene, parent)
	return s
}
