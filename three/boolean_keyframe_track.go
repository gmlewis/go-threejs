// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

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
