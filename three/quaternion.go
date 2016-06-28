package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Quaternion represents a quaternion.
type Quaternion js.Object

// NewQuaternion returns a new quaternion.
func (t *Three) NewQuaternion(x, y, z, w float64) *Quaternion {
	return (*Quaternion)(t.ctx.Get("Quaternion").New(x, y, z, w))
}

// X returns the x-component of the quaternion.
func (q *Quaternion) X() float64 {
	return (*js.Object)(q).Get("_x").Float()
}

// Y returns the x-component of the quaternion.
func (q *Quaternion) Y() float64 {
	return (*js.Object)(q).Get("_y").Float()
}

// Z returns the x-component of the quaternion.
func (q *Quaternion) Z() float64 {
	return (*js.Object)(q).Get("_z").Float()
}

// W returns the x-component of the quaternion.
func (q *Quaternion) W() float64 {
	return (*js.Object)(q).Get("_w").Float()
}

// Set sets the components of the quaternion.
func (q *Quaternion) Set(x, y, z, w float64) *Quaternion {
	(*js.Object)(q).Call("set", x, y, z, w)
	return q
}

// Clone returns a clone of this quaternion.
func (q *Quaternion) Clone() *Quaternion {
	v := (*js.Object)(q).Call("clone")
	return (*Quaternion)(v)
}

// Copy copies the values from src into q.
func (q *Quaternion) Copy(src *Quaternion) *Quaternion {
	v := (*js.Object)(q).Call("copy", (*js.Object)(src))
	return (*Quaternion)(v)
}
