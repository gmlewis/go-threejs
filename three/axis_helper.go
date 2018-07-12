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

// AxisHelper is an axis object to visualize the 3 axes
// in a simple way. The X axis is red. The Y axis is green.
// The Z axis is blue.
//
// http://threejs.org/docs/index.html#Reference/Extras.Helpers/AxisHelper
type AxisHelper struct{ *LineSegments }

// JSObject returns the underlying *js.Object.
func (a *AxisHelper) JSObject() *js.Object { return a.p }

// AxisHelper returns an AxisHelper JavaScript class.
func (t *Three) AxisHelper() *AxisHelper {
	p := t.ctx.Get("AxisHelper")
	return AxisHelperFromJSObject(p)
}

// AxisHelperFromJSObject returns a wrapper AxisHelper JavaScript class.
func AxisHelperFromJSObject(p *js.Object) *AxisHelper {
	return &AxisHelper{LineSegmentsFromJSObject(p)}
}

// NewAxisHelper returns a new AxisHelper object.
func (t *Three) NewAxisHelper(size float64) *AxisHelper {
	p := t.ctx.Get("AxisHelper").New(size)
	return AxisHelperFromJSObject(p)
}
