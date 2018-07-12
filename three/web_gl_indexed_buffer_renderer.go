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

// WebGLIndexedBufferRenderer represents a webglindexedbufferrenderer.
type WebGLIndexedBufferRenderer struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WebGLIndexedBufferRenderer) JSObject() *js.Object { return w.p }

// WebGLIndexedBufferRenderer returns a WebGLIndexedBufferRenderer JavaScript class.
func (t *Three) WebGLIndexedBufferRenderer() *WebGLIndexedBufferRenderer {
	p := t.ctx.Get("WebGLIndexedBufferRenderer")
	return WebGLIndexedBufferRendererFromJSObject(p)
}

// WebGLIndexedBufferRendererFromJSObject returns a wrapped WebGLIndexedBufferRenderer JavaScript class.
func WebGLIndexedBufferRendererFromJSObject(p *js.Object) *WebGLIndexedBufferRenderer {
	return &WebGLIndexedBufferRenderer{p: p}
}

// NewWebGLIndexedBufferRenderer returns a new WebGLIndexedBufferRenderer object.
func (t *Three) NewWebGLIndexedBufferRenderer(_gl, extensions, _infoRender float64) *WebGLIndexedBufferRenderer {
	p := t.ctx.Get("WebGLIndexedBufferRenderer").New(_gl, extensions, _infoRender)
	return WebGLIndexedBufferRendererFromJSObject(p)
}
