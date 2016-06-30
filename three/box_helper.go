// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// BoxHelper represents a boxhelper.
type BoxHelper struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *BoxHelper) JSObject() *js.Object { return t.p }

// BoxHelper returns a BoxHelper object.
func (t *Three) BoxHelper() *BoxHelper {
	p := t.ctx.Get("BoxHelper")
	return &BoxHelper{p: p}
}

// New returns a new BoxHelper object.
func (t *BoxHelper) New(object float64) *BoxHelper {
	p := t.p.New(object)
	return &BoxHelper{p: p}
}
