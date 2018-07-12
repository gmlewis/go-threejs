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

// RingBufferGeometry is the BufferGeometry port of RingGeometry.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/RingBufferGeometry
type RingBufferGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (r *RingBufferGeometry) JSObject() *js.Object { return r.p }

// RingBufferGeometry returns a RingBufferGeometry JavaScript class.
func (t *Three) RingBufferGeometry() *RingBufferGeometry {
	p := t.ctx.Get("RingBufferGeometry")
	return RingBufferGeometryFromJSObject(p)
}

// RingBufferGeometryFromJSObject returns a wrapped RingBufferGeometry JavaScript class.
func RingBufferGeometryFromJSObject(p *js.Object) *RingBufferGeometry {
	return &RingBufferGeometry{p: p}
}

// NewRingBufferGeometry returns a new RingBufferGeometry object.
//
//     innerRadius — Default is 0, but it doesn't work right when innerRadius is set to 0.
//     outerRadius — Default is 50.
//     thetaSegments — Number of segments. A higher number means the ring will be more round. Minimum is 3. Default is 8.
//     phiSegments — Minimum is 1. Default is 8.
//     thetaStart — Starting angle. Default is 0.
//     thetaLength — Central angle. Default is Math.PI * 2.
func (t *Three) NewRingBufferGeometry(innerRadius, outerRadius, thetaSegments, phiSegments, thetaStart, thetaLength float64) *RingBufferGeometry {
	p := t.ctx.Get("RingBufferGeometry").New(innerRadius, outerRadius, thetaSegments, phiSegments, thetaStart, thetaLength)
	return RingBufferGeometryFromJSObject(p)
}
