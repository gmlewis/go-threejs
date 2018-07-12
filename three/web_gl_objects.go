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

// WebGLObjects represents a webglobjects.
type WebGLObjects struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WebGLObjects) JSObject() *js.Object { return w.p }

// WebGLObjects returns a WebGLObjects JavaScript class.
func (t *Three) WebGLObjects() *WebGLObjects {
	p := t.ctx.Get("WebGLObjects")
	return WebGLObjectsFromJSObject(p)
}

// WebGLObjectsFromJSObject returns a wrapped WebGLObjects JavaScript class.
func WebGLObjectsFromJSObject(p *js.Object) *WebGLObjects {
	return &WebGLObjects{p: p}
}

// NewWebGLObjects returns a new WebGLObjects object.
func (t *Three) NewWebGLObjects(gl, properties, info float64) *WebGLObjects {
	p := t.ctx.Get("WebGLObjects").New(gl, properties, info)
	return WebGLObjectsFromJSObject(p)
}
