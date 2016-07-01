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
func (t *Three) NewBufferGeometryLoader(manager float64) *BufferGeometryLoader {
	p := t.ctx.Get("BufferGeometryLoader").New(manager)
	return &BufferGeometryLoader{p: p}
}

// Load TODO description.
func (b *BufferGeometryLoader) Load(url, onLoad, onProgress, onError float64) *BufferGeometryLoader {
	b.p.Call("load", url, onLoad, onProgress, onError)
	return b
}

// Parse TODO description.
func (b *BufferGeometryLoader) Parse(json float64) *BufferGeometryLoader {
	b.p.Call("parse", json)
	return b
}
