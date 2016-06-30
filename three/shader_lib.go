// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// ShaderLib represents a shaderlib.
type ShaderLib struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (s *ShaderLib) JSObject() *js.Object { return s.p }

// ShaderLib returns a ShaderLib object.
func (t *Three) ShaderLib() *ShaderLib {
	p := t.ctx.Get("ShaderLib")
	return &ShaderLib{p: p}
}

// New returns a new ShaderLib object.
func (s *ShaderLib) New() *ShaderLib {
	p := s.p.New()
	return &ShaderLib{p: p}
}
