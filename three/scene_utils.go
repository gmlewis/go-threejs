// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// SceneUtils contains useful utility functions for scene manipulation.
//
// http://threejs.org/docs/index.html#Reference/Extras/SceneUtils
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

// CreateMultiMaterialObject creates an new Object3D an new mesh for each material defined in materials.
// Beware that this is not the same as MultiMaterial which defines multiple material for 1 mesh.
// This is mostly useful for object that need a material and a wireframe implementation.
//
//     geometry -- The geometry for the Object.
//     materials -- The materials for the object.
func (s *SceneUtils) CreateMultiMaterialObject(geometry, materials float64) *SceneUtils {
	s.p.Call("createMultiMaterialObject", geometry, materials)
	return s
}

// Detach detaches the object from the parent and adds it back to the scene without moving in worldspace.
// Beware that to do this the matrixWorld needs to be updated, this can be done by calling the
// updateMatrixWorld method on the parent object.
//
//     child -- The object to remove from the parent
//     scene -- The scene to attach the object on.
//     parent -- The parent to detach the object from.
func (s *SceneUtils) Detach(child, parent, scene float64) *SceneUtils {
	s.p.Call("detach", child, parent, scene)
	return s
}

// Attach attaches the object to the parent without the moving the object in the worldspace.
// Beware that to do this the matrixWorld needs to be updated, this can be done by calling the
// updateMatrixWorld method on the parent object.
//
//     child -- The object to add to the parent
//     scene -- The scene to detach the object on.
//     parent -- The parent to attach the object from.
func (s *SceneUtils) Attach(child, scene, parent float64) *SceneUtils {
	s.p.Call("attach", child, scene, parent)
	return s
}
