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

// StringKeyframeTrack returns a StringKeyframeTrack object.
func (t *Three) StringKeyframeTrack() *StringKeyframeTrack {
	p := t.ctx.Get("StringKeyframeTrack")
	return &StringKeyframeTrack{p: p}
}

// New returns a new StringKeyframeTrack object.
func (s *StringKeyframeTrack) New(name, times, values, interpolation float64) *StringKeyframeTrack {
	p := s.p.New(name, times, values, interpolation)
	return &StringKeyframeTrack{p: p}
}
