// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// StereoCamera is a stereo projection camera with left and right eyes.
type StereoCamera struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *StereoCamera) JSObject() *js.Object { return t.p }

// StereoCamera returns a StereoCamera object.
func (t *Three) StereoCamera() *StereoCamera {
	p := t.ctx.Get("StereoCamera")
	return &StereoCamera{p: p}
}

// New returns a new StereoCamera object.
func (t *StereoCamera) New() *StereoCamera {
	p := t.p.New()
	return &StereoCamera{p: p}
}
