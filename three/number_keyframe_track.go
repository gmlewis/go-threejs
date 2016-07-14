// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

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
