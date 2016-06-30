// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// ColorKeyframeTrack represents a colorkeyframetrack.
type ColorKeyframeTrack struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *ColorKeyframeTrack) JSObject() *js.Object { return t.p }

// ColorKeyframeTrack returns a ColorKeyframeTrack object.
func (t *Three) ColorKeyframeTrack() *ColorKeyframeTrack {
	p := t.ctx.Get("ColorKeyframeTrack")
	return &ColorKeyframeTrack{p: p}
}

// New returns a new ColorKeyframeTrack object.
func (t *ColorKeyframeTrack) New(name, times, values, interpolation float64) *ColorKeyframeTrack {
	p := t.p.New(name, times, values, interpolation)
	return &ColorKeyframeTrack{p: p}
}
