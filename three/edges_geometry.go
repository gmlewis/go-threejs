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

// EdgesGeometry represents an edgesgeometry.
type EdgesGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (e *EdgesGeometry) JSObject() *js.Object { return e.p }

// EdgesGeometry returns an EdgesGeometry JavaScript class.
func (t *Three) EdgesGeometry() *EdgesGeometry {
	p := t.ctx.Get("EdgesGeometry")
	return EdgesGeometryFromJSObject(p)
}

// EdgesGeometryFromJSObject returns a wrapped EdgesGeometry JavaScript class.
func EdgesGeometryFromJSObject(p *js.Object) *EdgesGeometry {
	return &EdgesGeometry{p: p}
}

// NewEdgesGeometry returns a new EdgesGeometry object.
func (t *Three) NewEdgesGeometry(geometry, thresholdAngle float64) *EdgesGeometry {
	p := t.ctx.Get("EdgesGeometry").New(geometry, thresholdAngle)
	return EdgesGeometryFromJSObject(p)
}
