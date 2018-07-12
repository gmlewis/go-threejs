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

// WebGLBufferRenderer represents a webglbufferrenderer.
type WebGLBufferRenderer struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WebGLBufferRenderer) JSObject() *js.Object { return w.p }

// WebGLBufferRenderer returns a WebGLBufferRenderer JavaScript class.
func (t *Three) WebGLBufferRenderer() *WebGLBufferRenderer {
	p := t.ctx.Get("WebGLBufferRenderer")
	return WebGLBufferRendererFromJSObject(p)
}

// WebGLBufferRendererFromJSObject returns a wrapped WebGLBufferRenderer JavaScript class.
func WebGLBufferRendererFromJSObject(p *js.Object) *WebGLBufferRenderer {
	return &WebGLBufferRenderer{p: p}
}

// NewWebGLBufferRenderer returns a new WebGLBufferRenderer object.
func (t *Three) NewWebGLBufferRenderer(_gl, extensions, _infoRender float64) *WebGLBufferRenderer {
	p := t.ctx.Get("WebGLBufferRenderer").New(_gl, extensions, _infoRender)
	return WebGLBufferRendererFromJSObject(p)
}
