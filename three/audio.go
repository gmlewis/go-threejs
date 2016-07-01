// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Audio represents an audio.
type Audio struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (a *Audio) JSObject() *js.Object { return a.p }

// Audio returns an Audio JavaScript class.
func (t *Three) Audio() *Audio {
	p := t.ctx.Get("Audio")
	return &Audio{p: p}
}

// NewAudio returns a new Audio object.
func (t *Three) NewAudio(listener float64) *Audio {
	p := t.ctx.Get("Audio").New(listener)
	return &Audio{p: p}
}

// Load TODO description.
func (a *Audio) Load(file float64) *Audio {
	a.p.Call("load", file)
	return a
}

// SetNodeSource TODO description.
func (a *Audio) SetNodeSource(audioNode float64) *Audio {
	a.p.Call("setNodeSource", audioNode)
	return a
}

// SetBuffer TODO description.
func (a *Audio) SetBuffer(audioBuffer float64) *Audio {
	a.p.Call("setBuffer", audioBuffer)
	return a
}

// SetFilter TODO description.
func (a *Audio) SetFilter(value float64) *Audio {
	a.p.Call("setFilter", value)
	return a
}

// SetPlaybackRate TODO description.
func (a *Audio) SetPlaybackRate(value float64) *Audio {
	a.p.Call("setPlaybackRate", value)
	return a
}

// SetLoop TODO description.
func (a *Audio) SetLoop(value float64) *Audio {
	a.p.Call("setLoop", value)
	return a
}

// SetVolume TODO description.
func (a *Audio) SetVolume(value float64) *Audio {
	a.p.Call("setVolume", value)
	return a
}
