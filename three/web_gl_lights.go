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

// WebGLLights represents a webgllights.
type WebGLLights struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WebGLLights) JSObject() *js.Object { return w.p }

// WebGLLights returns a WebGLLights JavaScript class.
func (t *Three) WebGLLights() *WebGLLights {
	p := t.ctx.Get("WebGLLights")
	return WebGLLightsFromJSObject(p)
}

// WebGLLightsFromJSObject returns a wrapped WebGLLights JavaScript class.
func WebGLLightsFromJSObject(p *js.Object) *WebGLLights {
	return &WebGLLights{p: p}
}

// NewWebGLLights returns a new WebGLLights object.
func (t *Three) NewWebGLLights() *WebGLLights {
	p := t.ctx.Get("WebGLLights").New()
	return WebGLLightsFromJSObject(p)
}
