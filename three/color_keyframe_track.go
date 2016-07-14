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
