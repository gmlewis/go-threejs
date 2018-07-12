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

// WebGLShadowMap represents a webglshadowmap.
type WebGLShadowMap struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WebGLShadowMap) JSObject() *js.Object { return w.p }

// WebGLShadowMap returns a WebGLShadowMap JavaScript class.
func (t *Three) WebGLShadowMap() *WebGLShadowMap {
	p := t.ctx.Get("WebGLShadowMap")
	return WebGLShadowMapFromJSObject(p)
}

// WebGLShadowMapFromJSObject returns a wrapped WebGLShadowMap JavaScript class.
func WebGLShadowMapFromJSObject(p *js.Object) *WebGLShadowMap {
	return &WebGLShadowMap{p: p}
}

// NewWebGLShadowMap returns a new WebGLShadowMap object.
func (t *Three) NewWebGLShadowMap(_renderer, _lights, _objects float64) *WebGLShadowMap {
	p := t.ctx.Get("WebGLShadowMap").New(_renderer, _lights, _objects)
	return WebGLShadowMapFromJSObject(p)
}
