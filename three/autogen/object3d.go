package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Object3D represents an object3d.
type Object3D struct{ p *js.Object }

// Object3D returns an object3d object.
func (t *Three) Object3D() *Object3D {
	p := t.ctx.Get("Object3D")
	return &Object3D{p: p}
}

// NewObject3D returns a new object3d object.
func (t *Object3D) New() *Object3D {
	p := t.p.New()
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

// GetObjectById TODO description.
func (o *Object3D) GetObjectById(id float64) *Object3D {
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
func (o *Object3D) Copy(source, recursive float64) *Object3D {
	o.p.Call("copy", source, recursive)
	return o
}

