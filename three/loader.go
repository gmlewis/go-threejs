// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Loader represents a loader.
type Loader struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (l *Loader) JSObject() *js.Object { return l.p }

// Loader returns a Loader JavaScript class.
func (t *Three) Loader() *Loader {
	p := t.ctx.Get("Loader")
	return LoaderFromJSObject(p)
}

// LoaderFromJSObject returns a wrapped Loader JavaScript class.
func LoaderFromJSObject(p *js.Object) *Loader {
	return &Loader{p: p}
}

// NewLoader returns a new Loader object.
func (t *Three) NewLoader() *Loader {
	p := t.ctx.Get("Loader").New()
	return LoaderFromJSObject(p)
}

// ExtractURLBase TODO description.
func (l *Loader) ExtractURLBase(url float64) *Loader {
	l.p.Call("extractUrlBase", url)
	return l
}

// InitMaterials TODO description.
func (l *Loader) InitMaterials(materials, texturePath, crossOrigin float64) *Loader {
	l.p.Call("initMaterials", materials, texturePath, crossOrigin)
	return l
}

// Add TODO description.
func (l *Loader) Add(regex, loader float64) *Loader {
	l.p.Call("add", regex, loader)
	return l
}

// Get TODO description.
func (l *Loader) Get(file float64) *Loader {
	l.p.Call("get", file)
	return l
}
