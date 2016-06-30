// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// AudioBuffer represents an audiobuffer.
type AudioBuffer struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (a *AudioBuffer) JSObject() *js.Object { return a.p }

// AudioBuffer returns an AudioBuffer object.
func (t *Three) AudioBuffer() *AudioBuffer {
	p := t.ctx.Get("AudioBuffer")
	return &AudioBuffer{p: p}
}

// New returns a new AudioBuffer object.
func (a *AudioBuffer) New(context float64) *AudioBuffer {
	p := a.p.New(context)
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
