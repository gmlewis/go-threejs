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

// TorusGeometry represents a torus.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/TorusGeometry
type TorusGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *TorusGeometry) JSObject() *js.Object { return t.p }

// TorusGeometry returns a TorusGeometry JavaScript class.
func (t *Three) TorusGeometry() *TorusGeometry {
	p := t.ctx.Get("TorusGeometry")
	return TorusGeometryFromJSObject(p)
}

// TorusGeometryFromJSObject returns a wrapped TorusGeometry JavaScript class.
func TorusGeometryFromJSObject(p *js.Object) *TorusGeometry {
	return &TorusGeometry{p: p}
}

// NewTorusGeometry returns a new TorusGeometry object.
//
//     radius — Default is 100.
//     tube — Diameter of the tube. Default is 40.
//     radialSegments — Default is 8
//     tubularSegments — Default is 6.
//     arc — Central angle. Default is Math.PI * 2.
func (t *Three) NewTorusGeometry(radius, tube float64, radialSegments, tubularSegments int, arc float64) *TorusGeometry {
	p := t.ctx.Get("TorusGeometry").New(radius, tube, radialSegments, tubularSegments, arc)
	return TorusGeometryFromJSObject(p)
}
