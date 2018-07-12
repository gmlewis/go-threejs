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

// WebGLExtensions represents a webglextensions.
type WebGLExtensions struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WebGLExtensions) JSObject() *js.Object { return w.p }

// WebGLExtensions returns a WebGLExtensions JavaScript class.
func (t *Three) WebGLExtensions() *WebGLExtensions {
	p := t.ctx.Get("WebGLExtensions")
	return WebGLExtensionsFromJSObject(p)
}

// WebGLExtensionsFromJSObject returns a wrapped WebGLExtensions JavaScript class.
func WebGLExtensionsFromJSObject(p *js.Object) *WebGLExtensions {
	return &WebGLExtensions{p: p}
}

// NewWebGLExtensions returns a new WebGLExtensions object.
func (t *Three) NewWebGLExtensions(gl float64) *WebGLExtensions {
	p := t.ctx.Get("WebGLExtensions").New(gl)
	return WebGLExtensionsFromJSObject(p)
}
