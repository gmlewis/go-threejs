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

// TextureLoader returns a TextureLoader object.
func (t *Three) TextureLoader() *TextureLoader {
	p := t.ctx.Get("TextureLoader")
	return &TextureLoader{p: p}
}

// New returns a new TextureLoader object.
func (t *TextureLoader) New(manager float64) *TextureLoader {
	p := t.p.New(manager)
	return &TextureLoader{p: p}
}

// Load TODO description.
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
