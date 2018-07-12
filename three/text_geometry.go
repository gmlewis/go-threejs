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

// TextGeometry represents a 3D object of text as a single object.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/TextGeometry
type TextGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *TextGeometry) JSObject() *js.Object { return t.p }

// TextGeometry returns a TextGeometry JavaScript class.
func (t *Three) TextGeometry() *TextGeometry {
	p := t.ctx.Get("TextGeometry")
	return TextGeometryFromJSObject(p)
}

// TextGeometryFromJSObject returns a wrapped TextGeometry JavaScript class.
func TextGeometryFromJSObject(p *js.Object) *TextGeometry {
	return &TextGeometry{p: p}
}

// NewTextGeometry returns a new TextGeometry object.
//
//     text — The text to be shown.
//     parameters — Object that can contains the following parameters.  TODO.
func (t *Three) NewTextGeometry(text string, parameters map[string]interface{}) *TextGeometry {
	p := t.ctx.Get("TextGeometry").New(text, parameters)
	return TextGeometryFromJSObject(p)
}
