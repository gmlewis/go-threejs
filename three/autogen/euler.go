package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Euler represents an euler.
type Euler struct{ p *js.Object }

// Euler returns an Euler object.
func (t *Three) Euler() *Euler {
	p := t.ctx.Get("Euler")
	return &Euler{p: p}
}

// New returns a new Euler object.
func (t *Euler) New(x, y, z, order float64) *Euler {
	p := t.p.New(x, y, z, order)
	return &Euler{p: p}
}

// X returns the x-component of the Euler.
func (e *Euler) X() float64 {
	return e.p.Get("x").Float()
}

// Y returns the y-component of the Euler.
func (e *Euler) Y() float64 {
	return e.p.Get("y").Float()
}

// Z returns the z-component of the Euler.
func (e *Euler) Z() float64 {
	return e.p.Get("z").Float()
}

// Order returns the order-component of the Euler.
func (e *Euler) Order() float64 {
	return e.p.Get("order").Float()
}

// SetX sets the x-component of the Euler.
func (e *Euler) SetX(value float64) *Euler {
	e.p.Set("x", value)
	return e
}

// SetY sets the y-component of the Euler.
func (e *Euler) SetY(value float64) *Euler {
	e.p.Set("y", value)
	return e
}

// SetZ sets the z-component of the Euler.
func (e *Euler) SetZ(value float64) *Euler {
	e.p.Set("z", value)
	return e
}

// SetOrder sets the order-component of the Euler.
func (e *Euler) SetOrder(value float64) *Euler {
	e.p.Set("order", value)
	return e
}

// Set TODO description.
func (e *Euler) Set(x, y, z, order float64) *Euler {
	e.p.Call("set", x, y, z, order)
	return e
}

// Clone TODO description.
func (e *Euler) Clone() *Euler {
	e.p.Call("clone")
	return e
}

// Copy TODO description.
func (e *Euler) Copy(euler float64) *Euler {
	e.p.Call("copy", euler)
	return e
}

// SetFromRotationMatrix TODO description.
func (e *Euler) SetFromRotationMatrix(m, order, update float64) *Euler {
	e.p.Call("setFromRotationMatrix", m, order, update)
	return e
}

// SetFromQuaternion TODO description.
func (e *Euler) SetFromQuaternion() *Euler {
	e.p.Call("setFromQuaternion")
	return e
}

// SetFromVector3 TODO description.
func (e *Euler) SetFromVector3(v, order float64) *Euler {
	e.p.Call("setFromVector3", v, order)
	return e
}

// Reorder TODO description.
func (e *Euler) Reorder() *Euler {
	e.p.Call("reorder")
	return e
}

// Equals TODO description.
func (e *Euler) Equals(euler float64) *Euler {
	e.p.Call("equals", euler)
	return e
}

// FromArray TODO description.
func (e *Euler) FromArray(array float64) *Euler {
	e.p.Call("fromArray", array)
	return e
}

// ToArray TODO description.
func (e *Euler) ToArray(array, offset float64) *Euler {
	e.p.Call("toArray", array, offset)
	return e
}

// ToVector3 TODO description.
func (e *Euler) ToVector3(optionalResult float64) *Euler {
	e.p.Call("toVector3", optionalResult)
	return e
}

// OnChange TODO description.
func (e *Euler) OnChange(callback float64) *Euler {
	e.p.Call("onChange", callback)
	return e
}

// OnChangeCallback TODO description.
func (e *Euler) OnChangeCallback() *Euler {
	e.p.Call("onChangeCallback")
	return e
}
