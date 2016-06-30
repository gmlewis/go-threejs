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

// WireframeHelper returns a WireframeHelper object.
func (t *Three) WireframeHelper() *WireframeHelper {
	p := t.ctx.Get("WireframeHelper")
	return &WireframeHelper{p: p}
}

// New returns a new WireframeHelper object.
func (w *WireframeHelper) New(object, hex float64) *WireframeHelper {
	p := w.p.New(object, hex)
	return &WireframeHelper{p: p}
}
