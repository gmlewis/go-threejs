// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

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
