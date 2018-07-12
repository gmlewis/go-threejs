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

// KeyframeTrack represents a keyframetrack.
type KeyframeTrack struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (k *KeyframeTrack) JSObject() *js.Object { return k.p }

// KeyframeTrack returns a KeyframeTrack JavaScript class.
func (t *Three) KeyframeTrack() *KeyframeTrack {
	p := t.ctx.Get("KeyframeTrack")
	return KeyframeTrackFromJSObject(p)
}

// KeyframeTrackFromJSObject returns a wrapped KeyframeTrack JavaScript class.
func KeyframeTrackFromJSObject(p *js.Object) *KeyframeTrack {
	return &KeyframeTrack{p: p}
}

// NewKeyframeTrack returns a new KeyframeTrack object.
func (t *Three) NewKeyframeTrack(name, times, values, interpolation float64) *KeyframeTrack {
	p := t.ctx.Get("KeyframeTrack").New(name, times, values, interpolation)
	return KeyframeTrackFromJSObject(p)
}

// InterpolantFactoryMethodDiscrete TODO description.
func (k *KeyframeTrack) InterpolantFactoryMethodDiscrete(result float64) *KeyframeTrack {
	k.p.Call("InterpolantFactoryMethodDiscrete", result)
	return k
}

// InterpolantFactoryMethodLinear TODO description.
func (k *KeyframeTrack) InterpolantFactoryMethodLinear(result float64) *KeyframeTrack {
	k.p.Call("InterpolantFactoryMethodLinear", result)
	return k
}

// InterpolantFactoryMethodSmooth TODO description.
func (k *KeyframeTrack) InterpolantFactoryMethodSmooth(result float64) *KeyframeTrack {
	k.p.Call("InterpolantFactoryMethodSmooth", result)
	return k
}

// SetInterpolation TODO description.
func (k *KeyframeTrack) SetInterpolation(interpolation float64) *KeyframeTrack {
	k.p.Call("setInterpolation", interpolation)
	return k
}

// GetInterpolation TODO description.
func (k *KeyframeTrack) GetInterpolation() *KeyframeTrack {
	k.p.Call("getInterpolation")
	return k
}

// GetValueSize TODO description.
func (k *KeyframeTrack) GetValueSize() *KeyframeTrack {
	k.p.Call("getValueSize")
	return k
}

// Shift TODO description.
func (k *KeyframeTrack) Shift(timeOffset float64) *KeyframeTrack {
	k.p.Call("shift", timeOffset)
	return k
}

// Scale TODO description.
func (k *KeyframeTrack) Scale(timeScale float64) *KeyframeTrack {
	k.p.Call("scale", timeScale)
	return k
}

// Trim TODO description.
func (k *KeyframeTrack) Trim(startTime, endTime float64) *KeyframeTrack {
	k.p.Call("trim", startTime, endTime)
	return k
}

// Validate TODO description.
func (k *KeyframeTrack) Validate() *KeyframeTrack {
	k.p.Call("validate")
	return k
}

// Optimize TODO description.
func (k *KeyframeTrack) Optimize() *KeyframeTrack {
	k.p.Call("optimize")
	return k
}

// Parse TODO description.
func (k *KeyframeTrack) Parse(json float64) *KeyframeTrack {
	k.p.Call("parse", json)
	return k
}

// ToJSON TODO description.
func (k *KeyframeTrack) ToJSON(track float64) *KeyframeTrack {
	k.p.Call("toJSON", track)
	return k
}

// _getTrackTypeForValueTypeName TODO description.
func (k *KeyframeTrack) _getTrackTypeForValueTypeName(typName float64) *KeyframeTrack {
	k.p.Call("_getTrackTypeForValueTypeName", typName)
	return k
}
