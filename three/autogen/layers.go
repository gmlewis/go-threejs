package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Layers represents a layers.
type Layers struct{ p *js.Object }

// Layers returns a layers object.
func (t *Three) Layers() *Layers {
	p := t.ctx.Get("Layers")
	return &Layers{p: p}
}

// NewLayers returns a new layers object.
func (t *Layers) New() *Layers {
	p := t.p.New()
	return &Layers{p: p}
}

// Set TODO description.
func (l *Layers) Set(channel float64) *Layers {
	l.p.Call("set", channel)
	return l
}

// Enable TODO description.
func (l *Layers) Enable(channel float64) *Layers {
	l.p.Call("enable", channel)
	return l
}

// Toggle TODO description.
func (l *Layers) Toggle(channel float64) *Layers {
	l.p.Call("toggle", channel)
	return l
}

// Disable TODO description.
func (l *Layers) Disable(channel float64) *Layers {
	l.p.Call("disable", channel)
	return l
}

// Test TODO description.
func (l *Layers) Test(layers float64) *Layers {
	l.p.Call("test", layers)
	return l
}
