// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// BufferGeometryLoader is a loader for loading a BufferGeometry.
//
// http://threejs.org/docs/index.html#Reference/Loaders/BufferGeometryLoader
type BufferGeometryLoader struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (b *BufferGeometryLoader) JSObject() *js.Object { return b.p }

// BufferGeometryLoader returns a BufferGeometryLoader JavaScript class.
func (t *Three) BufferGeometryLoader() *BufferGeometryLoader {
	p := t.ctx.Get("BufferGeometryLoader")
	return BufferGeometryLoaderFromJSObject(p)
}

// BufferGeometryLoaderFromJSObject returns a wrapped BufferGeometryLoader JavaScript class.
func BufferGeometryLoaderFromJSObject(p *js.Object) *BufferGeometryLoader {
	return &BufferGeometryLoader{p: p}
}

// NewBufferGeometryLoader returns a new BufferGeometryLoader object.
func (t *Three) NewBufferGeometryLoader() *BufferGeometryLoader {
	p := t.ctx.Get("BufferGeometryLoader").New()
	return BufferGeometryLoaderFromJSObject(p)
}

// BufferGeometryLoadFunc is a callback function called by Load.
type BufferGeometryLoadFunc func(geometry *BufferGeometry, materials []*Material)

// onBufferGeometryLoadWrapperFunc wraps the geometry and materials to a typed LoadFunc.
func onBufferGeometryLoadWrapperFunc(onLoad BufferGeometryLoadFunc) func(geometry, materials *js.Object) {
	return func(geom, materialArray *js.Object) {
		var materials []*Material
		if materialArray != nil && materialArray != js.Undefined {
			for i := 0; i < materialArray.Length(); i++ {
				materials = append(materials, MaterialFromJSObject(materialArray.Index(i)))
			}
		}
		onLoad(BufferGeometryFromJSObject(geom), materials)
	}
}

// Load TODO description.
func (b *BufferGeometryLoader) Load(url string, onLoad BufferGeometryLoadFunc, onProgress, onError interface{}) *BufferGeometryLoader {
	onLoadWrapper := onBufferGeometryLoadWrapperFunc(onLoad)
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
