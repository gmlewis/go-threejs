package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Spherical represents a spherical.
type Spherical struct{ p *js.Object }

// Spherical returns a spherical object.
func (t *Three) Spherical() *Spherical {
	p := t.ctx.Get("Spherical")
	return &Spherical{p: p}
}

// NewSpherical returns a new spherical object.
func (t *Spherical) New(radius, phi, theta float64) *Spherical {
	p := t.p.New(radius, phi, theta)
	return &Spherical{p: p}
}

// Set TODO description.
func (s *Spherical) Set(radius, phi, theta float64) *Spherical {
	s.p.Call("set", radius, phi, theta)
	return s
}

// Clone TODO description.
func (s *Spherical) Clone() *Spherical {
	s.p.Call("clone")
	return s
}

// Copy TODO description.
func (s *Spherical) Copy(other float64) *Spherical {
	s.p.Call("copy", other)
	return s
}

// MakeSafe TODO description.
func (s *Spherical) MakeSafe() *Spherical {
	s.p.Call("makeSafe")
	return s
}

// SetFromVector3 TODO description.
func (s *Spherical) SetFromVector3(vec3 float64) *Spherical {
	s.p.Call("setFromVector3", vec3)
	return s
}

