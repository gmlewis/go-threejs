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

// SceneUtils contains useful utility functions for scene manipulation.
//
// http://threejs.org/docs/index.html#Reference/Extras/SceneUtils
type SceneUtils struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (s *SceneUtils) JSObject() *js.Object { return s.p }

// SceneUtils returns a SceneUtils JavaScript class.
func (t *Three) SceneUtils() *SceneUtils {
	p := t.ctx.Get("SceneUtils")
	return SceneUtilsFromJSObject(p)
}

// SceneUtilsFromJSObject returns a wrapped SceneUtils JavaScript class.
func SceneUtilsFromJSObject(p *js.Object) *SceneUtils {
	return &SceneUtils{p: p}
}

// NewSceneUtils returns a new SceneUtils object.
func (t *Three) NewSceneUtils() *SceneUtils {
	p := t.ctx.Get("SceneUtils").New()
	return SceneUtilsFromJSObject(p)
}

// CreateMultiMaterialObject creates an new Object3D an new mesh for each material defined in materials.
// Beware that this is not the same as MultiMaterial which defines multiple material for 1 mesh.
// This is mostly useful for object that need a material and a wireframe implementation.
//
//     geometry -- The geometry for the Object.
//     materials -- The materials for the object.
func (s *SceneUtils) CreateMultiMaterialObject(geometry JSObject, materials []JSObject) *Group {
	array := []*js.Object{}
	for i := 0; i < len(materials); i++ {
		array = append(array, materials[i].JSObject())
	}
	return GroupFromJSObject(s.p.Call("createMultiMaterialObject", geometry.JSObject(), array))
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
