// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WireframeGeometry represents a wireframegeometry.
type WireframeGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WireframeGeometry) JSObject() *js.Object { return w.p }

// WireframeGeometry returns a WireframeGeometry object.
func (t *Three) WireframeGeometry() *WireframeGeometry {
	p := t.ctx.Get("WireframeGeometry")
	return &WireframeGeometry{p: p}
}

// New returns a new WireframeGeometry object.
func (w *WireframeGeometry) New(geometry float64) *WireframeGeometry {
	p := w.p.New(geometry)
	return &WireframeGeometry{p: p}
}
