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

// SceneUtils returns a SceneUtils object.
func (t *Three) SceneUtils() *SceneUtils {
	p := t.ctx.Get("SceneUtils")
	return &SceneUtils{p: p}
}

// New returns a new SceneUtils object.
func (s *SceneUtils) New() *SceneUtils {
	p := s.p.New()
	return &SceneUtils{p: p}
}

// CreateMultiMaterialObject TODO description.
func (s *SceneUtils) CreateMultiMaterialObject(geometry, materials float64) *SceneUtils {
	s.p.Call("createMultiMaterialObject", geometry, materials)
	return s
}

// Detach TODO description.
func (s *SceneUtils) Detach(child, parent, scene float64) *SceneUtils {
	s.p.Call("detach", child, parent, scene)
	return s
}

// Attach TODO description.
func (s *SceneUtils) Attach(child, scene, parent float64) *SceneUtils {
	s.p.Call("attach", child, scene, parent)
	return s
}
