// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// BufferGeometryLoader represents a buffergeometryloader.
type BufferGeometryLoader struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (b *BufferGeometryLoader) JSObject() *js.Object { return b.p }

// BufferGeometryLoader returns a BufferGeometryLoader JavaScript class.
func (t *Three) BufferGeometryLoader() *BufferGeometryLoader {
	p := t.ctx.Get("BufferGeometryLoader")
	return &BufferGeometryLoader{p: p}
}

// NewBufferGeometryLoader returns a new BufferGeometryLoader object.
func (t *Three) NewBufferGeometryLoader() *BufferGeometryLoader {
	p := t.ctx.Get("BufferGeometryLoader").New()
	return &BufferGeometryLoader{p: p}
}

// Load TODO description.
func (b *BufferGeometryLoader) Load(url string, onLoad LoadFunc, onProgress, onError interface{}) *BufferGeometryLoader {
	onLoadWrapper := onLoadWrapperFunc(onLoad)
	switch {
	case onProgress != nil && onError != nil:
		b.p.Call("load", url, onLoadWrapper, onProgress, onError)
	case onProgress != nil:
		b.p.Call("load", url, onLoadWrapper, onProgress)
	default:
		b.p.Call("load", url, onLoadWrapper)
	}
	return b
}

// Parse TODO description.
func (b *BufferGeometryLoader) Parse(json float64) *BufferGeometryLoader {
	b.p.Call("parse", json)
	return b
}
