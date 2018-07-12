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

// StringKeyframeTrack represents a stringkeyframetrack.
type StringKeyframeTrack struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (s *StringKeyframeTrack) JSObject() *js.Object { return s.p }

// StringKeyframeTrack returns a StringKeyframeTrack JavaScript class.
func (t *Three) StringKeyframeTrack() *StringKeyframeTrack {
	p := t.ctx.Get("StringKeyframeTrack")
	return StringKeyframeTrackFromJSObject(p)
}

// StringKeyframeTrackFromJSObject returns a wrapped StringKeyframeTrack JavaScript class.
func StringKeyframeTrackFromJSObject(p *js.Object) *StringKeyframeTrack {
	return &StringKeyframeTrack{p: p}
}

// NewStringKeyframeTrack returns a new StringKeyframeTrack object.
func (t *Three) NewStringKeyframeTrack(name, times, values, interpolation float64) *StringKeyframeTrack {
	p := t.ctx.Get("StringKeyframeTrack").New(name, times, values, interpolation)
	return StringKeyframeTrackFromJSObject(p)
}
