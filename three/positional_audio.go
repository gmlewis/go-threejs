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

// PositionalAudio represents a positionalaudio.
type PositionalAudio struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (p *PositionalAudio) JSObject() *js.Object { return p.p }

// PositionalAudio returns a PositionalAudio JavaScript class.
func (t *Three) PositionalAudio() *PositionalAudio {
	p := t.ctx.Get("PositionalAudio")
	return PositionalAudioFromJSObject(p)
}

// PositionalAudioFromJSObject returns a wrapped PositionalAudio JavaScript class.
func PositionalAudioFromJSObject(p *js.Object) *PositionalAudio {
	return &PositionalAudio{p: p}
}

// NewPositionalAudio returns a new PositionalAudio object.
func (t *Three) NewPositionalAudio(listener float64) *PositionalAudio {
	p := t.ctx.Get("PositionalAudio").New(listener)
	return PositionalAudioFromJSObject(p)
}

// SetRefDistance TODO description.
func (p *PositionalAudio) SetRefDistance(value float64) *PositionalAudio {
	p.p.Call("setRefDistance", value)
	return p
}

// SetRolloffFactor TODO description.
func (p *PositionalAudio) SetRolloffFactor(value float64) *PositionalAudio {
	p.p.Call("setRolloffFactor", value)
	return p
}

// SetDistanceModel TODO description.
func (p *PositionalAudio) SetDistanceModel(value float64) *PositionalAudio {
	p.p.Call("setDistanceModel", value)
	return p
}

// SetMaxDistance TODO description.
func (p *PositionalAudio) SetMaxDistance(value float64) *PositionalAudio {
	p.p.Call("setMaxDistance", value)
	return p
}
