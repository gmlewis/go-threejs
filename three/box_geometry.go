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

// BoxGeometry represents a quadrilateral geometry primitive.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/BoxGeometry
type BoxGeometry struct{ *Geometry }

// JSObject returns the underlying *js.Object.
func (b *BoxGeometry) JSObject() *js.Object { return b.p }

// BoxGeometry returns a BoxGeometry JavaScript class.
func (t *Three) BoxGeometry() *BoxGeometry {
	p := t.ctx.Get("BoxGeometry")
	return BoxGeometryFromJSObject(p)
}

// BoxGeometryFromJSObject returns a wrapped BoxGeometry JavaScript class.
func BoxGeometryFromJSObject(p *js.Object) *BoxGeometry {
	return &BoxGeometry{GeometryFromJSObject(p)}
}

// BoxGeometryOpts represents options passed to the BoxGeometry constructor.
type BoxGeometryOpts struct {
	WidthSegments  int // Number of segmented faces along the width of the sides. Default is 1.
	HeightSegments int // Number of segmented faces along the height of the sides. Default is 1.
	DepthSegments  int // Number of segmented faces along the depth of the sides. Default is 1.
}

// NewBoxGeometry returns a new BoxGeometry object.
//
//     width — Width of the sides on the X axis.
//     height — Height of the sides on the Y axis.
//     depth — Depth of the sides on the Z axis.
func (t *Three) NewBoxGeometry(width, height, depth float64, opts *BoxGeometryOpts) *BoxGeometry {
	p := t.ctx.Get("BoxGeometry")
	if opts != nil {
		p = p.New(width, height, depth, opts.WidthSegments, opts.HeightSegments, opts.DepthSegments)
	} else {
		p = p.New(width, height, depth)
	}
	return BoxGeometryFromJSObject(p)
}
