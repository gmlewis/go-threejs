package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// AudioBuffer represents an audiobuffer.
type AudioBuffer struct{ p *js.Object }

// AudioBuffer returns an audiobuffer object.
func (t *Three) AudioBuffer() *AudioBuffer {
	p := t.ctx.Get("AudioBuffer")
	return &AudioBuffer{p: p}
}

// NewAudioBuffer returns a new audiobuffer object.
func (t *AudioBuffer) New(context float64) *AudioBuffer {
	p := t.p.New(context)
	return &AudioBuffer{p: p}
}

// Load TODO description.
func (a *AudioBuffer) Load(file float64) *AudioBuffer {
	a.p.Call("load", file)
	return a
}

// OnReady TODO description.
func (a *AudioBuffer) OnReady(callback float64) *AudioBuffer {
	a.p.Call("onReady", callback)
	return a
}

