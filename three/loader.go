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
