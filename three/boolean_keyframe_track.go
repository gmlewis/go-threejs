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
func (t *BooleanKeyframeTrack) JSObject() *js.Object { return t.p }

// BooleanKeyframeTrack returns a BooleanKeyframeTrack object.
func (t *Three) BooleanKeyframeTrack() *BooleanKeyframeTrack {
	p := t.ctx.Get("BooleanKeyframeTrack")
	return &BooleanKeyframeTrack{p: p}
}

// New returns a new BooleanKeyframeTrack object.
func (t *BooleanKeyframeTrack) New(name, times, values float64) *BooleanKeyframeTrack {
	p := t.p.New(name, times, values)
	return &BooleanKeyframeTrack{p: p}
}
