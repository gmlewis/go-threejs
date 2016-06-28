package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// LensFlare represents a lensflare.
type LensFlare struct{ p *js.Object }

// LensFlare returns a LensFlare object.
func (t *Three) LensFlare() *LensFlare {
	p := t.ctx.Get("LensFlare")
	return &LensFlare{p: p}
}

// New returns a new LensFlare object.
func (t *LensFlare) New(texture, size, distance, blending, color float64) *LensFlare {
	p := t.p.New(texture, size, distance, blending, color)
	return &LensFlare{p: p}
}

// Add TODO description.
func (l *LensFlare) Add(texture, size, distance, blending, color, opacity float64) *LensFlare {
	l.p.Call("add", texture, size, distance, blending, color, opacity)
	return l
}

// Copy TODO description.
func (l *LensFlare) Copy(source float64) *LensFlare {
	l.p.Call("copy", source)
	return l
}

