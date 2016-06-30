// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// ArrowHelper represents an arrowhelper.
type ArrowHelper struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (a *ArrowHelper) JSObject() *js.Object { return a.p }

// ArrowHelper returns an ArrowHelper object.
func (t *Three) ArrowHelper() *ArrowHelper {
	p := t.ctx.Get("ArrowHelper")
	return &ArrowHelper{p: p}
}

// New returns a new ArrowHelper object.
func (a *ArrowHelper) New() *ArrowHelper {
	p := a.p.New()
	return &ArrowHelper{p: p}
}

// SetLength TODO description.
func (a *ArrowHelper) SetLength(length, headLength, headWidth float64) *ArrowHelper {
	a.p.Call("setLength", length, headLength, headWidth)
	return a
}

// SetColor TODO description.
func (a *ArrowHelper) SetColor(color float64) *ArrowHelper {
	a.p.Call("setColor", color)
	return a
}
