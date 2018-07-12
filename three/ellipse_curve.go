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

// EllipseCurve represents a 2D curve in the shape of an ellipse.
//
// http://threejs.org/docs/index.html#Reference/Extras.Curves/EllipseCurve
type EllipseCurve struct{ *Curve }

// JSObject returns the underlying *js.Object.
func (e *EllipseCurve) JSObject() *js.Object { return e.p }

// EllipseCurve returns an EllipseCurve JavaScript class.
func (t *Three) EllipseCurve() *EllipseCurve {
	p := t.ctx.Get("EllipseCurve")
	return EllipseCurveFromJSObject(p)
}

// EllipseCurveFromJSObject returns a wrapped EllipseCurve JavaScript class.
func EllipseCurveFromJSObject(p *js.Object) *EllipseCurve {
	return &EllipseCurve{CurveFromJSObject(p)}
}

// EllipseCurveOpts represents options passed to the EllipseCurve constructor.
//     StartAngle – The start angle of the curve in radians starting from the middle right side
//     EndAngle – The end angle of the curve in radians starting from the middle right side
//     Clockwise – Whether the ellipse is clockwise
//     Rotation – The rotation angle of the ellipse in radians, counterclockwise from the positive X axis (optional)
//
//     Note: When going clockwise it's best to set the start angle to (Math.PI * 2) and then work towards lower numbers.
type EllipseCurveOpts struct {
	StartAngle float64
	EndAngle   float64
	Clockwise  bool
	Rotation   float64
}

// NewEllipseCurve returns a new EllipseCurve object.
//
//     aX – The X center of the ellipse
//     aY – The Y center of the ellipse
//     xRadius – The radius of the ellipse in the x direction
//     yRadius – The radius of the ellipse in the y direction
func (t *Three) NewEllipseCurve(aX, aY, xRadius, yRadius float64, opts *EllipseCurveOpts) *EllipseCurve {
	var p *js.Object
	if opts != nil {
		p = t.ctx.Get("EllipseCurve").New(aX, aY, xRadius, yRadius, opts.StartAngle, opts.EndAngle, opts.Clockwise, opts.Rotation)
	} else {
		p = t.ctx.Get("EllipseCurve").New(aX, aY, xRadius, yRadius)
	}
	return EllipseCurveFromJSObject(p)
}

// GetPoint TODO description.
func (e *EllipseCurve) GetPoint(t float64) *EllipseCurve {
	e.p.Call("getPoint", t)
	return e
}
