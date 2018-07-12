// Copyright 2016 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	return ParametricGeometryFromJSObject(p)
}

// ParametricGeometryFromJSObject returns a wrapped ParametricGeometry JavaScript class.
func ParametricGeometryFromJSObject(p *js.Object) *ParametricGeometry {
	return &ParametricGeometry{p: p}
}

// NewParametricGeometry returns a new ParametricGeometry object.
//
//     fn — A function that takes in a u and v value each between 0 and 1 and returns a Vector3
//     slices — The count of slices to use for the parametric function
//     stacks — The count of stacks to use for the parametric function
func (t *Three) NewParametricGeometry(fn func(u, v float64) *js.Object, slices, stacks int) *ParametricGeometry {
	p := t.ctx.Get("ParametricGeometry").New(fn, slices, stacks)
	return ParametricGeometryFromJSObject(p)
}
