package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// SpriteMaterial represents a spritematerial.
type SpriteMaterial struct{ p *js.Object }

// SpriteMaterial returns a spritematerial object.
func (t *Three) SpriteMaterial() *SpriteMaterial {
	p := t.ctx.Get("SpriteMaterial")
	return &SpriteMaterial{p: p}
}

// NewSpriteMaterial returns a new spritematerial object.
func (t *SpriteMaterial) New(parameters float64) *SpriteMaterial {
	p := t.p.New(parameters)
	return &SpriteMaterial{p: p}
}

// Copy TODO description.
func (s *SpriteMaterial) Copy(source float64) *SpriteMaterial {
	s.p.Call("copy", source)
	return s
}

