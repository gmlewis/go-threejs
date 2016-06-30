// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// BoundingBoxHelper represents a boundingboxhelper.
type BoundingBoxHelper struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (b *BoundingBoxHelper) JSObject() *js.Object { return b.p }

// BoundingBoxHelper returns a BoundingBoxHelper object.
func (t *Three) BoundingBoxHelper() *BoundingBoxHelper {
	p := t.ctx.Get("BoundingBoxHelper")
	return &BoundingBoxHelper{p: p}
}

// New returns a new BoundingBoxHelper object.
func (b *BoundingBoxHelper) New(object, hex float64) *BoundingBoxHelper {
	p := b.p.New(object, hex)
	return &BoundingBoxHelper{p: p}
}
