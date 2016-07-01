// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// AxisHelper represents an axishelper.
type AxisHelper struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (a *AxisHelper) JSObject() *js.Object { return a.p }

// AxisHelper returns an AxisHelper JavaScript class.
func (t *Three) AxisHelper() *AxisHelper {
	p := t.ctx.Get("AxisHelper")
	return &AxisHelper{p: p}
}

// NewAxisHelper returns a new AxisHelper object.
func (t *Three) NewAxisHelper(size float64) *AxisHelper {
	p := t.ctx.Get("AxisHelper").New(size)
	return &AxisHelper{p: p}
}
