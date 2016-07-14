// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

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
