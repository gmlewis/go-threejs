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

// AxisHelper returns an AxisHelper object.
func (t *Three) AxisHelper() *AxisHelper {
	p := t.ctx.Get("AxisHelper")
	return &AxisHelper{p: p}
}

// New returns a new AxisHelper object.
func (a *AxisHelper) New(size float64) *AxisHelper {
	p := a.p.New(size)
	return &AxisHelper{p: p}
}
