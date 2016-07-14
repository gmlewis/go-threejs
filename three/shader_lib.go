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

// ShaderLib returns a ShaderLib JavaScript class.
func (t *Three) ShaderLib() *ShaderLib {
	p := t.ctx.Get("ShaderLib")
	return ShaderLibFromJSObject(p)
}

// ShaderLibFromJSObject returns a wrapped ShaderLib JavaScript class.
func ShaderLibFromJSObject(p *js.Object) *ShaderLib {
	return &ShaderLib{p: p}
}

// NewShaderLib returns a new ShaderLib object.
func (t *Three) NewShaderLib() *ShaderLib {
	p := t.ctx.Get("ShaderLib").New()
	return ShaderLibFromJSObject(p)
}
