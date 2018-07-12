// Copyright 2016 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
