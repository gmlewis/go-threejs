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

// FontLoader represents a fontloader.
type FontLoader struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (f *FontLoader) JSObject() *js.Object { return f.p }

// FontLoader returns a FontLoader JavaScript class.
func (t *Three) FontLoader() *FontLoader {
	p := t.ctx.Get("FontLoader")
	return FontLoaderFromJSObject(p)
}

// FontLoaderFromJSObject returns a wrapped FontLoader JavaScript class.
func FontLoaderFromJSObject(p *js.Object) *FontLoader {
	return &FontLoader{p: p}
}

// NewFontLoader returns a new FontLoader object.
func (t *Three) NewFontLoader(manager float64) *FontLoader {
	p := t.ctx.Get("FontLoader").New(manager)
	return FontLoaderFromJSObject(p)
}

// Load TODO description.
func (f *FontLoader) Load(url, onLoad, onProgress, onError float64) *FontLoader {
	f.p.Call("load", url, onLoad, onProgress, onError)
	return f
}
