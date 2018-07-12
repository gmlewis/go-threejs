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

// DodecahedronGeometry represents a dodecagedron.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/DodecahedronGeometry
type DodecahedronGeometry struct{ *PolyhedronGeometry }

// JSObject returns the underlying *js.Object.
func (d *DodecahedronGeometry) JSObject() *js.Object { return d.p }

// DodecahedronGeometry returns a DodecahedronGeometry JavaScript class.
func (t *Three) DodecahedronGeometry() *DodecahedronGeometry {
	p := t.ctx.Get("DodecahedronGeometry")
	return DodecahedronGeometryFromJSObject(p)
}

// DodecahedronGeometryFromJSObject returns a wrapped DodecahedronGeometry JavaScript class.
func DodecahedronGeometryFromJSObject(p *js.Object) *DodecahedronGeometry {
	return &DodecahedronGeometry{PolyhedronGeometryFromJSObject(p)}
}

// NewDodecahedronGeometry returns a new DodecahedronGeometry object.
//
//     radius — Radius of the dodecahedron. Default is 1.
//     detail — Default is 0. Setting this to a value greater than 0 adds
//         vertices making it no longer a dodecahedron.
func (t *Three) NewDodecahedronGeometry(radius, detail float64) *DodecahedronGeometry {
	p := t.ctx.Get("DodecahedronGeometry").New(radius, detail)
	return DodecahedronGeometryFromJSObject(p)
}
