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

// Line represents a continuous line.
//
// http://threejs.org/docs/index.html#Reference/Objects/Line
type Line struct{ *Object3D }

// JSObject returns the underlying *js.Object.
func (l *Line) JSObject() *js.Object { return l.p }

// Line returns a Line JavaScript class.
func (t *Three) Line() *Line {
	p := t.ctx.Get("Line")
	return LineFromJSObject(p)
}

// LineFromJSObject returns a wrapper Line JavaScript class.
func LineFromJSObject(p *js.Object) *Line {
	return &Line{Object3DFromJSObject(p)}
}

// NewLine returns a new Line object.
//
//     geometry — Vertices representing the line segment(s).
//     material — Material for the line. Default is LineBasicMaterial.
func (t *Three) NewLine(geometry JSObject, material JSObject) *Line {
	p := t.ctx.Get("Line").New(geometry.JSObject(), material.JSObject())
	return LineFromJSObject(p)
}
