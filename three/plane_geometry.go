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

// PlaneGeometry represent a plane.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/PlaneGeometry
type PlaneGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (p *PlaneGeometry) JSObject() *js.Object { return p.p }

// PlaneGeometry returns a PlaneGeometry JavaScript class.
func (t *Three) PlaneGeometry() *PlaneGeometry {
	p := t.ctx.Get("PlaneGeometry")
	return PlaneGeometryFromJSObject(p)
}

// PlaneGeometryFromJSObject returns a wrapped PlaneGeometry JavaScript class.
func PlaneGeometryFromJSObject(p *js.Object) *PlaneGeometry {
	return &PlaneGeometry{p: p}
}

// NewPlaneGeometry returns a new PlaneGeometry object.
//
//     width — Width along the X axis.
//     height — Height along the Y axis.
//     widthSegments — Optional. Default is 1.
//     heightSegments — Optional. Default is 1.
func (t *Three) NewPlaneGeometry(width, height float64, widthSegments, heightSegments int) *PlaneGeometry {
	p := t.ctx.Get("PlaneGeometry").New(width, height, widthSegments, heightSegments)
	return PlaneGeometryFromJSObject(p)
}
