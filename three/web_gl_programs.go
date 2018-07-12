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

// WebGLPrograms represents a webglprograms.
type WebGLPrograms struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WebGLPrograms) JSObject() *js.Object { return w.p }

// WebGLPrograms returns a WebGLPrograms JavaScript class.
func (t *Three) WebGLPrograms() *WebGLPrograms {
	p := t.ctx.Get("WebGLPrograms")
	return WebGLProgramsFromJSObject(p)
}

// WebGLProgramsFromJSObject returns a wrapped WebGLPrograms JavaScript class.
func WebGLProgramsFromJSObject(p *js.Object) *WebGLPrograms {
	return &WebGLPrograms{p: p}
}

// NewWebGLPrograms returns a new WebGLPrograms object.
func (t *Three) NewWebGLPrograms(renderer, capabilities float64) *WebGLPrograms {
	p := t.ctx.Get("WebGLPrograms").New(renderer, capabilities)
	return WebGLProgramsFromJSObject(p)
}
