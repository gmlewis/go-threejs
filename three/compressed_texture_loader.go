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
func (t *CompressedTextureLoader) JSObject() *js.Object { return t.p }

// CompressedTextureLoader returns a CompressedTextureLoader object.
func (t *Three) CompressedTextureLoader() *CompressedTextureLoader {
	p := t.ctx.Get("CompressedTextureLoader")
	return &CompressedTextureLoader{p: p}
}

// New returns a new CompressedTextureLoader object.
func (t *CompressedTextureLoader) New(manager float64) *CompressedTextureLoader {
	p := t.p.New(manager)
	return &CompressedTextureLoader{p: p}
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
