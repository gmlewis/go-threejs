// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// AudioAnalyser represents an audioanalyser.
type AudioAnalyser struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (a *AudioAnalyser) JSObject() *js.Object { return a.p }

// AudioAnalyser returns an AudioAnalyser object.
func (t *Three) AudioAnalyser() *AudioAnalyser {
	p := t.ctx.Get("AudioAnalyser")
	return &AudioAnalyser{p: p}
}

// New returns a new AudioAnalyser object.
func (a *AudioAnalyser) New(audio, fftSize float64) *AudioAnalyser {
	p := a.p.New(audio, fftSize)
	return &AudioAnalyser{p: p}
}

// GetData TODO description.
func (a *AudioAnalyser) GetData() *AudioAnalyser {
	a.p.Call("getData")
	return a
}
