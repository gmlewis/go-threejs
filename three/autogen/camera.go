package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Camera represents a camera.
type Camera struct{ p *js.Object }

// Camera returns a camera object.
func (t *Three) Camera() *Camera {
	p := t.ctx.Get("Camera")
	return &Camera{p: p}
}

// NewCamera returns a new camera object.
func (t *Camera) New() *Camera {
	p := t.p.New()
	return &Camera{p: p}
}

// Copy TODO description.
func (c *Camera) Copy(source float64) *Camera {
	c.p.Call("copy", source)
	return c
}

