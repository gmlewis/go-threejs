package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// UniformsLib represents an uniformslib.
type UniformsLib struct{ p *js.Object }

// UniformsLib returns an uniformslib object.
func (t *Three) UniformsLib() *UniformsLib {
	p := t.ctx.Get("UniformsLib")
	return &UniformsLib{p: p}
}

// NewUniformsLib returns a new uniformslib object.
func (t *UniformsLib) New() *UniformsLib {
	p := t.p.New()
	return &UniformsLib{p: p}
}

