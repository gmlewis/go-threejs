// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// PlaneGeometry represent a plane.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/PlaneGeometry
type PlaneGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *PlaneGeometry) JSObject() *js.Object { return t.p }

// PlaneGeometry returns a PlaneGeometry object.
func (t *Three) PlaneGeometry() *PlaneGeometry {
	p := t.ctx.Get("PlaneGeometry")
	return &PlaneGeometry{p: p}
}

// New returns a new PlaneGeometry object.
//
//     width — Width along the X axis.
//     height — Height along the Y axis.
//     widthSegments — Optional. Default is 1.
//     heightSegments — Optional. Default is 1.
func (t *PlaneGeometry) New(width, height float64, widthSegments, heightSegments int) *PlaneGeometry {
	p := t.p.New(width, height, widthSegments, heightSegments)
	return &PlaneGeometry{p: p}
}
