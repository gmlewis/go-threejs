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

// ColorKeyframeTrack represents a colorkeyframetrack.
type ColorKeyframeTrack struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (c *ColorKeyframeTrack) JSObject() *js.Object { return c.p }

// ColorKeyframeTrack returns a ColorKeyframeTrack JavaScript class.
func (t *Three) ColorKeyframeTrack() *ColorKeyframeTrack {
	p := t.ctx.Get("ColorKeyframeTrack")
	return ColorKeyframeTrackFromJSObject(p)
}

// ColorKeyframeTrackFromJSObject returns a wrapped ColorKeyframeTrack JavaScript class.
func ColorKeyframeTrackFromJSObject(p *js.Object) *ColorKeyframeTrack {
	return &ColorKeyframeTrack{p: p}
}

// NewColorKeyframeTrack returns a new ColorKeyframeTrack object.
func (t *Three) NewColorKeyframeTrack(name, times, values, interpolation float64) *ColorKeyframeTrack {
	p := t.ctx.Get("ColorKeyframeTrack").New(name, times, values, interpolation)
	return ColorKeyframeTrackFromJSObject(p)
}
