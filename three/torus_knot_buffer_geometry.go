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

// TorusKnotBufferGeometry is the BufferGeometry port of TorusKnotGeometry.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/TorusKnotBufferGeometry
type TorusKnotBufferGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *TorusKnotBufferGeometry) JSObject() *js.Object { return t.p }

// TorusKnotBufferGeometry returns a TorusKnotBufferGeometry JavaScript class.
func (t *Three) TorusKnotBufferGeometry() *TorusKnotBufferGeometry {
	p := t.ctx.Get("TorusKnotBufferGeometry")
	return TorusKnotBufferGeometryFromJSObject(p)
}

// TorusKnotBufferGeometryFromJSObject returns a wrapped TorusKnotBufferGeometry JavaScript class.
func TorusKnotBufferGeometryFromJSObject(p *js.Object) *TorusKnotBufferGeometry {
	return &TorusKnotBufferGeometry{p: p}
}

// NewTorusKnotBufferGeometry returns a new TorusKnotBufferGeometry object.
//
//     radius — Default is 100.
//     tube — Diameter of the tube. Default is 40.
//     tubularSegments — Default is 64.
//     radialSegments — Default is 8.
//     p — This value determines, how many times the geometry winds around its axis of rotational symmetry. Default is 2.
//     q — This value determines, how many times the geometry winds around a circle in the interior of the torus. Default is 3.
func (t *Three) NewTorusKnotBufferGeometry(radius, tube, tubularSegments, radialSegments, p, q float64) *TorusKnotBufferGeometry {
	s := t.ctx.Get("TorusKnotBufferGeometry").New(radius, tube, tubularSegments, radialSegments, p, q)
	return TorusKnotBufferGeometryFromJSObject(s)
}
