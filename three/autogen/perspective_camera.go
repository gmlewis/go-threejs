package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// PerspectiveCamera represents a perspectivecamera.
type PerspectiveCamera struct{ p *js.Object }

// PerspectiveCamera returns a PerspectiveCamera object.
func (t *Three) PerspectiveCamera() *PerspectiveCamera {
	p := t.ctx.Get("PerspectiveCamera")
	return &PerspectiveCamera{p: p}
}

// New returns a new PerspectiveCamera object.
func (t *PerspectiveCamera) New(fov, aspect, near, far float64) *PerspectiveCamera {
	p := t.p.New(fov, aspect, near, far)
	return &PerspectiveCamera{p: p}
}

// SetLens TODO description.
func (p *PerspectiveCamera) SetLens(focalLength, frameHeight float64) *PerspectiveCamera {
	p.p.Call("setLens", focalLength, frameHeight)
	return p
}

// SetViewOffset TODO description.
func (p *PerspectiveCamera) SetViewOffset(fullWidth, fullHeight, x, y, width, height float64) *PerspectiveCamera {
	p.p.Call("setViewOffset", fullWidth, fullHeight, x, y, width, height)
	return p
}

// Copy TODO description.
func (p *PerspectiveCamera) Copy(source float64) *PerspectiveCamera {
	p.p.Call("copy", source)
	return p
}

// ToJSON TODO description.
func (p *PerspectiveCamera) ToJSON(meta float64) *PerspectiveCamera {
	p.p.Call("toJSON", meta)
	return p
}

