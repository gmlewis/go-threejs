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

// QuaternionKeyframeTrack represents a quaternionkeyframetrack.
type QuaternionKeyframeTrack struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (q *QuaternionKeyframeTrack) JSObject() *js.Object { return q.p }

// QuaternionKeyframeTrack returns a QuaternionKeyframeTrack JavaScript class.
func (t *Three) QuaternionKeyframeTrack() *QuaternionKeyframeTrack {
	p := t.ctx.Get("QuaternionKeyframeTrack")
	return QuaternionKeyframeTrackFromJSObject(p)
}

// QuaternionKeyframeTrackFromJSObject returns a wrapped QuaternionKeyframeTrack JavaScript class.
func QuaternionKeyframeTrackFromJSObject(p *js.Object) *QuaternionKeyframeTrack {
	return &QuaternionKeyframeTrack{p: p}
}

// NewQuaternionKeyframeTrack returns a new QuaternionKeyframeTrack object.
func (t *Three) NewQuaternionKeyframeTrack(name, times, values, interpolation float64) *QuaternionKeyframeTrack {
	p := t.ctx.Get("QuaternionKeyframeTrack").New(name, times, values, interpolation)
	return QuaternionKeyframeTrackFromJSObject(p)
}

// InterpolantFactoryMethodLinear TODO description.
func (q *QuaternionKeyframeTrack) InterpolantFactoryMethodLinear(result float64) *QuaternionKeyframeTrack {
	q.p.Call("InterpolantFactoryMethodLinear", result)
	return q
}
