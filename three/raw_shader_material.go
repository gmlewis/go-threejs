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

// RawShaderMaterial returns a RawShaderMaterial object.
func (t *Three) RawShaderMaterial() *RawShaderMaterial {
	p := t.ctx.Get("RawShaderMaterial")
	return &RawShaderMaterial{p: p}
}

// New returns a new RawShaderMaterial object.
func (r *RawShaderMaterial) New(parameters float64) *RawShaderMaterial {
	p := r.p.New(parameters)
	return &RawShaderMaterial{p: p}
}
