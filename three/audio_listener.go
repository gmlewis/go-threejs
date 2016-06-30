// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// AudioListener represents an audiolistener.
type AudioListener struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (a *AudioListener) JSObject() *js.Object { return a.p }

// AudioListener returns an AudioListener object.
func (t *Three) AudioListener() *AudioListener {
	p := t.ctx.Get("AudioListener")
	return &AudioListener{p: p}
}

// New returns a new AudioListener object.
func (a *AudioListener) New() *AudioListener {
	p := a.p.New()
	return &AudioListener{p: p}
}

// RemoveFilter TODO description.
func (a *AudioListener) RemoveFilter(float64) *AudioListener {
	a.p.Call("removeFilter")
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
