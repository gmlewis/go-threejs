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

// WebGLCapabilities represents a webglcapabilities.
type WebGLCapabilities struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WebGLCapabilities) JSObject() *js.Object { return w.p }

// WebGLCapabilities returns a WebGLCapabilities JavaScript class.
func (t *Three) WebGLCapabilities() *WebGLCapabilities {
	p := t.ctx.Get("WebGLCapabilities")
	return WebGLCapabilitiesFromJSObject(p)
}

// WebGLCapabilitiesFromJSObject returns a wrapped WebGLCapabilities JavaScript class.
func WebGLCapabilitiesFromJSObject(p *js.Object) *WebGLCapabilities {
	return &WebGLCapabilities{p: p}
}

// NewWebGLCapabilities returns a new WebGLCapabilities object.
func (t *Three) NewWebGLCapabilities(gl, extensions, parameters float64) *WebGLCapabilities {
	p := t.ctx.Get("WebGLCapabilities").New(gl, extensions, parameters)
	return WebGLCapabilitiesFromJSObject(p)
}
