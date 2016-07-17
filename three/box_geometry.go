// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

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
	widthSegments  int // Number of segmented faces along the width of the sides. Default is 1.
	heightSegments int // Number of segmented faces along the height of the sides. Default is 1.
	depthSegments  int // Number of segmented faces along the depth of the sides. Default is 1.
}

// NewBoxGeometry returns a new BoxGeometry object.
//
//     width — Width of the sides on the X axis.
//     height — Height of the sides on the Y axis.
//     depth — Depth of the sides on the Z axis.
func (t *Three) NewBoxGeometry(width, height, depth float64, opts *BoxGeometryOpts) *BoxGeometry {
	p := t.ctx.Get("BoxGeometry")
	if opts != nil {
		p = p.New(width, height, depth, opts.widthSegments, opts.heightSegments, opts.depthSegments)
	} else {
		p = p.New(width, height, depth)
	}
	return BoxGeometryFromJSObject(p)
}
