package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Sprite represents a sprite.
type Sprite struct{ p *js.Object }

// Sprite returns a Sprite object.
func (t *Three) Sprite() *Sprite {
	p := t.ctx.Get("Sprite")
	return &Sprite{p: p}
}

// New returns a new Sprite object.
func (t *Sprite) New() *Sprite {
	p := t.p.New()
	return &Sprite{p: p}
}

