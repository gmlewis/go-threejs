// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// RawShaderMaterial represents a rawshadermaterial.
type RawShaderMaterial struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (r *RawShaderMaterial) JSObject() *js.Object { return r.p }

// RawShaderMaterial returns a RawShaderMaterial JavaScript class.
func (t *Three) RawShaderMaterial() *RawShaderMaterial {
	p := t.ctx.Get("RawShaderMaterial")
	return RawShaderMaterialFromJSObject(p)
}

// RawShaderMaterialFromJSObject returns a wrapped RawShaderMaterial JavaScript class.
func RawShaderMaterialFromJSObject(p *js.Object) *RawShaderMaterial {
	return &RawShaderMaterial{p: p}
}

// NewRawShaderMaterial returns a new RawShaderMaterial object.
func (t *Three) NewRawShaderMaterial(parameters float64) *RawShaderMaterial {
	p := t.ctx.Get("RawShaderMaterial").New(parameters)
	return RawShaderMaterialFromJSObject(p)
}
