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

// NumberKeyframeTrack returns a NumberKeyframeTrack object.
func (t *Three) NumberKeyframeTrack() *NumberKeyframeTrack {
	p := t.ctx.Get("NumberKeyframeTrack")
	return &NumberKeyframeTrack{p: p}
}

// New returns a new NumberKeyframeTrack object.
func (n *NumberKeyframeTrack) New(name, times, values, interpolation float64) *NumberKeyframeTrack {
	p := n.p.New(name, times, values, interpolation)
	return &NumberKeyframeTrack{p: p}
}
