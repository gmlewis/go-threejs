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

// ArrowHelper represents a 3D arrow object.
//
// http://threejs.org/docs/index.html#Reference/Extras.Helpers/ArrowHelper
type ArrowHelper struct{ *Object3D }

// JSObject returns the underlying *js.Object.
func (a *ArrowHelper) JSObject() *js.Object { return a.p }

// ArrowHelper returns an ArrowHelper JavaScript class.
func (t *Three) ArrowHelper() *ArrowHelper {
	p := t.ctx.Get("ArrowHelper")
	return ArrowHelperFromJSObject(p)
}

// ArrowHelperFromJSObject returns a wrapped ArrowHelper JavaScript class.
func ArrowHelperFromJSObject(p *js.Object) *ArrowHelper {
	return &ArrowHelper{Object3DFromJSObject(p)}
}

// ArrowHelperOpts represents optional arguments that can be passed to
// NewArrowHelper.
type ArrowHelperOpts struct {
	Length     *float64
	Color      *Color
	HeadLength *float64
	HeadWidth  *float64
}

// NewArrowHelper returns a new ArrowHelper object.
func (t *Three) NewArrowHelper(dir, origin *Vector3, opts *ArrowHelperOpts) *ArrowHelper {
	var (
		length     interface{} = js.Undefined
		color      interface{} = js.Undefined
		headLength interface{} = js.Undefined
		headWidth  interface{} = js.Undefined
	)
	if opts != nil {
		if opts.Length != nil {
			length = *opts.Length
		}
		if opts.Color != nil {
			color = opts.Color.JSObject()
		}
		if opts.HeadLength != nil {
			headLength = *opts.HeadLength
		}
		if opts.HeadWidth != nil {
			headWidth = *opts.HeadWidth
		}
	}
	p := t.ctx.Get("ArrowHelper").New(dir.JSObject(), origin.JSObject(), length, color, headLength, headWidth)
	return ArrowHelperFromJSObject(p)
}

// SetLength TODO description.
func (a *ArrowHelper) SetLength(length, headLength, headWidth float64) *ArrowHelper {
	a.p.Call("setLength", length, headLength, headWidth)
	return a
}

// SetColor TODO description.
func (a *ArrowHelper) SetColor(color float64) *ArrowHelper {
	a.p.Call("setColor", color)
	return a
}
