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

// SphereBufferGeometry is the BufferGeometry port of SphereGeometry.
type SphereBufferGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (s *SphereBufferGeometry) JSObject() *js.Object { return s.p }

// SphereBufferGeometry returns a SphereBufferGeometry JavaScript class.
func (t *Three) SphereBufferGeometry() *SphereBufferGeometry {
	p := t.ctx.Get("SphereBufferGeometry")
	return SphereBufferGeometryFromJSObject(p)
}

// SphereBufferGeometryFromJSObject returns a wrapped SphereBufferGeometry JavaScript class.
func SphereBufferGeometryFromJSObject(p *js.Object) *SphereBufferGeometry {
	return &SphereBufferGeometry{p: p}
}

// NewSphereBufferGeometry returns a new SphereBufferGeometry object.
//
//     radius — sphere radius. Default is 50.
//     widthSegments — number of horizontal segments. Minimum value is 3, and the default is 8.
//     heightSegments — number of vertical segments. Minimum value is 2, and the default is 6.
//     phiStart — specify horizontal starting angle. Default is 0.
//     phiLength — specify horizontal sweep angle size. Default is Math.PI * 2.
//     thetaStart — specify vertical starting angle. Default is 0.
//     thetaLength — specify vertical sweep angle size. Default is Math.PI.
func (t *Three) NewSphereBufferGeometry(radius, widthSegments, heightSegments, phiStart, phiLength, thetaStart, thetaLength float64) *SphereBufferGeometry {
	p := t.ctx.Get("SphereBufferGeometry").New(radius, widthSegments, heightSegments, phiStart, phiLength, thetaStart, thetaLength)
	return SphereBufferGeometryFromJSObject(p)
}
