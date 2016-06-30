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

// PositionalAudio returns a PositionalAudio object.
func (t *Three) PositionalAudio() *PositionalAudio {
	p := t.ctx.Get("PositionalAudio")
	return &PositionalAudio{p: p}
}

// New returns a new PositionalAudio object.
func (p *PositionalAudio) New(listener float64) *PositionalAudio {
	t := p.p.New(listener)
	return &PositionalAudio{p: t}
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
