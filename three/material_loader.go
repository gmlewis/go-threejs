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

// MaterialLoader represents a materialloader.
type MaterialLoader struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (m *MaterialLoader) JSObject() *js.Object { return m.p }

// MaterialLoader returns a MaterialLoader JavaScript class.
func (t *Three) MaterialLoader() *MaterialLoader {
	p := t.ctx.Get("MaterialLoader")
	return MaterialLoaderFromJSObject(p)
}

// MaterialLoaderFromJSObject returns a wrapped MaterialLoader JavaScript class.
func MaterialLoaderFromJSObject(p *js.Object) *MaterialLoader {
	return &MaterialLoader{p: p}
}

// NewMaterialLoader returns a new MaterialLoader object.
func (t *Three) NewMaterialLoader(manager float64) *MaterialLoader {
	p := t.ctx.Get("MaterialLoader").New(manager)
	return MaterialLoaderFromJSObject(p)
}

// Load TODO description.
func (m *MaterialLoader) Load(url, onLoad, onProgress, onError float64) *MaterialLoader {
	m.p.Call("load", url, onLoad, onProgress, onError)
	return m
}

// SetTextures TODO description.
func (m *MaterialLoader) SetTextures(value float64) *MaterialLoader {
	m.p.Call("setTextures", value)
	return m
}

// GetTexture TODO description.
func (m *MaterialLoader) GetTexture(name float64) *MaterialLoader {
	m.p.Call("getTexture", name)
	return m
}

// Parse TODO description.
func (m *MaterialLoader) Parse(json float64) *MaterialLoader {
	m.p.Call("parse", json)
	return m
}
