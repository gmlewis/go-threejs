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
func (t *ParametricGeometry) JSObject() *js.Object { return t.p }

// ParametricGeometry returns a ParametricGeometry object.
func (t *Three) ParametricGeometry() *ParametricGeometry {
	p := t.ctx.Get("ParametricGeometry")
	return &ParametricGeometry{p: p}
}

// New returns a new ParametricGeometry object.
//
//     fn — A function that takes in a u and v value each between 0 and 1 and returns a Vector3
//     slices — The count of slices to use for the parametric function
//     stacks — The count of stacks to use for the parametric function
func (t *ParametricGeometry) New(fn func(u, v float64) *js.Object, slices, stacks int) *ParametricGeometry {
	p := t.p.New(fn, slices, stacks)
	return &ParametricGeometry{p: p}
}
