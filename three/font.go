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

// Font represents a font.
type Font struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (f *Font) JSObject() *js.Object { return f.p }

// Font returns a Font JavaScript class.
func (t *Three) Font() *Font {
	p := t.ctx.Get("Font")
	return FontFromJSObject(p)
}

// FontFromJSObject returns a wrapped Font JavaScript class.
func FontFromJSObject(p *js.Object) *Font {
	return &Font{p: p}
}

// NewFont returns a new Font object.
func (t *Three) NewFont(data float64) *Font {
	p := t.ctx.Get("Font").New(data)
	return FontFromJSObject(p)
}

// GenerateShapes TODO description.
func (f *Font) GenerateShapes(text, size, divisions float64) *Font {
	f.p.Call("generateShapes", text, size, divisions)
	return f
}
