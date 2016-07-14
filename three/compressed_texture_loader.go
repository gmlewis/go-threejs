// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CompressedTextureLoader represents a compressedtextureloader.
type CompressedTextureLoader struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (c *CompressedTextureLoader) JSObject() *js.Object { return c.p }

// CompressedTextureLoader returns a CompressedTextureLoader JavaScript class.
func (t *Three) CompressedTextureLoader() *CompressedTextureLoader {
	p := t.ctx.Get("CompressedTextureLoader")
	return CompressedTextureLoaderFromJSObject(p)
}

// CompressedTextureLoaderFromJSObject returns a wrapped CompressedTextureLoader JavaScript class.
func CompressedTextureLoaderFromJSObject(p *js.Object) *CompressedTextureLoader {
	return &CompressedTextureLoader{p: p}
}

// NewCompressedTextureLoader returns a new CompressedTextureLoader object.
func (t *Three) NewCompressedTextureLoader(manager float64) *CompressedTextureLoader {
	p := t.ctx.Get("CompressedTextureLoader").New(manager)
	return CompressedTextureLoaderFromJSObject(p)
}

// Load TODO description.
func (c *CompressedTextureLoader) Load(url, onLoad, onProgress, onError float64) *CompressedTextureLoader {
	c.p.Call("load", url, onLoad, onProgress, onError)
	return c
}

// SetPath TODO description.
func (c *CompressedTextureLoader) SetPath(value float64) *CompressedTextureLoader {
	c.p.Call("setPath", value)
	return c
}
