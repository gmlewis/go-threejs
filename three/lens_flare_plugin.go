package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// LensFlarePlugin represents a lensflareplugin.
type LensFlarePlugin struct{ p *js.Object }

// LensFlarePlugin returns a LensFlarePlugin object.
func (t *Three) LensFlarePlugin() *LensFlarePlugin {
	p := t.ctx.Get("LensFlarePlugin")
	return &LensFlarePlugin{p: p}
}

// New returns a new LensFlarePlugin object.
func (t *LensFlarePlugin) New(renderer, flares float64) *LensFlarePlugin {
	p := t.p.New(renderer, flares)
	return &LensFlarePlugin{p: p}
}

