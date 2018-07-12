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

// AudioBuffer represents an audiobuffer.
type AudioBuffer struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (a *AudioBuffer) JSObject() *js.Object { return a.p }

// AudioBuffer returns an AudioBuffer JavaScript class.
func (t *Three) AudioBuffer() *AudioBuffer {
	p := t.ctx.Get("AudioBuffer")
	return AudioBufferFromJSObject(p)
}

// AudioBufferFromJSObject returns a wrapped AudioBuffer JavaScript class.
func AudioBufferFromJSObject(p *js.Object) *AudioBuffer {
	return &AudioBuffer{p: p}
}

// NewAudioBuffer returns a new AudioBuffer object.
func (t *Three) NewAudioBuffer(context float64) *AudioBuffer {
	p := t.ctx.Get("AudioBuffer").New(context)
	return AudioBufferFromJSObject(p)
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
