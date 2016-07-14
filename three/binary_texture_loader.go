// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// BinaryTextureLoader represents a binarytextureloader.
type BinaryTextureLoader struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (b *BinaryTextureLoader) JSObject() *js.Object { return b.p }

// BinaryTextureLoader returns a BinaryTextureLoader JavaScript class.
func (t *Three) BinaryTextureLoader() *BinaryTextureLoader {
	p := t.ctx.Get("BinaryTextureLoader")
	return BinaryTextureLoaderFromJSObject(p)
}

// BinaryTextureLoaderFromJSObject returns a wrapped BinaryTextureLoader JavaScript class.
func BinaryTextureLoaderFromJSObject(p *js.Object) *BinaryTextureLoader {
	return &BinaryTextureLoader{p: p}
}

// NewBinaryTextureLoader returns a new BinaryTextureLoader object.
func (t *Three) NewBinaryTextureLoader() *BinaryTextureLoader {
	p := t.ctx.Get("BinaryTextureLoader").New()
	return BinaryTextureLoaderFromJSObject(p)
}

// Load TODO description.
func (b *BinaryTextureLoader) Load(url, onLoad, onProgress, onError float64) *BinaryTextureLoader {
	b.p.Call("load", url, onLoad, onProgress, onError)
	return b
}
