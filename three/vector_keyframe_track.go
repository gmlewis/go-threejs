// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// VectorKeyframeTrack represents a vectorkeyframetrack.
type VectorKeyframeTrack struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *VectorKeyframeTrack) JSObject() *js.Object { return t.p }

// VectorKeyframeTrack returns a VectorKeyframeTrack object.
func (t *Three) VectorKeyframeTrack() *VectorKeyframeTrack {
	p := t.ctx.Get("VectorKeyframeTrack")
	return &VectorKeyframeTrack{p: p}
}

// New returns a new VectorKeyframeTrack object.
func (t *VectorKeyframeTrack) New(name, times, values, interpolation float64) *VectorKeyframeTrack {
	p := t.p.New(name, times, values, interpolation)
	return &VectorKeyframeTrack{p: p}
}
