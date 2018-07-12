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

// IcosahedronGeometry represents an icosahedron geometry.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/IcosahedronGeometry
type IcosahedronGeometry struct{ *PolyhedronGeometry }

// JSObject returns the underlying *js.Object.
func (i *IcosahedronGeometry) JSObject() *js.Object { return i.p }

// IcosahedronGeometry returns an IcosahedronGeometry JavaScript class.
func (t *Three) IcosahedronGeometry() *IcosahedronGeometry {
	p := t.ctx.Get("IcosahedronGeometry")
	return IcosahedronGeometryFromJSObject(p)
}

// IcosahedronGeometryFromJSObject returns a wrapped IcosahedronGeometry JavaScript class.
func IcosahedronGeometryFromJSObject(p *js.Object) *IcosahedronGeometry {
	return &IcosahedronGeometry{PolyhedronGeometryFromJSObject(p)}
}

// NewIcosahedronGeometry returns a new IcosahedronGeometry object.
//
//     radius — Default is 1.
//     detail — Default is 0. Setting this to a value greater than 0 adds more
//         vertices making it no longer an icosahedron. When detail is greater
//         than 1, it's effectively a sphere.
func (t *Three) NewIcosahedronGeometry(radius, detail float64) *IcosahedronGeometry {
	p := t.ctx.Get("IcosahedronGeometry").New(radius, detail)
	return IcosahedronGeometryFromJSObject(p)
}
