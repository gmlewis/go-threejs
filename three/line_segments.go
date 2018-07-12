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

// LineSegments represents a series of lines.
//
// http://threejs.org/docs/index.html#Reference/Objects/LineSegments
type LineSegments struct{ *Line }

// JSObject returns the underlying *js.Object.
func (l *LineSegments) JSObject() *js.Object { return l.p }

// LineSegments returns a LineSegments JavaScript class.
func (t *Three) LineSegments() *LineSegments {
	p := t.ctx.Get("LineSegments")
	return LineSegmentsFromJSObject(p)
}

// LineSegmentsFromJSObject returns a wrapper LineSegments JavaScript class.
func LineSegmentsFromJSObject(p *js.Object) *LineSegments {
	return &LineSegments{LineFromJSObject(p)}
}

// NewLineSegments returns a new LineSegments object.
//
//     geometry — Vertices representing the line segment(s).
//     material — Material for the line. Default is LineBasicMaterial.
func (t *Three) NewLineSegments(geometry []*js.Object, material *js.Object) *LineSegments {
	p := t.ctx.Get("LineSegments").New(geometry, material)
	return LineSegmentsFromJSObject(p)
}
