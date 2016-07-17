// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// AxisHelper is an axis object to visualize the 3 axes
// in a simple way. The X axis is red. The Y axis is green.
// The Z axis is blue.
//
// http://threejs.org/docs/index.html#Reference/Extras.Helpers/AxisHelper
type AxisHelper struct{ *LineSegments }

// JSObject returns the underlying *js.Object.
func (a *AxisHelper) JSObject() *js.Object { return a.p }

// AxisHelper returns an AxisHelper JavaScript class.
func (t *Three) AxisHelper() *AxisHelper {
	p := t.ctx.Get("AxisHelper")
	return AxisHelperFromJSObject(p)
}

// AxisHelperFromJSObject returns a wrapper AxisHelper JavaScript class.
func AxisHelperFromJSObject(p *js.Object) *AxisHelper {
	return &AxisHelper{LineSegmentsFromJSObject(p)}
}

// NewAxisHelper returns a new AxisHelper object.
func (t *Three) NewAxisHelper(size float64) *AxisHelper {
	p := t.ctx.Get("AxisHelper").New(size)
	return AxisHelperFromJSObject(p)
}
