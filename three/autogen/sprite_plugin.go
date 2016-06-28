package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// SpritePlugin represents a spriteplugin.
type SpritePlugin struct{ p *js.Object }

// SpritePlugin returns a SpritePlugin object.
func (t *Three) SpritePlugin() *SpritePlugin {
	p := t.ctx.Get("SpritePlugin")
	return &SpritePlugin{p: p}
}

// New returns a new SpritePlugin object.
func (t *SpritePlugin) New(renderer, sprites float64) *SpritePlugin {
	p := t.p.New(renderer, sprites)
	return &SpritePlugin{p: p}
}

