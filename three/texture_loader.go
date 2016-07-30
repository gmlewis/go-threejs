// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// TextureLoader is a class for loading a texture.
type TextureLoader struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *TextureLoader) JSObject() *js.Object { return t.p }

// TextureLoader returns a TextureLoader JavaScript class.
func (t *Three) TextureLoader() *TextureLoader {
	p := t.ctx.Get("TextureLoader")
	return &TextureLoader{p: p}
}

// NewTextureLoader creates a new TextureLoader object.
//
//     manager — The loadingManager for the loader to use. Default is THREE.DefaultLoadingManager.
func (t *Three) NewTextureLoader(manager float64) *TextureLoader {
	p := t.ctx.Get("TextureLoader").New(manager)
	return &TextureLoader{p: p}
}

// Load begins loading from url and passes the loaded texture to onLoad.
//
//     url — required
//     onLoad — Will be called when load completes. The argument will be the loaded texture.
//     onProgress — Will be called while load progresses. The argument will be the XmlHttpRequest instance, that contain .total and .loaded bytes.
//     onError — Will be called when load errors.
func (t *TextureLoader) Load(url, onLoad, onProgress, onError float64) *TextureLoader {
	t.p.Call("load", url, onLoad, onProgress, onError)
	return t
}

// SetCrossOrigin TODO description.
func (t *TextureLoader) SetCrossOrigin(value float64) *TextureLoader {
	t.p.Call("setCrossOrigin", value)
	return t
}

// SetPath TODO description.
func (t *TextureLoader) SetPath(value float64) *TextureLoader {
	t.p.Call("setPath", value)
	return t
}
