// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// JSONLoader represents a jsonloader.
type JSONLoader struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *JSONLoader) JSObject() *js.Object { return t.p }

// JSONLoader returns a JSONLoader object.
func (t *Three) JSONLoader() *JSONLoader {
	p := t.ctx.Get("JSONLoader")
	return &JSONLoader{p: p}
}

// New returns a new JSONLoader object.
func (t *JSONLoader) New(manager float64) *JSONLoader {
	p := t.p.New(manager)
	return &JSONLoader{p: p}
}

// StatusDomElement returns the statusDomElement-component of the JSONLoader.
func (j *JSONLoader) StatusDomElement() float64 {
	return j.p.Get("statusDomElement").Float()
}

// Load TODO description.
func (j *JSONLoader) Load(url, onLoad, onProgress, onError float64) *JSONLoader {
	j.p.Call("load", url, onLoad, onProgress, onError)
	return j
}

// SetTexturePath TODO description.
func (j *JSONLoader) SetTexturePath(value float64) *JSONLoader {
	j.p.Call("setTexturePath", value)
	return j
}

// Parse TODO description.
func (j *JSONLoader) Parse(json, texturePath float64) *JSONLoader {
	j.p.Call("parse", json, texturePath)
	return j
}
