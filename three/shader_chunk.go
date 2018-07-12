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

// ShaderChunk represents a shaderchunk.
//
// http://threejs.org/docs/index.html#Reference/Renderers.Shaders/ShaderChunk
type ShaderChunk struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (s *ShaderChunk) JSObject() *js.Object { return s.p }

// ShaderChunk returns a ShaderChunk JavaScript class.
func (t *Three) ShaderChunk() *ShaderChunk {
	p := t.ctx.Get("ShaderChunk")
	return ShaderChunkFromJSObject(p)
}

// ShaderChunkFromJSObject returns a wrapped ShaderChunk JavaScript class.
func ShaderChunkFromJSObject(p *js.Object) *ShaderChunk {
	return &ShaderChunk{p: p}
}

// NewShaderChunk returns a new ShaderChunk object.
func (t *Three) NewShaderChunk() *ShaderChunk {
	p := t.ctx.Get("ShaderChunk").New()
	return ShaderChunkFromJSObject(p)
}
