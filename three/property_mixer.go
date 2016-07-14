// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// PropertyMixer represents a propertymixer.
type PropertyMixer struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (p *PropertyMixer) JSObject() *js.Object { return p.p }

// PropertyMixer returns a PropertyMixer JavaScript class.
func (t *Three) PropertyMixer() *PropertyMixer {
	p := t.ctx.Get("PropertyMixer")
	return PropertyMixerFromJSObject(p)
}

// PropertyMixerFromJSObject returns a wrapped PropertyMixer JavaScript class.
func PropertyMixerFromJSObject(p *js.Object) *PropertyMixer {
	return &PropertyMixer{p: p}
}

// NewPropertyMixer returns a new PropertyMixer object.
func (t *Three) NewPropertyMixer(binding, typName, valueSize float64) *PropertyMixer {
	p := t.ctx.Get("PropertyMixer").New(binding, typName, valueSize)
	return PropertyMixerFromJSObject(p)
}

// Accumulate TODO description.
func (p *PropertyMixer) Accumulate(accuIndex, weight float64) *PropertyMixer {
	p.p.Call("accumulate", accuIndex, weight)
	return p
}

// Apply TODO description.
func (p *PropertyMixer) Apply(accuIndex float64) *PropertyMixer {
	p.p.Call("apply", accuIndex)
	return p
}

// SaveOriginalState TODO description.
func (p *PropertyMixer) SaveOriginalState() *PropertyMixer {
	p.p.Call("saveOriginalState")
	return p
}

// RestoreOriginalState TODO description.
func (p *PropertyMixer) RestoreOriginalState() *PropertyMixer {
	p.p.Call("restoreOriginalState")
	return p
}

// _select TODO description.
func (p *PropertyMixer) _select(buffer, dstOffset, srcOffset, t, stride float64) *PropertyMixer {
	p.p.Call("_select", buffer, dstOffset, srcOffset, t, stride)
	return p
}

// _slerp TODO description.
func (p *PropertyMixer) _slerp(buffer, dstOffset, srcOffset, t, stride float64) *PropertyMixer {
	p.p.Call("_slerp", buffer, dstOffset, srcOffset, t, stride)
	return p
}

// _lerp TODO description.
func (p *PropertyMixer) _lerp(buffer, dstOffset, srcOffset, t, stride float64) *PropertyMixer {
	p.p.Call("_lerp", buffer, dstOffset, srcOffset, t, stride)
	return p
}
