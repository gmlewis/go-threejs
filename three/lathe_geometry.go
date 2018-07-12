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

// LatheGeometry represents a mesh with axial symmetry.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/LatheGeometry
type LatheGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (l *LatheGeometry) JSObject() *js.Object { return l.p }

// LatheGeometry returns a LatheGeometry JavaScript class.
func (t *Three) LatheGeometry() *LatheGeometry {
	p := t.ctx.Get("LatheGeometry")
	return LatheGeometryFromJSObject(p)
}

// LatheGeometryFromJSObject returns a wrapped LatheGeometry JavaScript class.
func LatheGeometryFromJSObject(p *js.Object) *LatheGeometry {
	return &LatheGeometry{p: p}
}

// NewLatheGeometry returns a new LatheGeometry object.
//
//     points — Slice of Vector2s.
//     segments — the number of circumference segments to generate. Default is 12.
//     phiStart — the starting angle in radians. Default is 0.
//     phiLength — the radian (0 to 2PI) range of the lathed section 2PI is a
//         closed lathe, less than 2PI is a portion. Default is 2*PI
func (t *Three) NewLatheGeometry(points []*js.Object, segments int, phiStart, phiLength float64) *LatheGeometry {
	p := t.ctx.Get("LatheGeometry").New(points, segments, phiStart, phiLength)
	return LatheGeometryFromJSObject(p)
}
