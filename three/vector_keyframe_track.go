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
func (v *VectorKeyframeTrack) JSObject() *js.Object { return v.p }

// VectorKeyframeTrack returns a VectorKeyframeTrack JavaScript class.
func (t *Three) VectorKeyframeTrack() *VectorKeyframeTrack {
	p := t.ctx.Get("VectorKeyframeTrack")
	return VectorKeyframeTrackFromJSObject(p)
}

// VectorKeyframeTrackFromJSObject returns a wrapped VectorKeyframeTrack JavaScript class.
func VectorKeyframeTrackFromJSObject(p *js.Object) *VectorKeyframeTrack {
	return &VectorKeyframeTrack{p: p}
}

// NewVectorKeyframeTrack returns a new VectorKeyframeTrack object.
func (t *Three) NewVectorKeyframeTrack(name, times, values, interpolation float64) *VectorKeyframeTrack {
	p := t.ctx.Get("VectorKeyframeTrack").New(name, times, values, interpolation)
	return VectorKeyframeTrackFromJSObject(p)
}
