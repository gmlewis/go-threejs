// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// ShaderMaterial represents a shadermaterial.
type ShaderMaterial struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (s *ShaderMaterial) JSObject() *js.Object { return s.p }

// ShaderMaterial returns a ShaderMaterial JavaScript class.
func (t *Three) ShaderMaterial() *ShaderMaterial {
	p := t.ctx.Get("ShaderMaterial")
	return ShaderMaterialFromJSObject(p)
}

// ShaderMaterialFromJSObject returns a wrapped ShaderMaterial JavaScript class.
func ShaderMaterialFromJSObject(p *js.Object) *ShaderMaterial {
	return &ShaderMaterial{p: p}
}

// NewShaderMaterial returns a new ShaderMaterial object.
func (t *Three) NewShaderMaterial(parameters float64) *ShaderMaterial {
	p := t.ctx.Get("ShaderMaterial").New(parameters)
	return ShaderMaterialFromJSObject(p)
}

// Copy TODO description.
func (s *ShaderMaterial) Copy(source *ShaderMaterial) *ShaderMaterial {
	s.p.Call("copy", source.p)
	return s
}

// ToJSON TODO description.
func (s *ShaderMaterial) ToJSON(meta float64) *ShaderMaterial {
	s.p.Call("toJSON", meta)
	return s
}
