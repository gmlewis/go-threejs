package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// UniformsUtils represents an uniformsutils.
type UniformsUtils struct{ p *js.Object }

// UniformsUtils returns an uniformsutils object.
func (t *Three) UniformsUtils() *UniformsUtils {
	p := t.ctx.Get("UniformsUtils")
	return &UniformsUtils{p: p}
}

// NewUniformsUtils returns a new uniformsutils object.
func (t *UniformsUtils) New() *UniformsUtils {
	p := t.p.New()
	return &UniformsUtils{p: p}
}

// Merge TODO description.
func (u *UniformsUtils) Merge(uniforms float64) *UniformsUtils {
	u.p.Call("merge", uniforms)
	return u
}

// Clone TODO description.
func (u *UniformsUtils) Clone(uniforms_src float64) *UniformsUtils {
	u.p.Call("clone", uniforms_src)
	return u
}

