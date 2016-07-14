// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// ShaderChunk represents a shaderchunk.
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
