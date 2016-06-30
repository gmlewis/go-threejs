// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// ImmediateRenderObject represents an immediaterenderobject.
type ImmediateRenderObject struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *ImmediateRenderObject) JSObject() *js.Object { return t.p }

// ImmediateRenderObject returns an ImmediateRenderObject object.
func (t *Three) ImmediateRenderObject() *ImmediateRenderObject {
	p := t.ctx.Get("ImmediateRenderObject")
	return &ImmediateRenderObject{p: p}
}

// New returns a new ImmediateRenderObject object.
func (t *ImmediateRenderObject) New(material float64) *ImmediateRenderObject {
	p := t.p.New(material)
	return &ImmediateRenderObject{p: p}
}
