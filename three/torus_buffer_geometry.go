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

// TorusBufferGeometry represents a torusbuffergeometry.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/TorusBufferGeometry
type TorusBufferGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *TorusBufferGeometry) JSObject() *js.Object { return t.p }

// TorusBufferGeometry returns a TorusBufferGeometry JavaScript class.
func (t *Three) TorusBufferGeometry() *TorusBufferGeometry {
	p := t.ctx.Get("TorusBufferGeometry")
	return TorusBufferGeometryFromJSObject(p)
}

// TorusBufferGeometryFromJSObject returns a wrapped TorusBufferGeometry JavaScript class.
func TorusBufferGeometryFromJSObject(p *js.Object) *TorusBufferGeometry {
	return &TorusBufferGeometry{p: p}
}

// NewTorusBufferGeometry returns a new TorusBufferGeometry object.
//
//     radius — Default is 100.
//     tube — Diameter of the tube. Default is 40.
//     radialSegments — Default is 8
//     tubularSegments — Default is 6.
//     arc — Central angle. Default is Math.PI * 2.
func (t *Three) NewTorusBufferGeometry(radius, tube, radialSegments, tubularSegments, arc float64) *TorusBufferGeometry {
	p := t.ctx.Get("TorusBufferGeometry").New(radius, tube, radialSegments, tubularSegments, arc)
	return TorusBufferGeometryFromJSObject(p)
}
