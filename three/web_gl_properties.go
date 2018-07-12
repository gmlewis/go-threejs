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

// WebGLProperties represents a webglproperties.
type WebGLProperties struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WebGLProperties) JSObject() *js.Object { return w.p }

// WebGLProperties returns a WebGLProperties JavaScript class.
func (t *Three) WebGLProperties() *WebGLProperties {
	p := t.ctx.Get("WebGLProperties")
	return WebGLPropertiesFromJSObject(p)
}

// WebGLPropertiesFromJSObject returns a wrapped WebGLProperties JavaScript class.
func WebGLPropertiesFromJSObject(p *js.Object) *WebGLProperties {
	return &WebGLProperties{p: p}
}

// NewWebGLProperties returns a new WebGLProperties object.
func (t *Three) NewWebGLProperties() *WebGLProperties {
	p := t.ctx.Get("WebGLProperties").New()
	return WebGLPropertiesFromJSObject(p)
}
