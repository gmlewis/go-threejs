// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// VideoTexture represents a videotexture.
type VideoTexture struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *VideoTexture) JSObject() *js.Object { return t.p }

// VideoTexture returns a VideoTexture object.
func (t *Three) VideoTexture() *VideoTexture {
	p := t.ctx.Get("VideoTexture")
	return &VideoTexture{p: p}
}

// New returns a new VideoTexture object.
func (t *VideoTexture) New(video, mapping, wrapS, wrapT, magFilter, minFilter, format, typ, anisotropy float64) *VideoTexture {
	p := t.p.New(video, mapping, wrapS, wrapT, magFilter, minFilter, format, typ, anisotropy)
	return &VideoTexture{p: p}
}
