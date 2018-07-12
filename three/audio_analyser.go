// Copyright 2016 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// AudioAnalyser represents an audioanalyser.
type AudioAnalyser struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (a *AudioAnalyser) JSObject() *js.Object { return a.p }

// AudioAnalyser returns an AudioAnalyser JavaScript class.
func (t *Three) AudioAnalyser() *AudioAnalyser {
	p := t.ctx.Get("AudioAnalyser")
	return AudioAnalyserFromJSObject(p)
}

// AudioAnalyserFromJSObject returns a wrapped AudioAnalyser JavaScript class.
func AudioAnalyserFromJSObject(p *js.Object) *AudioAnalyser {
	return &AudioAnalyser{p: p}
}

// NewAudioAnalyser returns a new AudioAnalyser object.
func (t *Three) NewAudioAnalyser(audio, fftSize float64) *AudioAnalyser {
	p := t.ctx.Get("AudioAnalyser").New(audio, fftSize)
	return AudioAnalyserFromJSObject(p)
}

// GetData TODO description.
func (a *AudioAnalyser) GetData() *AudioAnalyser {
	a.p.Call("getData")
	return a
}
