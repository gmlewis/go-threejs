// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Layers represents a layers.
type Layers struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (l *Layers) JSObject() *js.Object { return l.p }

// Layers returns a Layers JavaScript class.
func (t *Three) Layers() *Layers {
	p := t.ctx.Get("Layers")
	return LayersFromJSObject(p)
}

// LayersFromJSObject returns a wrapped Layers JavaScript class.
func LayersFromJSObject(p *js.Object) *Layers {
	return &Layers{p: p}
}

// NewLayers returns a new Layers object.
func (t *Three) NewLayers() *Layers {
	p := t.ctx.Get("Layers").New()
	return LayersFromJSObject(p)
}

// Set TODO description.
func (l *Layers) Set(channel float64) *Layers {
	l.p.Call("set", channel)
	return l
}

// Enable TODO description.
func (l *Layers) Enable(channel float64) *Layers {
	l.p.Call("enable", channel)
	return l
}

// Toggle TODO description.
func (l *Layers) Toggle(channel float64) *Layers {
	l.p.Call("toggle", channel)
	return l
}

// Disable TODO description.
func (l *Layers) Disable(channel float64) *Layers {
	l.p.Call("disable", channel)
	return l
}

// Test TODO description.
func (l *Layers) Test(layers float64) *Layers {
	l.p.Call("test", layers)
	return l
}
