// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WireframeHelper represents a wireframehelper.
type WireframeHelper struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WireframeHelper) JSObject() *js.Object { return w.p }

// WireframeHelper returns a WireframeHelper JavaScript class.
func (t *Three) WireframeHelper() *WireframeHelper {
	p := t.ctx.Get("WireframeHelper")
	return &WireframeHelper{p: p}
}

// NewWireframeHelper returns a new WireframeHelper object.
func (t *Three) NewWireframeHelper(object, hex float64) *WireframeHelper {
	p := t.ctx.Get("WireframeHelper").New(object, hex)
	return &WireframeHelper{p: p}
}
