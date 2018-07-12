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

// LatheBufferGeometry represents a lathebuffergeometry.
type LatheBufferGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (l *LatheBufferGeometry) JSObject() *js.Object { return l.p }

// LatheBufferGeometry returns a LatheBufferGeometry JavaScript class.
func (t *Three) LatheBufferGeometry() *LatheBufferGeometry {
	p := t.ctx.Get("LatheBufferGeometry")
	return LatheBufferGeometryFromJSObject(p)
}

// LatheBufferGeometryFromJSObject returns a wrapped LatheBufferGeometry JavaScript class.
func LatheBufferGeometryFromJSObject(p *js.Object) *LatheBufferGeometry {
	return &LatheBufferGeometry{p: p}
}

// NewLatheBufferGeometry returns a new LatheBufferGeometry object.
func (t *Three) NewLatheBufferGeometry(points, segments, phiStart, phiLength float64) *LatheBufferGeometry {
	p := t.ctx.Get("LatheBufferGeometry").New(points, segments, phiStart, phiLength)
	return LatheBufferGeometryFromJSObject(p)
}
