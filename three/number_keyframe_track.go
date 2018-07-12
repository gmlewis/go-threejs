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

// NumberKeyframeTrack represents a numberkeyframetrack.
type NumberKeyframeTrack struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (n *NumberKeyframeTrack) JSObject() *js.Object { return n.p }

// NumberKeyframeTrack returns a NumberKeyframeTrack JavaScript class.
func (t *Three) NumberKeyframeTrack() *NumberKeyframeTrack {
	p := t.ctx.Get("NumberKeyframeTrack")
	return NumberKeyframeTrackFromJSObject(p)
}

// NumberKeyframeTrackFromJSObject returns a wrapped NumberKeyframeTrack JavaScript class.
func NumberKeyframeTrackFromJSObject(p *js.Object) *NumberKeyframeTrack {
	return &NumberKeyframeTrack{p: p}
}

// NewNumberKeyframeTrack returns a new NumberKeyframeTrack object.
func (t *Three) NewNumberKeyframeTrack(name, times, values, interpolation float64) *NumberKeyframeTrack {
	p := t.ctx.Get("NumberKeyframeTrack").New(name, times, values, interpolation)
	return NumberKeyframeTrackFromJSObject(p)
}
