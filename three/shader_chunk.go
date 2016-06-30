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
func (t *ShaderChunk) JSObject() *js.Object { return t.p }

// ShaderChunk returns a ShaderChunk object.
func (t *Three) ShaderChunk() *ShaderChunk {
	p := t.ctx.Get("ShaderChunk")
	return &ShaderChunk{p: p}
}

// New returns a new ShaderChunk object.
func (t *ShaderChunk) New() *ShaderChunk {
	p := t.p.New()
	return &ShaderChunk{p: p}
}
