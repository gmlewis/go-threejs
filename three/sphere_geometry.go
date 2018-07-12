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

// SphereGeometry represents a sphere.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/SphereGeometry
type SphereGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (s *SphereGeometry) JSObject() *js.Object { return s.p }

// SphereGeometry returns a SphereGeometry JavaScript class.
func (t *Three) SphereGeometry() *SphereGeometry {
	p := t.ctx.Get("SphereGeometry")
	return SphereGeometryFromJSObject(p)
}

// SphereGeometryFromJSObject returns a wrapped SphereGeometry JavaScript class.
func SphereGeometryFromJSObject(p *js.Object) *SphereGeometry {
	return &SphereGeometry{p: p}
}

// SphereGeometryOpts represents options for constructing a SphereGeometry.
//
//     phiStart — specify horizontal starting angle. Default is 0.
//     phiLength — specify horizontal sweep angle size. Default is Math.PI * 2.
//     thetaStart — specify vertical starting angle. Default is 0.
//     thetaLength — specify vertical sweep angle size. Default is Math.PI.
type SphereGeometryOpts struct {
	phiStart    float64
	phiLength   float64
	thetaStart  float64
	thetaLength float64
}

// NewSphereGeometry returns a new SphereGeometry object.
//
//     radius — sphere radius. Default is 50.
//     widthSegments — number of horizontal segments. Minimum value is 3, and the default is 8.
//     heightSegments — number of vertical segments. Minimum value is 2, and the default is 6.
func (t *Three) NewSphereGeometry(radius float64, widthSegments, heightSegments int, opts *SphereGeometryOpts) *SphereGeometry {
	var p *js.Object
	if opts != nil {
		p = t.ctx.Get("SphereGeometry").New(radius, widthSegments, heightSegments, opts.phiStart, opts.phiLength, opts.thetaStart, opts.thetaLength)
	} else {
		p = t.ctx.Get("SphereGeometry").New(radius, widthSegments, heightSegments)
	}
	return SphereGeometryFromJSObject(p)
}
