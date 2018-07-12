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

// BooleanKeyframeTrack represents a booleankeyframetrack.
type BooleanKeyframeTrack struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (b *BooleanKeyframeTrack) JSObject() *js.Object { return b.p }

// BooleanKeyframeTrack returns a BooleanKeyframeTrack JavaScript class.
func (t *Three) BooleanKeyframeTrack() *BooleanKeyframeTrack {
	p := t.ctx.Get("BooleanKeyframeTrack")
	return BooleanKeyframeTrackFromJSObject(p)
}

// BooleanKeyframeTrackFromJSObject returns a wrapped BooleanKeyframeTrack JavaScript class.
func BooleanKeyframeTrackFromJSObject(p *js.Object) *BooleanKeyframeTrack {
	return &BooleanKeyframeTrack{p: p}
}

// NewBooleanKeyframeTrack returns a new BooleanKeyframeTrack object.
func (t *Three) NewBooleanKeyframeTrack(name, times, values float64) *BooleanKeyframeTrack {
	p := t.ctx.Get("BooleanKeyframeTrack").New(name, times, values)
	return BooleanKeyframeTrackFromJSObject(p)
}
