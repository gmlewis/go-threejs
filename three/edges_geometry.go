// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// EdgesGeometry represents an edgesgeometry.
type EdgesGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *EdgesGeometry) JSObject() *js.Object { return t.p }

// EdgesGeometry returns an EdgesGeometry object.
func (t *Three) EdgesGeometry() *EdgesGeometry {
	p := t.ctx.Get("EdgesGeometry")
	return &EdgesGeometry{p: p}
}

// New returns a new EdgesGeometry object.
func (t *EdgesGeometry) New(geometry, thresholdAngle float64) *EdgesGeometry {
	p := t.p.New(geometry, thresholdAngle)
	return &EdgesGeometry{p: p}
}
