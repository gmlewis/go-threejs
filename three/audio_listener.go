package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// AudioListener represents an audiolistener.
type AudioListener struct{ p *js.Object }

// AudioListener returns an AudioListener object.
func (t *Three) AudioListener() *AudioListener {
	p := t.ctx.Get("AudioListener")
	return &AudioListener{p: p}
}

// New returns a new AudioListener object.
func (t *AudioListener) New() *AudioListener {
	p := t.p.New()
	return &AudioListener{p: p}
}

// RemoveFilter TODO description.
func (a *AudioListener) RemoveFilter(  float64) *AudioListener {
	a.p.Call("removeFilter",  )
	return a
}

// SetFilter TODO description.
func (a *AudioListener) SetFilter(value float64) *AudioListener {
	a.p.Call("setFilter", value)
	return a
}

// SetMasterVolume TODO description.
func (a *AudioListener) SetMasterVolume(value float64) *AudioListener {
	a.p.Call("setMasterVolume", value)
	return a
}

