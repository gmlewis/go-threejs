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

// PlaneBufferGeometry represents a planebuffergeometry.
type PlaneBufferGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (p *PlaneBufferGeometry) JSObject() *js.Object { return p.p }

// PlaneBufferGeometry returns a PlaneBufferGeometry JavaScript class.
func (t *Three) PlaneBufferGeometry() *PlaneBufferGeometry {
	p := t.ctx.Get("PlaneBufferGeometry")
	return PlaneBufferGeometryFromJSObject(p)
}

// PlaneBufferGeometryFromJSObject returns a wrapped PlaneBufferGeometry JavaScript class.
func PlaneBufferGeometryFromJSObject(p *js.Object) *PlaneBufferGeometry {
	return &PlaneBufferGeometry{p: p}
}

// NewPlaneBufferGeometry returns a new PlaneBufferGeometry object.
func (t *Three) NewPlaneBufferGeometry(width, height, widthSegments, heightSegments float64) *PlaneBufferGeometry {
	p := t.ctx.Get("PlaneBufferGeometry").New(width, height, widthSegments, heightSegments)
	return PlaneBufferGeometryFromJSObject(p)
}
