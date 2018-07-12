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

// AudioListener represents an audiolistener.
type AudioListener struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (a *AudioListener) JSObject() *js.Object { return a.p }

// AudioListener returns an AudioListener JavaScript class.
func (t *Three) AudioListener() *AudioListener {
	p := t.ctx.Get("AudioListener")
	return AudioListenerFromJSObject(p)
}

// AudioListenerFromJSObject returns a wrapped AudioListener JavaScript class.
func AudioListenerFromJSObject(p *js.Object) *AudioListener {
	return &AudioListener{p: p}
}

// NewAudioListener returns a new AudioListener object.
func (t *Three) NewAudioListener() *AudioListener {
	p := t.ctx.Get("AudioListener").New()
	return AudioListenerFromJSObject(p)
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
