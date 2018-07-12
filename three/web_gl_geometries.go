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

// WebGLGeometries represents a webglgeometries.
type WebGLGeometries struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WebGLGeometries) JSObject() *js.Object { return w.p }

// WebGLGeometries returns a WebGLGeometries JavaScript class.
func (t *Three) WebGLGeometries() *WebGLGeometries {
	p := t.ctx.Get("WebGLGeometries")
	return WebGLGeometriesFromJSObject(p)
}

// WebGLGeometriesFromJSObject returns a wrapped WebGLGeometries JavaScript class.
func WebGLGeometriesFromJSObject(p *js.Object) *WebGLGeometries {
	return &WebGLGeometries{p: p}
}

// NewWebGLGeometries returns a new WebGLGeometries object.
func (t *Three) NewWebGLGeometries(gl, properties, info float64) *WebGLGeometries {
	p := t.ctx.Get("WebGLGeometries").New(gl, properties, info)
	return WebGLGeometriesFromJSObject(p)
}
