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

// CatmullRomCurve3 represents a smooth 3D spline curve from
// a series of points using the Catmull-Rom algorithm.
//
// http://threejs.org/docs/index.html#Reference/Extras.Curves/CatmullRomCurve3
type CatmullRomCurve3 struct{ *Curve }

// JSObject returns the underlying *js.Object.
func (c *CatmullRomCurve3) JSObject() *js.Object { return c.p }

// CatmullRomCurve3 returns a CatmullRomCurve3 JavaScript class.
func (t *Three) CatmullRomCurve3() *CatmullRomCurve3 {
	p := t.ctx.Get("CatmullRomCurve3")
	return CatmullRomCurve3FromJSObject(p)
}

// CatmullRomCurve3FromJSObject returns a wrapped CatmullRomCurve3 JavaScript class.
func CatmullRomCurve3FromJSObject(p *js.Object) *CatmullRomCurve3 {
	return &CatmullRomCurve3{CurveFromJSObject(p)}
}

// NewCatmullRomCurve3 returns a new CatmullRomCurve3 object.
func (t *Three) NewCatmullRomCurve3(pts []*js.Object) *CatmullRomCurve3 {
	p := t.ctx.Get("CatmullRomCurve3").New(pts)
	return CatmullRomCurve3FromJSObject(p)
}
