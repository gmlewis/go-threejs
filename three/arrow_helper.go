// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// ArrowHelper represents an arrowhelper.
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

// NewArrowHelperOpts represents optional arguments that can be passed to
// NewArrowHelper.
type NewArrowHelperOpts struct {
	Length     *float64
	Color      *Color
	HeadLength *float64
	HeadWidth  *float64
}

// NewArrowHelper returns a new ArrowHelper object.
func (t *Three) NewArrowHelper(dir, origin *Vector3, opts *NewArrowHelperOpts) *ArrowHelper {
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
