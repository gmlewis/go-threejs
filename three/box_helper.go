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
func (b *BoxHelper) JSObject() *js.Object { return b.p }

// BoxHelper returns a BoxHelper JavaScript class.
func (t *Three) BoxHelper() *BoxHelper {
	p := t.ctx.Get("BoxHelper")
	return BoxHelperFromJSObject(p)
}

// BoxHelperFromJSObject returns a wrapped BoxHelper JavaScript class.
func BoxHelperFromJSObject(p *js.Object) *BoxHelper {
	return &BoxHelper{p: p}
}

// NewBoxHelper returns a new BoxHelper object.
func (t *Three) NewBoxHelper(object float64) *BoxHelper {
	p := t.ctx.Get("BoxHelper").New(object)
	return BoxHelperFromJSObject(p)
}
