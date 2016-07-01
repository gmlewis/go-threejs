// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Object3D is the JavaScript base class for scene graph objects.
//
// http://threejs.org/docs/index.html#Reference/Core/Object3D
type Object3D struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (o *Object3D) JSObject() *js.Object { return o.p }

// Object3D returns an Object3D JavaScript class.
func (t *Three) Object3D() *Object3D {
	p := t.ctx.Get("Object3D")
	return &Object3D{p: p}
}

// NewObject3D returns a new Object3D object.
func (t *Three) NewObject3D() *Object3D {
	p := t.ctx.Get("Object3D").New()
	return &Object3D{p: p}
}

// ApplyMatrix TODO description.
func (o *Object3D) ApplyMatrix(matrix float64) *Object3D {
	o.p.Call("applyMatrix", matrix)
	return o
}

// SetRotationFromAxisAngle TODO description.
func (o *Object3D) SetRotationFromAxisAngle(axis, angle float64) *Object3D {
	o.p.Call("setRotationFromAxisAngle", axis, angle)
	return o
}

// SetRotationFromEuler TODO description.
func (o *Object3D) SetRotationFromEuler(euler float64) *Object3D {
	o.p.Call("setRotationFromEuler", euler)
	return o
}

// SetRotationFromMatrix TODO description.
func (o *Object3D) SetRotationFromMatrix(m float64) *Object3D {
	o.p.Call("setRotationFromMatrix", m)
	return o
}

// SetRotationFromQuaternion TODO description.
func (o *Object3D) SetRotationFromQuaternion(q float64) *Object3D {
	o.p.Call("setRotationFromQuaternion", q)
	return o
}

// RotateOnAxis TODO description.
func (o *Object3D) RotateOnAxis() *Object3D {
	o.p.Call("rotateOnAxis")
	return o
}

// RotateX TODO description.
func (o *Object3D) RotateX() *Object3D {
	o.p.Call("rotateX")
	return o
}

// RotateY TODO description.
func (o *Object3D) RotateY() *Object3D {
	o.p.Call("rotateY")
	return o
}

// RotateZ TODO description.
func (o *Object3D) RotateZ() *Object3D {
	o.p.Call("rotateZ")
	return o
}

// TranslateOnAxis TODO description.
func (o *Object3D) TranslateOnAxis() *Object3D {
	o.p.Call("translateOnAxis")
	return o
}

// TranslateX TODO description.
func (o *Object3D) TranslateX() *Object3D {
	o.p.Call("translateX")
	return o
}

// TranslateY TODO description.
func (o *Object3D) TranslateY() *Object3D {
	o.p.Call("translateY")
	return o
}

// TranslateZ TODO description.
func (o *Object3D) TranslateZ() *Object3D {
	o.p.Call("translateZ")
	return o
}

// LocalToWorld TODO description.
func (o *Object3D) LocalToWorld(vector float64) *Object3D {
	o.p.Call("localToWorld", vector)
	return o
}

// WorldToLocal TODO description.
func (o *Object3D) WorldToLocal() *Object3D {
	o.p.Call("worldToLocal")
	return o
}

// LookAt TODO description.
func (o *Object3D) LookAt() *Object3D {
	o.p.Call("lookAt")
	return o
}

// Add TODO description.
func (o *Object3D) Add(object float64) *Object3D {
	o.p.Call("add", object)
	return o
}

// Remove TODO description.
func (o *Object3D) Remove(object float64) *Object3D {
	o.p.Call("remove", object)
	return o
}

// GetObjectByID TODO description.
func (o *Object3D) GetObjectByID(id float64) *Object3D {
	o.p.Call("getObjectById", id)
	return o
}

// GetObjectByName TODO description.
func (o *Object3D) GetObjectByName(name float64) *Object3D {
	o.p.Call("getObjectByName", name)
	return o
}

// GetObjectByProperty TODO description.
func (o *Object3D) GetObjectByProperty(name, value float64) *Object3D {
	o.p.Call("getObjectByProperty", name, value)
	return o
}

// GetWorldPosition TODO description.
func (o *Object3D) GetWorldPosition(optionalTarget float64) *Object3D {
	o.p.Call("getWorldPosition", optionalTarget)
	return o
}

// GetWorldQuaternion TODO description.
func (o *Object3D) GetWorldQuaternion() *Object3D {
	o.p.Call("getWorldQuaternion")
	return o
}

// GetWorldRotation TODO description.
func (o *Object3D) GetWorldRotation() *Object3D {
	o.p.Call("getWorldRotation")
	return o
}

// GetWorldScale TODO description.
func (o *Object3D) GetWorldScale() *Object3D {
	o.p.Call("getWorldScale")
	return o
}

// GetWorldDirection TODO description.
func (o *Object3D) GetWorldDirection() *Object3D {
	o.p.Call("getWorldDirection")
	return o
}

// Raycast TODO description.
func (o *Object3D) Raycast() *Object3D {
	o.p.Call("raycast")
	return o
}

// Traverse TODO description.
func (o *Object3D) Traverse(callback float64) *Object3D {
	o.p.Call("traverse", callback)
	return o
}

// TraverseVisible TODO description.
func (o *Object3D) TraverseVisible(callback float64) *Object3D {
	o.p.Call("traverseVisible", callback)
	return o
}

// TraverseAncestors TODO description.
func (o *Object3D) TraverseAncestors(callback float64) *Object3D {
	o.p.Call("traverseAncestors", callback)
	return o
}

// UpdateMatrix TODO description.
func (o *Object3D) UpdateMatrix() *Object3D {
	o.p.Call("updateMatrix")
	return o
}

// UpdateMatrixWorld TODO description.
func (o *Object3D) UpdateMatrixWorld(force float64) *Object3D {
	o.p.Call("updateMatrixWorld", force)
	return o
}

// ToJSON TODO description.
func (o *Object3D) ToJSON(meta float64) *Object3D {
	o.p.Call("toJSON", meta)
	return o
}

// Clone TODO description.
func (o *Object3D) Clone(recursive float64) *Object3D {
	o.p.Call("clone", recursive)
	return o
}

// Copy TODO description.
func (o *Object3D) Copy(source *Object3D, recursive bool) *Object3D {
	o.p.Call("copy", source.p, recursive)
	return o
}

// UUID returns the property of the same name.
func (o *Object3D) UUID() int {
	return o.p.Get("uuid").Int()
}

// Name returns the property of the same name.
func (o *Object3D) Name() string {
	return o.p.Get("name").String()
}

// Type returns the property of the same name.
func (o *Object3D) Type() string {
	return o.p.Get("type").String()
}

// Parent returns the property of the same name.
func (o *Object3D) Parent() *js.Object {
	return o.p.Get("parent")
}

// Children returns the property of the same name.
func (o *Object3D) Children() *js.Object {
	return o.p.Get("children")
}

// Up returns the property of the same name.
func (o *Object3D) Up() *Vector3 {
	return &Vector3{p: o.p.Get("up")}
}

// Position returns the property of the same name.
func (o *Object3D) Position() *Vector3 {
	return &Vector3{p: o.p.Get("position")}
}

// Rotation returns the property of the same name.
func (o *Object3D) Rotation() *Euler {
	return &Euler{p: o.p.Get("rotation")}
}

// Quaternion returns the property of the same name.
func (o *Object3D) Quaternion() *Quaternion {
	return &Quaternion{p: o.p.Get("quaternion")}
}

// Scale returns the property of the same name.
func (o *Object3D) Scale() *Vector3 {
	return &Vector3{p: o.p.Get("scale")}
}

// ModelViewMatrix returns the property of the same name.
func (o *Object3D) ModelViewMatrix() *Matrix4 {
	return &Matrix4{p: o.p.Get("modelViewMatrix")}
}

// NormalMatrix returns the property of the same name.
func (o *Object3D) NormalMatrix() *Matrix3 {
	return &Matrix3{p: o.p.Get("normalMatrix")}
}

/* TODO:
this.matrixAutoUpdate = THREE.Object3D.DefaultMatrixAutoUpdate;

this.userData = {};
*/

// Matrix returns the property of the same name.
func (o *Object3D) Matrix() *Matrix4 {
	return &Matrix4{p: o.p.Get("matrix")}
}

// SetMatrix sets the matrix property.
func (o *Object3D) SetMatrix(value *Matrix4) *Object3D {
	o.p.Set("matrix", value.p)
	return o
}

// MatrixWorld returns the property of the same name.
func (o *Object3D) MatrixWorld() *Matrix4 {
	return &Matrix4{p: o.p.Get("matrixWorld")}
}

// SetMatrixWorld sets the matrixWorld property.
func (o *Object3D) SetMatrixWorld(value *Matrix4) *Object3D {
	o.p.Set("matrixWorld", value.p)
	return o
}

// RotationAutoUpdate returns the property of the same name.
func (o *Object3D) RotationAutoUpdate() bool {
	return o.p.Get("rotationAutoUpdate").Bool()
}

// SetRotationAutoUpdate sets the rotationAutoUpdate property.
func (o *Object3D) SetRotationAutoUpdate(value bool) *Object3D {
	o.p.Set("rotationAutoUpdate", value)
	return o
}

// MatrixWorldNeedsUpdate returns the property of the same name.
func (o *Object3D) MatrixWorldNeedsUpdate() bool {
	return o.p.Get("matrixWorldNeedsUpdate").Bool()
}

// SetMatrixWorldNeedsUpdate sets the matrixWorldNeedsUpdate property.
func (o *Object3D) SetMatrixWorldNeedsUpdate(value bool) *Object3D {
	o.p.Set("matrixWorldNeedsUpdate", value)
	return o
}

// Layers returns the property of the same name.
func (o *Object3D) Layers() *Layers {
	return &Layers{p: o.p.Get("layers")}
}

// SetLayers sets the layers property.
func (o *Object3D) SetLayers(value *Layers) *Object3D {
	o.p.Set("layers", value.p)
	return o
}

// Visible returns the property of the same name.
func (o *Object3D) Visible() bool {
	return o.p.Get("visible").Bool()
}

// SetVisible sets the visible property.
func (o *Object3D) SetVisible(value bool) *Object3D {
	o.p.Set("visible", value)
	return o
}

// CastShadow returns the property of the same name.
func (o *Object3D) CastShadow() bool {
	return o.p.Get("castShadow").Bool()
}

// SetCastShadow sets the castShadow property.
func (o *Object3D) SetCastShadow(value bool) *Object3D {
	o.p.Set("castShadow", value)
	return o
}

// ReceiveShadow returns the property of the same name.
func (o *Object3D) ReceiveShadow() bool {
	return o.p.Get("receiveShadow").Bool()
}

// SetReceiveShadow sets the receiveShadow property.
func (o *Object3D) SetReceiveShadow(value bool) *Object3D {
	o.p.Set("receiveShadow", value)
	return o
}

// FrustumCulled returns the property of the same name.
func (o *Object3D) FrustumCulled() bool {
	return o.p.Get("frustumCulled").Bool()
}

// SetFrustumCulled sets the frustumCulled property.
func (o *Object3D) SetFrustumCulled(value bool) *Object3D {
	o.p.Set("frustumCulled", value)
	return o
}

// RenderOrder returns the property of the same name.
func (o *Object3D) RenderOrder() int {
	return o.p.Get("renderOrder").Int()
}

// SetRenderOrder sets the renderOrder property.
func (o *Object3D) SetRenderOrder(value int) *Object3D {
	o.p.Set("renderOrder", value)
	return o
}
