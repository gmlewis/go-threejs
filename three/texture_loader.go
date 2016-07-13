// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// TextureLoader represents a textureloader.
type TextureLoader struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *TextureLoader) JSObject() *js.Object { return t.p }

// TextureLoader returns a TextureLoader JavaScript class.
func (t *Three) TextureLoader() *TextureLoader {
	p := t.ctx.Get("TextureLoader")
	return &TextureLoader{p: p}
}

// NewTextureLoader returns a new TextureLoader object.
func (t *Three) NewTextureLoader() *TextureLoader {
	p := t.ctx.Get("TextureLoader").New()
	return &TextureLoader{p: p}
}

// TextureLoadFunc is a callback function called by Load.
type TextureLoadFunc func(geometry *Geometry, materials []*Material)

// onTextureLoadWrapperFunc wraps the geometry and materials to a typed LoadFunc.
func onTextureLoadWrapperFunc(onLoad TextureLoadFunc) func(geometry, materials *js.Object) {
	return func(geom, materialArray *js.Object) {
		var materials []*Material
		if materialArray != nil && materialArray != js.Undefined {
			for i := 0; i < materialArray.Length(); i++ {
				materials = append(materials, MaterialFromJSObject(materialArray.Index(i)))
			}
		}
		onLoad(GeometryFromJSObject(geom), materials)
	}
}

// Load TODO description.
func (t *TextureLoader) Load(url string, onLoad TextureLoadFunc, onProgress, onError interface{}) *Texture {
	var onLoadWrapper func(geometry, materials *js.Object)
	if onLoad != nil {
		onLoadWrapper = onTextureLoadWrapperFunc(onLoad)
	}
	var p *js.Object
	switch {
	case onLoad != nil && onProgress != nil && onError != nil:
		p = t.p.Call("load", url, onLoadWrapper, onProgress, onError)
	case onLoad != nil && onProgress != nil:
		p = t.p.Call("load", url, onLoadWrapper, onProgress)
	case onLoad != nil:
		p = t.p.Call("load", url, onLoadWrapper)
	default:
		p = t.p.Call("load", url)
	}
	return TextureFromJSObject(p)
}

// SetCrossOrigin TODO description.
func (t *TextureLoader) SetCrossOrigin(value string) *TextureLoader {
	t.p.Call("setCrossOrigin", value)
	return t
}

// SetPath TODO description.
func (t *TextureLoader) SetPath(value float64) *TextureLoader {
	t.p.Call("setPath", value)
	return t
}
