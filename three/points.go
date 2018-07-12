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

// Points is a class for displaying particles in the form of variable size points.
//
// http://threejs.org/docs/index.html#Reference/Objects/Points
type Points struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (p *Points) JSObject() *js.Object { return p.p }

// Points returns a Points JavaScript class.
func (t *Three) Points() *Points {
	p := t.ctx.Get("Points")
	return PointsFromJSObject(p)
}

// PointsFromJSObject returns a wrapped Points JavaScript class.
func PointsFromJSObject(p *js.Object) *Points {
	return &Points{p: p}
}

// NewPoints returns a new Points object.
//
//     geometry — an instance of geometry.
//     material — an instance of material (optional).
func (t *Three) NewPoints(geometry, material *js.Object) *Points {
	p := t.ctx.Get("Points").New(geometry, material)
	return PointsFromJSObject(p)
}
