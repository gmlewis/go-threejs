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
func (t *ShaderLib) JSObject() *js.Object { return t.p }

// ShaderLib returns a ShaderLib object.
func (t *Three) ShaderLib() *ShaderLib {
	p := t.ctx.Get("ShaderLib")
	return &ShaderLib{p: p}
}

// New returns a new ShaderLib object.
func (t *ShaderLib) New() *ShaderLib {
	p := t.p.New()
	return &ShaderLib{p: p}
}
