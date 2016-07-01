// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// ImageLoader represents an imageloader.
type ImageLoader struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (i *ImageLoader) JSObject() *js.Object { return i.p }

// ImageLoader returns an ImageLoader JavaScript class.
func (t *Three) ImageLoader() *ImageLoader {
	p := t.ctx.Get("ImageLoader")
	return &ImageLoader{p: p}
}

// NewImageLoader returns a new ImageLoader object.
func (t *Three) NewImageLoader(manager float64) *ImageLoader {
	p := t.ctx.Get("ImageLoader").New(manager)
	return &ImageLoader{p: p}
}

// Load TODO description.
func (i *ImageLoader) Load(url, onLoad, onProgress, onError float64) *ImageLoader {
	i.p.Call("load", url, onLoad, onProgress, onError)
	return i
}

// SetCrossOrigin TODO description.
func (i *ImageLoader) SetCrossOrigin(value float64) *ImageLoader {
	i.p.Call("setCrossOrigin", value)
	return i
}

// SetPath TODO description.
func (i *ImageLoader) SetPath(value float64) *ImageLoader {
	i.p.Call("setPath", value)
	return i
}
