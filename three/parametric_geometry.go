// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// ParametricGeometry represents geometry defined by a parametric surface.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/ParametricGeometry
type ParametricGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (p *ParametricGeometry) JSObject() *js.Object { return p.p }

// ParametricGeometry returns a ParametricGeometry JavaScript class.
func (t *Three) ParametricGeometry() *ParametricGeometry {
	p := t.ctx.Get("ParametricGeometry")
	return &ParametricGeometry{p: p}
}

// NewParametricGeometry returns a new ParametricGeometry object.
//
//     fn — A function that takes in a u and v value each between 0 and 1 and returns a Vector3
//     slices — The count of slices to use for the parametric function
//     stacks — The count of stacks to use for the parametric function
func (t *Three) NewParametricGeometry(fn func(u, v float64) *js.Object, slices, stacks int) *ParametricGeometry {
	return &ParametricGeometry{p: t.ctx.Get("ParametricGeometry").New(fn, slices, stacks)}
}
