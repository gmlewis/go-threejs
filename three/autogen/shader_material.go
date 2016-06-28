package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// ShaderMaterial represents a shadermaterial.
type ShaderMaterial struct{ p *js.Object }

// ShaderMaterial returns a shadermaterial object.
func (t *Three) ShaderMaterial() *ShaderMaterial {
	p := t.ctx.Get("ShaderMaterial")
	return &ShaderMaterial{p: p}
}

// NewShaderMaterial returns a new shadermaterial object.
func (t *ShaderMaterial) New(parameters float64) *ShaderMaterial {
	p := t.p.New(parameters)
	return &ShaderMaterial{p: p}
}

// Copy TODO description.
func (s *ShaderMaterial) Copy(source float64) *ShaderMaterial {
	s.p.Call("copy", source)
	return s
}

// ToJSON TODO description.
func (s *ShaderMaterial) ToJSON(meta float64) *ShaderMaterial {
	s.p.Call("toJSON", meta)
	return s
}

