package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// SpriteMaterial represents a material for a sprite.
//
// http://threejs.org/docs/index.html#Reference/Materials/SpriteMaterial
type SpriteMaterial struct{ p *js.Object }

// SpriteMaterial returns a SpriteMaterial object.
func (t *Three) SpriteMaterial() *SpriteMaterial {
	p := t.ctx.Get("SpriteMaterial")
	return &SpriteMaterial{p: p}
}

// New returns a new SpriteMaterial object.
//
// parameters is an object defining the the default properties:
//     color - color of the sprite
//     map - the texture map
//     rotation - the rotation of the sprite
//     fog - whether or not to use the scene fog
func (t *SpriteMaterial) New(parameters map[string]interface{}) *SpriteMaterial {
	p := t.p.New(parameters)
	return &SpriteMaterial{p: p}
}

// Copy TODO description.
func (s *SpriteMaterial) Copy(source *SpriteMaterial) *SpriteMaterial {
	s.p.Call("copy", source.p)
	return s
}
