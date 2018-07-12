// Copyright 2016 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// VideoTexture represents a videotexture.
type VideoTexture struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (v *VideoTexture) JSObject() *js.Object { return v.p }

// VideoTexture returns a VideoTexture JavaScript class.
func (t *Three) VideoTexture() *VideoTexture {
	p := t.ctx.Get("VideoTexture")
	return VideoTextureFromJSObject(p)
}

// VideoTextureFromJSObject returns a wrapped VideoTexture JavaScript class.
func VideoTextureFromJSObject(p *js.Object) *VideoTexture {
	return &VideoTexture{p: p}
}

// NewVideoTexture returns a new VideoTexture object.
func (t *Three) NewVideoTexture(video, mapping, wrapS, wrapT, magFilter, minFilter, format, typ, anisotropy float64) *VideoTexture {
	p := t.ctx.Get("VideoTexture").New(video, mapping, wrapS, wrapT, magFilter, minFilter, format, typ, anisotropy)
	return VideoTextureFromJSObject(p)
}
